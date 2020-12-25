package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"sync"
	"os"
	"fmt"
	"time"
	"bufio"
	"strings"
	/*
	"encoding/csv"
	"encoding/json"
	*/

	pb "2020_2-Sistemas_Distribuidos/Tarea2/datanode_namenode"
	pb2 "2020_2-Sistemas_Distribuidos/Tarea2/client_namenode"

	//"github.com/streadway/amqp"
	"google.golang.org/grpc"
)


//constantes de direccionesx
const (
	portClient = ":9020"
	port = ":9010"
	addrDN1 = "10.6.40.238:9053"
	addrDN2 = "10.6.40.239:9054"
	addrDN3 = "10.6.40.240:9055"
)

//struct de servidores
type server struct {
	pb.UnimplementedGreeterServer
}

type server2 struct {
	pb2.UnimplementedGreeterServer
}


//util para el algoritmo centrlizado hacer que no todos modifiquen el archivo a la vez
type Touching struct {
	files int
	mu sync.Mutex
}

//globales utiles
var using Touching
var quienSoy string


/*
Funcion del algoritmo centralizado que recibe una propuesta del DataNode, verifica
que esten activos y en el caso de que haya algun nodo no activo, entonces este generá una
nueva propuesta
*/
func (s *server) Propose(ctx context.Context, in *pb.Distribution) (*pb.Accepted, error) {
	//TEST_MSG log.Printf("Oh se han comunicado!")
	//SHOW_DETAILS log.Printf("Distribution %s %s %s %s", in.GetName(), in.GetChunk1(), in.GetChunk2(), in.GetChunk3())
	var on1, on2, on3 = true, true, true

	//SHOW_DETAILS log.Printf("Voy a checkear DN1")
	if !check(addrDN1) {
		//SHOW_DETAILS log.Printf("No me respondo DN1 F")
		on1 = false
	} else {
		//SHOW_DETAILS log.Printf("Esta vivo DN1!")
	}

	if !check(addrDN2) {
		//SHOW_DETAILS log.Printf("No me respondo DN2 F")
		on2 = false
	} else {
		//SHOW_DETAILS log.Printf("Esta vivo DN2!")
	}

	if !check(addrDN3) {
		//SHOW_DETAILS log.Printf("No me respondo DN3 F")
		on3 = false
	} else {
		//SHOW_DETAILS log.Printf("Esta vivo DN3!")
	}

	var quant = 0
	if(on1){
		quant++
	}

	if(on2){
		quant++
	}

	if(on3){
		quant++
	}

	var accept *pb.Accepted
	if quant == 3 {
		accept = &pb.Accepted{Chunk1: in.GetChunk1(), Chunk2: in.GetChunk2(), Chunk3: in.GetChunk3()}
	} else {
		val1, _ := strconv.ParseUint(in.GetChunk1(), 10, 64)
		val2, _ := strconv.ParseUint(in.GetChunk2(), 10, 64)
		val3, _ := strconv.ParseUint(in.GetChunk3(), 10, 64)
		cantChunks := val1 + val2 + val3
		var chunks1, chunks2 uint64
		chunks1, chunks2 = 0, 0

		if quant == 2 {
			for i := uint64(0); cantChunks > i; i++ {
				if i % 2 == 0 {
					chunks1++
				} else {
					chunks2++
				}
			}
			
			if on1 == false {
				accept = &pb.Accepted{Chunk1: "0", Chunk2: strconv.FormatUint(chunks1, 10), Chunk3: strconv.FormatUint(chunks2, 10)}
			}

			if on2 == false {
				accept = &pb.Accepted{Chunk1: strconv.FormatUint(chunks1, 10), Chunk2: "0", Chunk3: strconv.FormatUint(chunks2, 10)}
			}

			if on3 == false {
				accept = &pb.Accepted{Chunk1: strconv.FormatUint(chunks1, 10), Chunk2: strconv.FormatUint(chunks2, 10), Chunk3: "0"}
			}
		} else {
			if (on1 == false && on2 == false) {
				accept = &pb.Accepted{Chunk1: "0", Chunk2: "0", Chunk3: strconv.FormatUint(cantChunks, 10)}
			}

			if (on2 == false && on3 == false) {
				accept = &pb.Accepted{Chunk1: strconv.FormatUint(cantChunks, 10), Chunk2: "0", Chunk3: "0"}
			}

			if (on1 == false && on3 == false) {
				accept = &pb.Accepted{Chunk1: "0", Chunk2: strconv.FormatUint(cantChunks, 10), Chunk3: "0"}
			}
		}
	}
	//SHOW_DETAILS log.Printf("Nueva Distribution %s %s %s", accept.GetChunk1(), accept.GetChunk2(), accept.GetChunk3())
	result1, _ := strconv.ParseUint(accept.GetChunk1(), 10, 64)
	result2, _ := strconv.ParseUint(accept.GetChunk2(), 10, 64)
	result3, _ := strconv.ParseUint(accept.GetChunk3(), 10, 64)
	using.mu.Lock()
	defer using.mu.Unlock()
	writeInFile(in.GetName(), result1, result2, result3)
	return accept, nil
}


