package main

import (
	"context"
	"log"
	"math"
	"strings"
	//"io/ioutil"
	"math/rand"
	"fmt"
	"os"
	"strconv"
	"io"
	//"bufio"
	/*
	"encoding/csv"
	*/
	"time"

	pb "2020_2-Sistemas_Distribuidos/Tarea2/client_datanode"
	pb2 "2020_2-Sistemas_Distribuidos/Tarea2/client_namenode"

	"google.golang.org/grpc"
)

//struct util para almacenar la parte y el chunk correspondiente
type pieces struct {
	part string
	chunk []byte
}

//constante de direcciones 
const (
	addrDN1 = "10.6.40.238:9000"
	addrDN2 = "10.6.40.239:9001"
	addrDN3 = "10.6.40.240:9002"
	addrName = "10.6.40.241:9020"
)

//variables gloables utiles para despues
var cDN1, cDN2, cDN3 pb.GreeterClient
var connDN1, connDN2, connDN3 *grpc.ClientConn

//este realiza la conexion del cliente con algun DataNode
func makeConnection(address string) (pb.GreeterClient, *grpc.ClientConn){
	log.Printf("Haciendo conexion %s", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se conecto: %v", err)
	}
	c := pb.NewGreeterClient(conn)
	log.Printf("logre conexion %s", address)
	
	return c, conn
}

//este realiza la conexion del cliente con el NameNode
func makeConnectionToName(address string) (pb2.GreeterClient, *grpc.ClientConn){
	log.Printf("Haciendo conexion %s", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se conecto: %v", err)
	}
	c := pb2.NewGreeterClient(conn)
	log.Printf("logre conexion %s", address)
	
	return c, conn
}

