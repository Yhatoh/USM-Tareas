package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	//"strconv"
	pb "tarea1/tunel"
	pb2 "tarea1/tunel2"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

const (
	portClient = ":9009"
	portCamion = ":9008"
)

//Structs

//structus tipo server
type server struct {
	pb.UnimplementedGreeterServer
}

type server2 struct {
	pb2.UnimplementedGreeterServer
}

//structs para manejo de productos
type Products struct {
	value, idT                                        int32
	id, product, name, dest, prioritario, date, state string
}

type SafeQueue struct {
	queue []Products
	mux   sync.Mutex
}

type SafeMap struct {
	states map[int32]Products
	mux    sync.Mutex
}

//struct para el uso de rabbitMQ
type Complete struct {
	Id       string `json:"id"`
	Tipo     string `json:"tipo"`
	Valor    string `json:"valor"`
	Origen   string `json:"origen"`
	Destino  string `json:"destino"`
	Intentos string `json:"intentos"`
	Date     string `json:"date"`
}

//Global
var idTrace int32 = 0
var retail, priori, norm SafeQueue
var states SafeMap

//Recibe un producto que haya sido pedido y lo guarda en
//registro logistica.csv
func writeFile(pack Products) {
	data, _ := os.OpenFile("logistica.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	w := csv.NewWriter(data)
	defer data.Close()
	wr := make([]string, 8)
	wr[0] = pack.date
	wr[1] = strconv.Itoa(int(pack.idT))
	if pack.prioritario == "0" {
		wr[2] = "normal"
	} else if pack.prioritario == "1" {
		wr[2] = "prioritario"
	} else {
		wr[2] = "retail"
	}
	wr[3] = pack.product
	wr[4] = strconv.Itoa(int(pack.value))
	wr[5] = pack.name
	wr[6] = pack.dest
	wr[7] = strconv.Itoa(int(pack.idT))

	_ = w.Write(wr)
	w.Flush()
	return
}

//inicializa el csv ingresando los campos
func initCSV() {
	data, _ := os.OpenFile("logistica.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	w := csv.NewWriter(data)
	defer data.Close()
	wr := make([]string, 8)
	wr[0] = "timestamp"
	wr[1] = "id-paquete"
	wr[2] = "tipo"
	wr[3] = "nombre"
	wr[4] = "valor"
	wr[5] = "origen"
	wr[6] = "destino"
	wr[7] = "seguimiento"

	_ = w.Write(wr)
	w.Flush()
}

//Core function

//Order permite comunicacion entre cliente y logistica
//cuando el cliente le manda un paquete este lo recibe
//lo almacena y entrega id de seguimiento
func (s *server) Order(ctx context.Context, in *pb.Product) (*pb.Tracing, error) {
	log.Printf("Compra de %s %s %d %s %s %s %s", in.GetId(), in.GetProduct(), in.GetValue(), in.GetType(), in.GetName(), in.GetDest(), in.GetPrioritario())
	idTrace = idTrace + 1

	log.Printf("Toma id de seguimiento %d", idTrace)
	go storeInWarehouse(in, idTrace)
	return &pb.Tracing{Id: idTrace,
		Name: in.GetName()}, nil
}

//Trace permite comunicacion entre cliente y logistica
//cuando el cliente le mande una ide de seguimiento
//este le entrega el estado del paquete correspondiente
//a la id
func (s *server) Trace(ctx context.Context, in *pb.Tracing) (*pb.State, error) {
	log.Printf("Mmmmm, la tienda %s le pregunta a los camiones respecto al seguimiento: %d", in.GetName(), in.GetId())
	states.mux.Lock()
	defer states.mux.Unlock()

	return &pb.State{Estado: states.states[in.GetId()].state}, nil
}

//GimmePackeT permite comunicacion entre logistica y camion
//Cuando un camion le dice que esta esperando, logistica
//revisa si hay un paquete disponible para ese camion
//si es que lo hay se lo entrega si no le dice le entrega -2
//para decir que no hay paquete, ademas aprovechad de cambiar
//el estado del paquete entregado de paso
func (s2 *server2) GimmePacket(ctx context.Context, in *pb2.Truck) (*pb2.Packet, error) {
	var prod1 Products

	if in.GetTypes() == 1 {
		if isEmpty(retail.queue) && isEmpty(priori.queue) {
			return &pb2.Packet{Id: -2}, nil
		} else if !isEmpty(retail.queue) {
			retail.mux.Lock()
			prod1 = front(retail.queue)
			retail.queue = dequeue(retail.queue)
			defer retail.mux.Unlock()
		} else {
			priori.mux.Lock()
			prod1 = front(priori.queue)
			priori.queue = dequeue(priori.queue)
			defer priori.mux.Unlock()
		}
	} else {
		if isEmpty(priori.queue) && isEmpty(norm.queue) {
			return &pb2.Packet{Id: -2}, nil
		} else if !isEmpty(priori.queue) {
			priori.mux.Lock()
			prod1 = front(priori.queue)
			priori.queue = dequeue(priori.queue)
			defer priori.mux.Unlock()
		} else {
			norm.mux.Lock()
			prod1 = front(norm.queue)
			norm.queue = dequeue(norm.queue)
			defer norm.mux.Unlock()
		}
	}
	log.Printf("Toma paquete")
	states.mux.Lock()
	states.states[prod1.idT] = Products{prod1.value,
		prod1.idT,
		prod1.id,
		prod1.product,
		prod1.name,
		prod1.dest,
		prod1.prioritario,
		prod1.date,
		"En camino"}
	states.mux.Unlock()
	return &pb2.Packet{Id: prod1.idT,
		Types:  prod1.prioritario,
		Value:  prod1.value,
		Origin: prod1.name,
		Dest:   prod1.dest,
		Trys:   "0",
		Date:   "0"}, nil
}

//ImBack permite comunicacion entre logistica y camion
//cuando un camion ya haya realizado una entrega este le
//dice a logistica la informacion respecto a su viaje con
//ese paquete
func (s2 *server2) ImBack(ctx context.Context, in *pb2.Packet) (*pb2.Tick, error) {
	states.mux.Lock()
	prod1 := states.states[in.Id]
	if in.GetDate() == "0" {
		states.states[prod1.idT] = Products{prod1.value,
			prod1.idT,
			prod1.id,
			prod1.product,
			prod1.name,
			prod1.dest,
			prod1.prioritario,
			in.GetDate(),
			"No entregado"}
	} else {
		states.states[prod1.idT] = Products{prod1.value,
			prod1.idT,
			prod1.id,
			prod1.product,
			prod1.name,
			prod1.dest,
			prod1.prioritario,
			in.GetDate(),
			"Entregado"}
	}
	atoi, _ := strconv.Atoi(in.GetTrys())
	go sendFianance(states.states[prod1.idT], atoi)
	states.mux.Unlock()
	return &pb2.Tick{Ok: 1}, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//Recibe un producto y la cantidad de intentos. Le entrega mediante rabbitMQ los datos del producto a finanzas.
func sendFianance(prod Products, intentos int) {
	sval := strconv.Itoa(int(prod.value))
	sint := strconv.Itoa(intentos)
	pio := "xd"
	if prod.prioritario == "0" {
		pio = "normal"
	} else if prod.prioritario == "1" {
		pio = "prioritario"
	} else {
		pio = "retail"
	}

	conn, err := amqp.Dial("amqp://test:test@10.6.40.240:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	body := Complete{Id: strconv.Itoa(int(prod.idT)), Tipo: pio, Valor: sval, Origen: "xd", Destino: "xd", Intentos: sint, Date: prod.date}
	bo, _ := json.Marshal(body)

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bo,
		})
	log.Printf("Ahi va un completado finanzas")
	failOnError(err, "Failed to publish a message")
}

//storeInWarehouse recibe un producto a almacenar, donde
//segun su tipo lo posiciona en alguna cola, y lo escribe
//en su registro
func storeInWarehouse(in *pb.Product, idTrace int32) {
	prod := Products{in.GetValue(),
		idTrace,
		in.GetId(),
		in.GetProduct(),
		in.GetName(),
		in.GetDest(),
		in.GetPrioritario(),
		time.Now().Format("2006-01-02 15:04:05"),
		"En bodega"}
	writeFile(prod)
	states.mux.Lock()
	states.states[idTrace] = prod
	defer states.mux.Unlock()
	if prod.prioritario == "0" {
		norm.mux.Lock()
		norm.queue = enqueue(norm.queue, prod)
		defer norm.mux.Unlock()
	} else if prod.prioritario == "1" {
		priori.mux.Lock()
		priori.queue = enqueue(priori.queue, prod)
		defer priori.mux.Unlock()
	} else {
		retail.mux.Lock()
		retail.queue = enqueue(retail.queue, prod)
		defer retail.mux.Unlock()
	}
	//log.Printf("retail: %v\n", retail)
	//log.Printf("priori: %v\n", priori)
	//log.Printf("normal: %v\n", norm)
}

//Queue function

//isEmpty permite sabe si la cola esta vacia
func isEmpty(queue []Products) bool {
	if len(queue) == 0 {
		return true
	} else {
		return false
	}
}

//dequeue saca el primer elemento de la cola
func dequeue(queue []Products) []Products {
	return queue[1:]
}

//enqueue mete un nuevo elemento a la cola
func enqueue(queue []Products, prod Products) []Products {
	return append(queue, prod)
}

//entrega el valor que esta al frente de la cola
func front(queue []Products) Products {
	return queue[0]
}

//Connection function
//permite conexion con cliente
func connClient() {
	lis, err := net.Listen("tcp", portClient)
	if err != nil {
		log.Fatalf("Falle al escuchar: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falle al servir: %v", err)
	}
}

//permite conexion con camion
func connCamion() {
	lis, err := net.Listen("tcp", portCamion)
	if err != nil {
		log.Fatalf("Falle al escuchar: %v", err)
	}
	s2 := grpc.NewServer()
	pb2.RegisterGreeterServer(s2, &server2{})
	if err := s2.Serve(lis); err != nil {
		log.Fatalf("Falle al servir: %v", err)
	}
}

func main() {
	initCSV()
	states.states = make(map[int32]Products)
	log.Printf("Iniciando Logistica\n")

	go connClient()
	go connCamion()
	for {

	}
}