/*
Funcion del algoritmo distribuido que recibe la distribucion realizada y los manda para guardar en el registro
*/
func (s *server) Proposedis(ctx context.Context, in *pb.Distribution) (*pb.Distribution, error) {
	//TEST_MSG log.Printf("Oh se han comunicado!")
	//SHOW_DETAILS log.Printf("Distribution %s %s %s %s", in.GetName(), in.GetChunk1(), in.GetChunk2(), in.GetChunk3())
	result1, _ := strconv.ParseUint(in.GetChunk1(), 10, 64)
	result2, _ := strconv.ParseUint(in.GetChunk2(), 10, 64)
	result3, _ := strconv.ParseUint(in.GetChunk3(), 10, 64)
	
	writeInFile(in.GetName(), result1, result2, result3)
	return in, nil
}


/*
Función de comunicacion del cliente con el NameNode el cual recibe un nombre del libro, y le entrega
los nombres de las partes del libro y donde se encuentran
*/
func (s *server2) GimmePartsDirections(ctx context.Context, in *pb2.Book) (*pb2.Result, error) {
	//SHOW_DETAILS log.Printf("Me pidieron %s", in.GetName())
	result := readFile(in.GetName())
	//SHOW_DETAILS log.Printf("%s", result)
	return &pb2.Result{Msg: result}, nil
}


/*
Función de comunicacion del cliente con el NameNode, el cual recibe un mensaje para saber que debe
enviar todos los libros almacenados actualmente. Luego entrega al cliente los nombres de los libres
disponibles
*/
func (s *server2) GimmeBooks(in *pb2.Result, stream pb2.Greeter_GimmeBooksServer) error {
	//SHOW_DETAILS log.Printf("Te buscare todos los libritos")
	using.mu.Lock()
	defer using.mu.Unlock()
	f, err := os.Open("LOG.txt") 
  	defer f.Close()
    if err != nil { 
        log.Fatalf("failed to open") 
    }

    scanner := bufio.NewScanner(f) 
    scanner.Split(bufio.ScanLines) 
  
    for scanner.Scan() { 
    	text := scanner.Text()
        if strings.Contains(text, "parte") {
        	continue
        }

        text = strings.Split(text, " ")[0]
    	text = strings.Split(text, "_")[0]
    	log.Printf("%s", text)
        err := stream.Send(&pb2.Book{Name: text})
        if err != nil {
        	return err
        }
    }
	return nil
}