/*
Esta funcion contiene toda la base para pedir un libro que este almacendado 
primero pide todos los libros disponibles, luego se elige un libro. Se le pide
al NameNode las partes y direcciones de estas en los DataNodes.
Recibe la información y luego procede a pedirle a los DataNodes las partes del libro
las recibe junta todas las partes y guarde el libro en la carpeta results como NEWnombre_del_libro.
*/
func askForBook(){
	var menu int;
	
	log.Printf("[1] Pedir libro")
	log.Printf("[2] Deja de pedir libro")
	fmt.Scanf("%d", &menu)
	for menu != 2 {
		log.Printf("Libros disponibles:")
		c, conn := makeConnectionToName(addrName)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		stream, err := c.GimmeBooks(ctx, &pb2.Result{Msg: "1"})
		if err != nil {
			log.Fatalf("Error on get books: %v", err)
		}
		for {
			book, err2 := stream.Recv()
			if err2 == io.EOF {
				break
			}

			log.Printf("%s", book.GetName())
			if err2 != nil {
				log.Fatalf("F")
			}
		}
		cancel()
		conn.Close()
		var name string;
		log.Printf("Escriba libro a pedir:")
		fmt.Scanf("%s", &name)
		c2, conn2 := makeConnectionToName(addrName)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		r, err2 := c2.GimmePartsDirections(ctx, &pb2.Book{Name: name})
		if err2 != nil {
			log.Fatalf("Ayuda :c")
		}
		log.Printf("%s", r)
		cancel()
		conn2.Close()

		parts := strings.Split(r.GetMsg(), "/")
		chunks := []pieces{}
		newFileName := "NEW" + name + ".pdf"
		_, err = os.Create("./results/" + newFileName)

		if err != nil {
		     fmt.Println(err)
		     os.Exit(1)
		}

		//set the newFileName file to APPEND MODE!!
		// open files r and w

		file, err2 := os.OpenFile("./results/" + newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

		if err2 != nil {
		     fmt.Println(err2)
		     os.Exit(1)
		}

		// IMPORTANT! do not defer a file.Close when opening a file for APPEND mode!
		// defer file.Close()

		// just information on which part of the new file we are appending
		for _, part := range parts {
			log.Printf("Voy a pedir %s", part)

			partOfPart := strings.Split(part, " ")

			partOfPartOfPart := strings.Split(partOfPart[0], "_")

			addr := ""
			if strings.Contains(addrDN1, partOfPart[1]) {
				addr = addrDN1
			} else if strings.Contains(addrDN2, partOfPart[1]) {
				addr = addrDN2
			} else if strings.Contains(addrDN3, partOfPart[1]) {
				addr = addrDN3
			}
			c1, conn1 := makeConnection(addr)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			r, err := c1.GimmeChunk(ctx, &pb.Chunkpart{Id: partOfPartOfPart[2], Name: name})
			if err != nil {
				log.Fatalf("Ayuda :c")
			}
			piece := pieces{part: partOfPartOfPart[2], chunk: r.GetMsg()}
			chunks = append(chunks, piece)
			cancel()
			conn1.Close()
		}

		for j := uint64(0); j < uint64(len(chunks)); j++ {
			_, err := file.Write(chunks[j].chunk)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			file.Sync()
		}

		file.Close()
		log.Printf("[1] Pedir libro")
		log.Printf("[2] Deja de pedir libro")
		fmt.Scanf("%d", &menu)
	}

	log.Printf("Hasta la proxima :D")
}

/*
Esta función tiene toda la lógica de enviar un libro. Lo que hace es 
pedir el libro a enviar, separa el libro en chunks y se los manda a un DataNode
al azar para que el se encargue de distribuir y espera la respuesta de que se guardo el libro
*/
func sendBook() {
	var menu int;

	log.Printf("[1] Enviar libro")
	log.Printf("[2] No enviar libro")
	fmt.Scanf("%d", &menu)
	for menu != 2 {
		var name string;
		log.Printf("Escriba libro a subir:")
		fmt.Scanf("%s", &name)
	
		fileToBeChunked := "./test/" + name + ".pdf" // change here!
		file, err := os.Open(fileToBeChunked)
		if err != nil {
		    fmt.Println(err)
		    os.Exit(1)
		}
		defer file.Close()
		fileInfo, _ := file.Stat()
		var fileSize int64 = fileInfo.Size()
		const fileChunk = 250000 // 1 MB, change this to your requirement
		// calculate total number of parts the file will be chunked into
		totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
		log.Printf("Splitting to %d pieces.\n", totalPartsNum)

		dns := []int{1, 2, 3}
		dns2 := []int{1, 2, 3}
		random := rand.Intn(3)

		dns2[0] = dns[random]
		copy(dns[random:], dns[random + 1:])
		dns = dns[:len(dns) - 1]

		random = rand.Intn(2)

		dns2[1] = dns[random]
		copy(dns[random:], dns[random + 1:])
		dns = dns[:len(dns) - 1]

		dns2[2] = dns[0]

		for j := 0; j < 3; j++ {
			log.Printf("Lo voy a enviar a %d", dns2[j])
			if dns2[j] == 1 {
				cDN1, connDN1 = makeConnection(addrDN1)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				stream, err := cDN1.Upload(ctx)
				if err != nil {
					log.Printf("No pude conectarme, intentare con otro :c")
					continue
				}

				for i := uint64(0); i < totalPartsNum; i++ {
					partSize := int(math.Min(fileChunk, float64(fileSize - int64(i * fileChunk))))
					partBuffer := make([]byte, partSize)
					file.Read(partBuffer)
					stream.Send(&pb.Book{Name: name, Part: strconv.FormatUint(i, 10), Chunk: partBuffer})
				}

				r, err2 := stream.CloseAndRecv()
				if err2 != nil {
					log.Fatalf("Fallaste :c")
				}
				log.Printf("%s", r.GetMsg())
				cancel()
				connDN1.Close()
				break
			} else if dns2[j] == 2 {
				cDN2, connDN2 = makeConnection(addrDN2) 
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				stream, err := cDN2.Upload(ctx)
				if err != nil {
					log.Printf("No pude conectarme, intentare con otro :c")
					continue
				}

				for i := uint64(0); i < totalPartsNum; i++ {
					partSize := int(math.Min(fileChunk, float64(fileSize - int64(i * fileChunk))))
					partBuffer := make([]byte, partSize)
					file.Read(partBuffer)
					stream.Send(&pb.Book{Name: name, Part: strconv.FormatUint(i, 10), Chunk: partBuffer})
				}

				r, err2 := stream.CloseAndRecv()
				if err2 != nil {
					log.Fatalf("Fallaste :c")
				}
				log.Printf("%s", r.GetMsg())
				cancel()
				connDN2.Close()
				break
			} else {
				cDN3, connDN3 = makeConnection(addrDN3) 
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				stream, err := cDN3.Upload(ctx)
				if err != nil {
					log.Printf("No pude conectarme, intentare con otro :c")
					continue
				}

				for i := uint64(0); i < totalPartsNum; i++ {
					partSize := int(math.Min(fileChunk, float64(fileSize - int64(i * fileChunk))))
					partBuffer := make([]byte, partSize)
					file.Read(partBuffer)
					stream.Send(&pb.Book{Name: name, Part: strconv.FormatUint(i, 10), Chunk: partBuffer})
				}

				r, err2 := stream.CloseAndRecv()
				if err2 != nil {
					log.Fatalf("Fallaste :c")
				}
				log.Printf("%s", r.GetMsg())
				cancel()
				connDN3.Close()
				break
			}
		}
		log.Printf("[1] Enviar libro")
		log.Printf("[2] No enviar libro")
		fmt.Scanf("%d", &menu)
	}
	log.Printf("Gracias por compartir tus conocimientos")
}

/*
Pregunta si es Uploader y Downloader y luego parte el juego
*/
func main() {
	var upDown int
	log.Printf("[1] Uploader")
	log.Printf("[2] Downloader")
	fmt.Scanf("%d", &upDown)
	if upDown == 1 {
		sendBook()
	} else {
		askForBook()
	}
}