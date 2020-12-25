package main

import (
	"context"
	"log"
	"net"
	"fmt"
	"io"
	"strconv"
	"time"
	"os"
	"io/ioutil"
	"bufio"
	/*
	"encoding/csv"
	"encoding/json"
	"sync"
	*/

	pb "2020_2-Sistemas_Distribuidos/Tarea2/client_datanode"
	pb2 "2020_2-Sistemas_Distribuidos/Tarea2/datanode_namenode"
	pb3 "2020_2-Sistemas_Distribuidos/Tarea2/datanode_datanode"

	//"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

//definicion de servidos de struct
type server struct {
	pb.UnimplementedGreeterServer
}

type server2 struct {
	pb2.UnimplementedGreeterServer
}

type server3 struct {
	pb3.UnimplementedGreeterServer
}


//constantes definidos con puertos
const (
	//Receive Client
	portDN1 = ":9000"
	portDN2 = ":9001"
	portDN3 = ":9002"
	//Receive Datanode
	portDN1_2 = ":9050"
	portDN2_2 = ":9051"
	portDN3_2 = ":9052"
	//Receive Namenode
	portDN1_3 = ":9053"
	portDN2_3 = ":9054"
	portDN3_3 = ":9055"
	//IPS
	addrDN1 = "10.6.40.238:9050"
	addrDN2 = "10.6.40.239:9051"
	addrDN3 = "10.6.40.240:9052"
	addrNN = "10.6.40.241:9010"
	//Algorithm Ricart and Agrawala
	RELEASED = 1
	WANTED = 2
	HELD = 3
)

//variables globales utiles
var state = RELEASED
var quienSoy, distOrCentr int
var cName pb2.GreeterClient
var cDN1, cDN2, cDN3 pb3.GreeterClient
var connName, connDN1, connDN2, connDN3 *grpc.ClientConn

/*
funcion de comunicacion del cliente al DataNode, donde este recibe una serie de chunks
y luego la distribuye siguiendo el algoritmo centralizado o distribuido dependiendo de
como se haya comenzado el DataNode
*/
func (s *server) Upload(stream pb.Greeter_UploadServer) error {
	var books []*pb.Book
	for {

		book, error := stream.Recv()
		if error == io.EOF {
			break
		}
		//SHOW_DETAILS log.Printf("%s %s", book.GetName(), book.GetPart())
		books = append(books, book)
	}

	if distOrCentr == 1 {
		centralizedMutualExclusion(books)
	} else {
		distributedMutualExclusion(books)
	}
	return stream.SendAndClose(&pb.Result{Msg: "Se envio respuesta."})
}

/*
Funcion de comunicacion del cliente con el DataNode, donde este recibe el nombre de una parte 
y este le envia el Chunk correspondiente al clinte
*/
func (s *server) GimmeChunk(ctx context.Context, in *pb.Chunkpart) (*pb.Chunk, error) {
	//SHOW_DETAILS log.Printf("Me pidieron un chunk, toma...")
	//read a chunk
	currentChunkFileName := in.GetName() + "_" + in.GetId()

	newFileChunk, err := os.Open("./chunks/" + currentChunkFileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer newFileChunk.Close()

	chunkInfo, err := newFileChunk.Stat()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// calculate the bytes size of each chunk
	// we are not going to rely on previous data and constant

	var chunkSize int64 = chunkInfo.Size()
	chunkBufferBytes := make([]byte, chunkSize)

	// read into chunkBufferBytes
	reader := bufio.NewReader(newFileChunk)
	_, err = reader.Read(chunkBufferBytes)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &pb.Chunk{Msg: chunkBufferBytes}, nil
}

/*
funcion de comunicacion entre DataNode y NaemNode que le permite al Namenode checkear 
si es este datanode esta activo o no
*/
func (s *server2) Check(ctx context.Context, in *pb2.Ping) (*pb2.Pong, error) {
	//TEST_MSG log.Printf("Oh se han comunicado!")
	//SHOW_DETAILS log.Printf("Soy pingiado")
	return &pb2.Pong{Msg: "1"}, nil
}

/*
funcion de comunicacion entre DataNode y DataNode donde un datanode le envia a otros los chunks
que debe almecenar el
*/
func (s *server3) Savechunk(stream pb3.Greeter_SavechunkServer) error {
	//TEST_MSG log.Printf("Oh se han comunicado!")
	//SHOW_DETAILS log.Printf("Recibido chunks")
	var books []*pb3.Book
	for {

		book, error := stream.Recv()
		if error == io.EOF {
			//SHOW_DETAILS log.Printf("Voy a enviar")
			writeChunks(books)
			return stream.SendAndClose(&pb3.Result{Msg: "Se han recibidos chunks"})
		}
		//SHOW_DETAILS log.Printf("%s %s", book.GetName(), book.GetPart())
		books = append(books, book)
	}
}
/*
funcion de comunicacion entre DataNode y DataNode donde este recibe una priopuesta,
y verifica si es que es valida y en el caso que no lo sea le dice donde estuvo el problema
*/
func (s *server3) Showme(ctx context.Context, in *pb3.Distribution) (*pb3.Problems, error) {
	//TEST_MSG log.Printf("Oh se han comunicado!")
	//SHOW_DETAILS log.Printf("Mmmmm distribution %s %s %s", in.GetChunk1(), in.GetChunk2(), in.GetChunk3())
	var dn1, dn2, dn3 = "0", "0", "0"

	n1, _ := strconv.ParseUint(in.GetChunk1(), 10, 64)
	n2, _ := strconv.ParseUint(in.GetChunk2(), 10, 64)
	n3, _ := strconv.ParseUint(in.GetChunk3(), 10, 64)
	
	if  n1 > 0 {
		if !check(addrDN1) {
			//SHOW_DETAILS log.Printf("No ta 1, no va a funcionar propuesta")
			dn1 = "1"
		} else {
			//SHOW_DETAILS log.Printf("Ta vivo 1")
		}
	}

	if n2 > 0 {
		if !check(addrDN2) {
			//SHOW_DETAILS log.Printf("No ta 2, no va a funcionar propuesta")
			dn2 = "2"
		} else {
			//SHOW_DETAILS log.Printf("Ta vivo 2")
		}
	}

	if n3 > 0 {
		if !check(addrDN3) {
			//SHOW_DETAILS log.Printf("No ta 3, no va a funcionar propuesta")
			dn3 = "3"
		} else {
			//SHOW_DETAILS log.Printf("Ta vivo 3")
		}
	}
	return &pb3.Problems{Dn1: dn1, Dn2: dn2, Dn3: dn3}, nil
}

/*
funcion de comunicacion entre DataNode y DataNode que le permite a un DataNode checkear 
si es esta activo o no
*/
func (s *server3) Check2(ctx context.Context, in *pb3.Ping) (*pb3.Pong, error) {
	//TEST_MSG log.Printf("Oh se han comunicado!")
	//SHOW_DETAILS log.Printf("Soy pingiado")
	return &pb3.Pong{P: "1"}, nil
}

/*Funcion que sirve para aplicar ricart agrawala*/
func (s *server3) Ricart(ctx context.Context, in *pb3.Node) (*pb3.NoNode, error) {
	//TEST_MSG log.Printf("Oh se han comunicado!")
	//SHOW_DETAILS log.Printf("Mmmmm me hablo %s", in.GetWho())
	for (state == HELD || (state == WANTED && in.GetWho() < strconv.Itoa(quienSoy))) {}
	return &pb3.NoNode{Msg: "1"}, nil
}	

/*Funcion que envia un mensaje tipo check*/
func check(addr string) bool {
	c, conn := connect(addr)
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.Check2(ctx, &pb3.Ping{P: "1"})
	if err != nil {
		return false
	}

	return true
}


/*Fuincion que envia mensaje a travé sde Ricart para hacer la spesra*/
func waitForRicart(addr string, prip *bool) {
	c, conn := connect(addr)
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.Ricart(ctx, &pb3.Node{Who: strconv.Itoa(quienSoy)})
	if err != nil {
		*prip = true
	} else {
		*prip = false
	}

}


/*Envia una propuesta a un datanode vecino y recibe donde fallo*/
func sendProposalToDN(addr string, who string, chunks1 string, chunks2 string, chunks3 string, flag *int) []int {
	//SHOW_DETAILS log.Printf("Enviando a %s", addr)
	c, conn := connect(addr)
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Showme(ctx, &pb3.Distribution{Who: who, Chunk1: chunks1, Chunk2: chunks2, Chunk3: chunks3})
	culpable := []int{}
	if err != nil {
		//SHOW_DETAILS log.Printf("Fallaste en enviar a %s :c", addr)
		*flag = 1
	} else {
		//SHOW_DETAILS log.Printf("%s %s %s", r.GetDn1(), r.GetDn2(), r.GetDn3())
		if r.GetDn1() == "1" {
			culpable = append(culpable, 1)
		}

		if r.GetDn2()== "2" {
			culpable = append(culpable, 2)
		}

		if r.GetDn3()== "3" {
			culpable = append(culpable, 3)
		}
		*flag = 2
	}

	//SHOW_DETAILS log.Println("Estos le fallaron a", addr, culpable)
	return culpable
}

/*Realiza toda la lógica del algoritmo distribuido*/
func distributedMutualExclusion(books []*pb.Book){
	cantChunks := len(books)
	var chunks1, chunks2, chunks3 uint64
	whoIsOn := []bool{true, true, true}
	for {
		//SHOW_DETAILS log.Printf("Creando propuesta...")
		chunks1, chunks2, chunks3 = 0, 0, 0
		if whoIsOn[0] && whoIsOn[1] && whoIsOn[2] {
			for i := 0; cantChunks > i; i++ {
				if i % 3 == 0 {
					chunks1++
				} else if i % 3 == 1 {
					chunks2++
				} else {
					chunks3++
				}
			}
			//SHOW_DETAILS log.Printf("Viendo si me validan")
			flag1, flag2 := 0, 0
			if quienSoy == 1 {
				culpable := sendProposalToDN(addrDN2, "1", strconv.FormatUint(chunks1, 10), strconv.FormatUint(chunks2, 10), strconv.FormatUint(chunks3, 10), &flag1)
				culpable2 := sendProposalToDN(addrDN3, "1", strconv.FormatUint(chunks1, 10), strconv.FormatUint(chunks2, 10), strconv.FormatUint(chunks3, 10), &flag2)
				for cupl := 0; cupl < len(culpable); cupl++ {
					whoIsOn[cupl] = false
				}
				for cupl := 0; cupl < len(culpable2); cupl++ {
					whoIsOn[cupl] = false
				}
			} else if quienSoy == 2 {
				culpable := sendProposalToDN(addrDN1, "2", strconv.FormatUint(chunks1, 10), strconv.FormatUint(chunks2, 10), strconv.FormatUint(chunks3, 10), &flag1)
				culpable2 := sendProposalToDN(addrDN3, "2", strconv.FormatUint(chunks1, 10), strconv.FormatUint(chunks2, 10), strconv.FormatUint(chunks3, 10), &flag2)
				for cupl := 0; cupl < len(culpable); cupl++ {
					whoIsOn[cupl] = false
				}
				for cupl := 0; cupl < len(culpable2); cupl++ {
					whoIsOn[cupl] = false
				}
			} else if quienSoy == 3 {
				culpable := sendProposalToDN(addrDN1, "3", strconv.FormatUint(chunks1, 10), strconv.FormatUint(chunks2, 10), strconv.FormatUint(chunks3, 10), &flag1)
				culpable2 := sendProposalToDN(addrDN2, "3", strconv.FormatUint(chunks1, 10), strconv.FormatUint(chunks2, 10), strconv.FormatUint(chunks3, 10), &flag2)
				for cupl := 0; cupl < len(culpable); cupl++ {
					whoIsOn[cupl] = false
				}
				for cupl := 0; cupl < len(culpable2); cupl++ {
					whoIsOn[cupl] = false
				}
			}

			if flag1 == 2 && flag2 == 2 {
				//SHOW_DETAILS log.Printf("Listo pase libre pa mandar libro")
				break
			}

			if flag1 == 1 {
				if quienSoy == 1 {
					//SHOW_DETAILS log.Printf("Oh no 2 no esta")
					whoIsOn[1] = false
				} else if quienSoy == 2 || quienSoy == 3 {
					//SHOW_DETAILS log.Printf("Oh no 1 no esta")
					whoIsOn[0] = false
				}
			}

			if flag2 == 1 {
				if quienSoy == 1 || quienSoy == 2 {
					//SHOW_DETAILS log.Printf("Oh no 3 no esta")
					whoIsOn[2] = false
				} else if quienSoy == 3 {
					//SHOW_DETAILS log.Printf("Oh no 1 no esta")
					whoIsOn[1] = false
				}
			}
		} else if whoIsOn[0] && whoIsOn[1] {
			for i := 0; cantChunks > i; i++ {
				if i % 2 == 0 {
					chunks1++
				} else {
					chunks2++
				}
			}
			//SHOW_DETAILS log.Printf("Viendo si me validan")
			flag := 0
			if quienSoy == 1 {
				culpable := sendProposalToDN(addrDN2, "1", strconv.FormatUint(chunks1, 10), strconv.FormatUint(chunks2, 10), "0", &flag)
				for cupl := 0; cupl < len(culpable); cupl++ {
					whoIsOn[cupl] = false
				}
			} else if quienSoy == 2 {
				culpable := sendProposalToDN(addrDN1, "2", strconv.FormatUint(chunks1, 10), strconv.FormatUint(chunks2, 10), "0", &flag)
				for cupl := 0; cupl < len(culpable); cupl++ {
					whoIsOn[cupl] = false
				}
			}

			if flag == 2 {
				//SHOW_DETAILS log.Printf("Listo pase libre pa mandar libro")
				break
			}

			if flag == 1 {
				if quienSoy == 1 {
					//SHOW_DETAILS log.Printf("Oh no 2 no esta")
					whoIsOn[1] = false
				} else if quienSoy == 2 {
					//SHOW_DETAILS log.Printf("Oh no 1 no esta")
					whoIsOn[0] = false
				}
			}
		} else if whoIsOn[0] && whoIsOn[2] {
			for i := 0; cantChunks > i; i++ {
				if i % 2 == 0 {
					chunks1++
				} else {
					chunks3++
				}
			}
			//SHOW_DETAILS log.Printf("Viendo si me validan")
			flag := 0
			if quienSoy == 1 {
				culpable := sendProposalToDN(addrDN3, "1", strconv.FormatUint(chunks1, 10), "0", strconv.FormatUint(chunks3, 10), &flag)
				for cupl := 0; cupl < len(culpable); cupl++ {
					whoIsOn[cupl] = false
				}
			} else if quienSoy == 3 {
				culpable := sendProposalToDN(addrDN1, "3", strconv.FormatUint(chunks1, 10), "0", strconv.FormatUint(chunks3, 10), &flag)
				for cupl := 0; cupl < len(culpable); cupl++ {
					whoIsOn[cupl] = false
				}
			}

			if flag == 2 {
				//SHOW_DETAILS log.Printf("Listo pase libre pa mandar libro")
				break
			}

			if flag == 1 {
				if quienSoy == 1 {
					//SHOW_DETAILS log.Printf("Oh no 3 no esta")
					whoIsOn[2] = false
				} else if quienSoy == 3 {
					//SHOW_DETAILS log.Printf("Oh no 1 no esta")
					whoIsOn[0] = false
				}
			}
		} else if whoIsOn[1] && whoIsOn[2] {
			for i := 0; cantChunks > i; i++ {
				if i % 2 == 0 {
					chunks2++
				} else {
					chunks3++
				}
			}
			//SHOW_DETAILS log.Printf("Viendo si me validan")
			flag := 0
			if quienSoy == 2 {
				culpable := sendProposalToDN(addrDN3, "2", "0", strconv.FormatUint(chunks2, 10), strconv.FormatUint(chunks3, 10), &flag)
				for cupl := 0; cupl < len(culpable); cupl++ {
					whoIsOn[cupl] = false
				}
			} else if quienSoy == 3 {
				culpable := sendProposalToDN(addrDN2, "3", "0", strconv.FormatUint(chunks2, 10), strconv.FormatUint(chunks3, 10), &flag)
				for cupl := 0; cupl < len(culpable); cupl++ {
					whoIsOn[cupl] = false
				}
			}

			if flag == 2 {
				//SHOW_DETAILS log.Printf("Listo pase libre pa mandar libro")
				break
			}

			if flag == 1 {
				if quienSoy == 2 {
					//SHOW_DETAILS log.Printf("Oh no 3 no esta")
					whoIsOn[2] = false
				} else if quienSoy == 3 {
					//SHOW_DETAILS log.Printf("Oh no 2 no esta")
					whoIsOn[1] = false
				}
			}
		} else {
			//SHOW_DETAILS log.Printf("Oh soy el unico vivo prendido...")
			if quienSoy == 1 {
				chunks1 = uint64(cantChunks)
			} else if quienSoy == 2 {
				chunks2 = uint64(cantChunks)
			} else {
				chunks3 = uint64(cantChunks)
			}
			break
		}
	}
	//SHOW_DETAILS log.Printf("Distribution final final %s %s %s", strconv.FormatUint(chunks1, 10), strconv.FormatUint(chunks2, 10), strconv.FormatUint(chunks3, 10))
	soyTiempo := time.Now()
	
	state = WANTED
	mehablo1 := false
	mehablo2 := false

	if quienSoy == 1 {
		if whoIsOn[1] {
			mehablo1 = true
			go waitForRicart(addrDN2, &mehablo1)
		}
		if whoIsOn[2] {
			mehablo2 = true
			go waitForRicart(addrDN3, &mehablo2)
		}
	} else if quienSoy == 2 {
		if whoIsOn[0] {
			mehablo1 = true
			go waitForRicart(addrDN1, &mehablo1)
		}
		if whoIsOn[2] {
			mehablo2 = true
			go waitForRicart(addrDN3, &mehablo2)
		}
	} else {
		if whoIsOn[0] {
			mehablo1 = true
			go waitForRicart(addrDN1, &mehablo1)
		}
		if whoIsOn[1] {
			mehablo2 = true
			go waitForRicart(addrDN2, &mehablo2)
		}
	}

	//SHOW_DETAILS log.Printf("Esperando respuesta")
	for mehablo2 || mehablo1 {

	}
	state = HELD
	//SHOW_DETAILS log.Printf("VOY A ESCRIBIR!")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	r, err := cName.Proposedis(ctx, &pb2.Distribution{Name: books[0].GetName(), Chunk1: strconv.FormatUint(chunks1, 10), Chunk2: strconv.FormatUint(chunks2, 10), Chunk3: strconv.FormatUint(chunks3, 10)})
	if err != nil {
		log.Fatalf("could not greet: %v %s", err, r)
	}
	//SHOW_DETAILS log.Printf("Distribution final final %s %s %s", r.GetChunk1(), r.GetChunk2(), r.GetChunk3())
	cancel()
	//SHOW_DETAILS log.Printf("Ya escribi vamo vamo")
	state = RELEASED
	soyResultado := time.Since(soyTiempo)

	log.Printf("Me demore en escribir en el LOG %s", soyResultado)
	sendChunks(books, chunks1, chunks2, chunks3)
}

/*realiza toda la logica del algoritmo centrlizado*/
func centralizedMutualExclusion(books []*pb.Book){
	cantChunks := len(books)
	var chunks1, chunks2, chunks3 uint64
	chunks1, chunks2, chunks3 = 0, 0, 0
	i := 0
	for cantChunks > i {
		if i % 3 == 0 {
			chunks1++
		} else if i % 3 == 1 {
			chunks2++
		} else {
			chunks3++
		}
		i++
	}
	//SHOW_DETAILS log.Printf("D1: %d D2: %d D3: %d", chunks1, chunks2, chunks3)

	soyTiempo := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	r, err := cName.Propose(ctx, &pb2.Distribution{Name: books[0].GetName(), Chunk1: strconv.FormatUint(chunks1, 10), Chunk2: strconv.FormatUint(chunks2, 10), Chunk3: strconv.FormatUint(chunks3, 10)})
	if err != nil {
		log.Fatalf("could not greet: %v %s", err, r)
	}
	//SHOW_DETAILS log.Printf("Distribution final final %s %s %s", r.GetChunk1(), r.GetChunk2(), r.GetChunk3())
	
	cancel()
	soyResultado := time.Since(soyTiempo)
	log.Printf("Me demore en escribir en el LOG %s", soyResultado)
	chunks1, _ = strconv.ParseUint(r.GetChunk1(), 10, 64)
	chunks2, _ = strconv.ParseUint(r.GetChunk2(), 10, 64)
	chunks3, _ = strconv.ParseUint(r.GetChunk3(), 10, 64)
	sendChunks(books, chunks1, chunks2, chunks3)
}

/*
Guardalos chunks en el maquina del datanode
*/
func writeChunks(books []*pb3.Book){
	//SHOW_DETAILS log.Printf("Guargando Chunks")
	for i := 0; i < len(books); i++ {
		fileName := books[i].GetName() + "_" + books[i].GetPart()
		//SHOW_DETAILS log.Printf(fileName)
		_, err := os.Create("./chunks/" + fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// write/save buffer to disk
		ioutil.WriteFile("./chunks/" + fileName, books[i].GetChunk(), os.ModeAppend)

		//SHOW_DETAILS fmt.Println("Split to : ", fileName)
	}
}

/*Envia los chunks a los datanoes y almacena los suyos*/
func sendChunks(books []*pb.Book, chunks1 uint64, chunks2 uint64, chunks3 uint64) {
	for j := 1; j <= 3; j++ {
		if quienSoy == j {
			//SHOW_DETAILS log.Printf("Yo mismo")
			var booksAyu []*pb3.Book
			var extra, quant uint64

			if j == 1 {
				extra = 0
				quant = chunks1
			} else if j == 2 {
				extra = chunks1
				quant = chunks2
			} else {
				extra = chunks1 + chunks2
				quant = chunks3
			}

			for i := uint64(0); i < quant; i++ {
				booksAyu = append(booksAyu, &pb3.Book{Name: books[i + extra].GetName(), Part: books[i + extra].GetPart(), Chunk: books[i + extra].GetChunk()})
			}
			writeChunks(booksAyu)
		} else if j == 1 {
			if chunks1 > 0 {
				//SHOW_DETAILS log.Printf("Enviando a 1")
				cDN1, connDN1 = connect(addrDN1)
				defer connDN1.Close()
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				stream, err := cDN1.Savechunk(ctx)
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}

				for i := uint64(0); i < chunks1; i++ {
					stream.Send(&pb3.Book{Name: books[i].GetName(), Part: books[i].GetPart(), Chunk: books[i].GetChunk()})
				}

				r, err2 := stream.CloseAndRecv()
				if err2 != nil {
					log.Fatalf("Fallaste en enviar a 1 :c: %s %s", err2, r)
				}
				//SHOW_DETAILS log.Printf("%s", r.GetMsg())
				cancel()
			} else {
				//SHOW_DETAILS log.Printf("El DN1 no ta")
			}
		} else if j == 2 {
			if chunks2 > 0 {
				//SHOW_DETAILS log.Printf("Enviando a 2")
				cDN2, connDN2 = connect(addrDN2)
				defer connDN2.Close()
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				stream, err := cDN2.Savechunk(ctx)
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}

				for i := uint64(0); i < chunks2; i++ {
					stream.Send(&pb3.Book{Name: books[i + chunks1].GetName(), Part: books[i + chunks1].GetPart(), Chunk: books[i + chunks1].GetChunk()})
				}

				r, err2 := stream.CloseAndRecv()
				if err2 != nil {
					log.Fatalf("Fallaste en enviar a 2 :c: %s %s", err2, r)
				}
				cancel()
				//SHOW_DETAILS log.Printf("%s", r.GetMsg())
			} else {
				//SHOW_DETAILS log.Printf("El DN2 no ta")
			}
		} else {
			if chunks3 > 0 {
				//SHOW_DETAILS log.Printf("Enviando a 3")
				cDN3, connDN3 = connect(addrDN3)
				defer connDN3.Close()
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				stream, err := cDN3.Savechunk(ctx)
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}

				for i := uint64(0); i < chunks3; i++ {
					stream.Send(&pb3.Book{Name: books[i + chunks1 + chunks2].GetName(), Part: books[i + chunks1 + chunks2].GetPart(), Chunk: books[i + chunks2].GetChunk()})
				}

				r, err2 := stream.CloseAndRecv()
				if err2 != nil {
					log.Fatalf("Fallaste en enviar 3 :c: %s %s", err2, r)
				}
				cancel()
				//SHOW_DETAILS log.Printf("%s", r.GetMsg())
			} else {
				//SHOW_DETAILS log.Printf("El DN3 no ta")
			}
		}
		
		
	}
}