/*
Lee el archivo LOG y un string que contenga todas las partes y las dirreeciones respectivos, todas
separadas con un /
*/
func readFile(book string) string {
	using.mu.Lock()
	defer using.mu.Unlock()
	f, err := os.Open("LOG.txt") 
  	defer f.Close()
    if err != nil { 
        log.Fatalf("failed to open") 
    }

    scanner := bufio.NewScanner(f) 
    scanner.Split(bufio.ScanLines) 
    var text []string 
  
    for scanner.Scan() { 
        text = append(text, scanner.Text()) 
    }

    quantLines := -1
    i := 0
    finals := []string{}
    for _, eachLine := range text {
    	if quantLines == -1 {
	    	if strings.Contains(eachLine, book) {
	    		space := " "
	    		separeted := strings.Split(eachLine, space)
	    		quantLines, _ = strconv.Atoi(separeted[1])
	    	}
	    } else {
	    	if i >= quantLines {
	    		break
	    	}

	    	i++
	    	finals = append(finals, eachLine)
	    }
    }

    return strings.Join(finals, "/")
}

/*
Recibe la distribución del libre y la registra
*/
func writeInFile(name string, chunks1 uint64, chunks2 uint64, chunks3 uint64){
	
	using.files++
	f, err := os.OpenFile("LOG.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Abri mal el archivo si vamo vamo")
	}

	//SHOW_DETAILS log.Printf("Escribiendo LOG esperen porfa...")
	
	_, err = fmt.Fprintln(f, name + "_" + strconv.Itoa(using.files) + " " + strconv.FormatUint(chunks1 + chunks2 + chunks3, 10))
	if err != nil {
		log.Fatalf("%s", err)
		f.Close()
	}

	for i := uint64(0); i < chunks1; i++ {
		fmt.Fprintln(f, "parte_" + strconv.Itoa(using.files) + "_" + strconv.FormatUint(i, 10) + " " + "10.6.40.238")
	}

	for i := uint64(0); i < chunks2; i++ {
		fmt.Fprintln(f, "parte_" + strconv.Itoa(using.files) + "_" + strconv.FormatUint(i + chunks1, 10) + " " + "10.6.40.239")
	}

	for i := uint64(0); i < chunks3; i++ {
		fmt.Fprintln(f, "parte_" + strconv.Itoa(using.files) + "_" + strconv.FormatUint(i + chunks1 + chunks2, 10) + " " + "10.6.40.240")
	}	

	using.files = using.files
}

/*
Permite checkear si es que un datanode esta activo o no, intentando hacer comunnicacion, donde 
si no se puede entonces retorna false, si se puede quiere decir que esta activo entonces retorna true 
*/
func check(addr string) bool {
	c, conn := connect(addr)
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.Check(ctx, &pb.Ping{Msg: "1"})
	if err != nil {
		return false
	}

	return true
}

/*
Realiza la conexion del NameNode hacia el DataNode
*/
func connect(addr string) (pb.GreeterClient, *grpc.ClientConn){
	//SHOW_DETAILS log.Printf("Haciendo conexion %s", addr)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se conecto: %v", err)
	}
	c := pb.NewGreeterClient(conn)
	//SHOW_DETAILS log.Printf("logre conexion %s", addr)
	
	return c, conn
}

/*
Abre puerto para recibir mensajes de los DataNodes
*/
func openPortDN(port string) {
	//SHOW_DETAILS log.Printf("Preparando... %s", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Falle al escuchar: %v", err)
	}
	//SHOW_DETAILS log.Printf("Escuchando...")
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falle al servir: %v", err)
	}
}

/*
Abre puerto para recibir mensajes de clientes
*/
func openPortCL(port string) {
	//SHOW_DETAILS log.Printf("Preparando... %s", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Falle al escuchar: %v", err)
	}
	//SHOW_DETAILS log.Printf("Escuchando...")
	s := grpc.NewServer()
	pb2.RegisterGreeterServer(s, &server2{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falle al servir: %v", err)
	}
}


func main(){
	using.files = 0
	go openPortDN(port)
	openPortCL(portClient)
}