/*Abre puerto para cliente*/
func OpenPortToClient(port string) {
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
/*Abre puerto para NameNode*/
func OpenPortToNameNode(port string) {
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
/*Abre puerto para DataNode*/
func OpenPortToDataNode(port string) {
	//SHOW_DETAILS log.Printf("Preparando... %s", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Falle al escuchar: %v", err)
	}
	//SHOW_DETAILS log.Printf("Escuchando...")
	s := grpc.NewServer()
	pb3.RegisterGreeterServer(s, &server3{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falle al servir: %v", err)
	}
}

//permite conexion con servidor
func toServer(addr string) (pb2.GreeterClient, *grpc.ClientConn){
	//SHOW_DETAILS log.Printf("Haciendo conexion %s", addr)
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se conecto: %v", err)
	}
	c := pb2.NewGreeterClient(conn)
	//SHOW_DETAILS log.Printf("logre conexion %s", addr)
	
	return c, conn
}

//permite conexion con datanode
func connect(addr string) (pb3.GreeterClient, *grpc.ClientConn){
	//SHOW_DETAILS log.Printf("Haciendo conexion %s", addr)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se conecto: %v", err)
	}
	c := pb3.NewGreeterClient(conn)
	//SHOW_DETAILS log.Printf("logre conexion %s", addr)
	
	return c, conn
}
//se identifica el datanode y se elige el algoritmo
func main(){
	log.Printf("Dime quien soy porfa:")
	fmt.Scanf("%d", &quienSoy)
	log.Printf("Soy centralizado o distribuido:")
	log.Printf("[1] Centralizado")
	log.Printf("[2] Distribuido")
	fmt.Scanf("%d", &distOrCentr)
	
	cName, connName = toServer(addrNN);
	defer connName.Close()
	if quienSoy == 1 {
		log.Printf("I'm 1")
		go OpenPortToClient(portDN1)
		go OpenPortToDataNode(portDN1_2)
		go OpenPortToNameNode(portDN1_3)
	} else if quienSoy == 2 {
		log.Printf("I'm 2")
		go OpenPortToClient(portDN2)
		go OpenPortToDataNode(portDN2_2)
		go OpenPortToNameNode(portDN2_3)
	} else {
		log.Printf("I'm 3")
		go OpenPortToClient(portDN3)
		go OpenPortToDataNode(portDN3_2)
		go OpenPortToNameNode(portDN3_3)
	}

	matame := 0
	fmt.Scanf("%d", &matame)
}