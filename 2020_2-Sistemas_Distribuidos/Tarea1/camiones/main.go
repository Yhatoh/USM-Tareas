package main

import (
	"context"
	"encoding/csv"
	"os"
	"strconv"

	/*
		"os"
		"io"
		"encoding/csv"
	*/
	"time"
	/*
		"net"
	*/
	"fmt"
	"log"
	"math/rand"
	pb2 "tarea1/tunel2"

	"google.golang.org/grpc"
)

const (
	address = "10.6.40.240:9008"
)

//Recibe el producto y la id del camion que realizo el envio. 
///Escribe los datos del envio en el csv respectivo del camion
func writeHistory(id string, pack *pb2.Packet) {
	data, _ := os.OpenFile("camion"+id+".csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	w := csv.NewWriter(data)
	defer data.Close()
	wr := make([]string, 7)
	wr[0] = strconv.Itoa(int(pack.GetId()))
	if pack.GetTypes() == "0" {
		wr[1] = "normal"
	} else if pack.GetTypes() == "1" {
		wr[1] = "prioritario"
	} else {
		wr[1] = "retail"
	}

	wr[2] = strconv.Itoa(int(pack.GetValue()))
	wr[3] = pack.GetOrigin()
	wr[4] = pack.GetDest()
	wr[5] = pack.GetTrys()
	wr[6] = pack.GetDate()
	_ = w.Write(wr)
	w.Flush()
}

//Recibe la id de un camion. Genera su .csv
func initCSV(id string) {
	data, _ := os.OpenFile("camion"+id+".csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	w := csv.NewWriter(data)
	defer data.Close()
	wr := make([]string, 7)
	wr[0] = "id-paquete"
	wr[1] = "tipo"
	wr[2] = "valor"
	wr[3] = "origen"
	wr[4] = "destino"
	wr[5] = "intentos"
	wr[6] = "fecha-entrega"
	_ = w.Write(wr)
	w.Flush()
}

//Trucks
//Recibe el tiempo de espera, el tiempo de envio y la id 
//del camion. Espera, recibe e intenta realizar los pedidos.
func Truck1(wait, send, id int32) {
	for {
		var r, r2 *pb2.Packet
		var err2 error
		log.Printf("%s: dame producto\n", textTruck(id))

		for {
			conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("No se conecto: %v", err)
			}
			defer conn.Close()
			c := pb2.NewGreeterClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			r, err2 = c.GimmePacket(ctx, &pb2.Truck{Types: 1})

			if err2 != nil {
				log.Fatalf("could not greet: %v", err2)
			}
			if r.GetId() != -2 {
				break
			}
		}
		log.Printf("%s: recibi %v\n", textTruck(id), r)
		start := time.Now()
		log.Printf("%s: dame otro producto o me voy\n", textTruck(id))

		for {
			conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("No se conecto: %v", err)
			}
			defer conn.Close()
			c := pb2.NewGreeterClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			r2, err2 = c.GimmePacket(ctx, &pb2.Truck{Types: 1})

			if err2 != nil {
				log.Fatalf("could not greet: %v", err2)
			}
			if r2.GetId() != -2 {
				break
			}
			end := time.Now()
			diff := end.Sub(start).Seconds()
			if int32(diff) > wait {
				r2 = nil
				break
			}
		}
		if r2 == nil {
			log.Printf("%s: demoraron mucho adios\n", textTruck(id))
		} else {
			log.Printf("%s: recibi este otro %v\n", textTruck(id), r2)
		}

		travelling(r, r2, id, send)
	}
}

//normal
//Recibe el tiempo de espera, el tiempo de envio. Espera, 
//recibe e intenta realizar los pedidos.
func Truck2(wait, send int32) {
	for {
		var r, r2 *pb2.Packet
		var err2 error
		log.Printf("%s: dame producto\n", textTruck(0))

		for {
			conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("No se conecto: %v", err)
			}
			defer conn.Close()
			c := pb2.NewGreeterClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			r, err2 = c.GimmePacket(ctx, &pb2.Truck{Types: 2})

			if err2 != nil {
				log.Fatalf("could not greet: %v", err2)
			}
			if r.GetId() != -2 {
				break
			}
		}
		log.Printf("%s: recibi %v\n", textTruck(0), r)
		start := time.Now()
		log.Printf("%s: dame otro producto o me voy\n", textTruck(0))
		for {
			conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("No se conecto: %v", err)
			}
			defer conn.Close()
			c := pb2.NewGreeterClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			r2, err2 = c.GimmePacket(ctx, &pb2.Truck{Types: 2})

			if err2 != nil {
				log.Fatalf("could not greet: %v", err2)
			}
			if r2.GetId() != -2 {
				break
			}
			end := time.Now()
			diff := end.Sub(start).Seconds()
			if int32(diff) > wait {
				r2 = nil
				break
			}
		}
		if r2 == nil {
			log.Printf("%s: demoraron mucho adios\n", textTruck(0))
		} else {
			log.Printf("%s: recibi este otro %v\n", textTruck(0), r2)
		}

		travelling(r, r2, 0, send)
	}
}

//sendToLogistica recibe el paquete, si fue entregado o no
//y sus cantidad de intento, para mandarle esta informacion
//a logitistica, y aprovechando guardar esa informacion en 
//el registro del camion correspondiente
func sendToLogistica(pack *pb2.Packet, deli int32, trys int32, id int32) {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se conecto: %v", err)
	}
	defer conn.Close()
	c := pb2.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var dateIn string
	if deli == 0 {
		dateIn = "0"
	} else {
		dateIn = time.Now().Format("2006-01-02 15:04:05")
	}
	pack2 := &pb2.Packet{Id: pack.GetId(),
		Types:  pack.GetTypes(),
		Value:  pack.GetValue(),
		Origin: pack.GetOrigin(),
		Dest:   pack.GetDest(),
		Trys:   strconv.Itoa(int(trys)),
		Date:   dateIn}
	writeHistory(strconv.Itoa(int(id)), pack2)
	_, err3 := c.ImBack(ctx, pack2)
	if err3 != nil {
		log.Fatalf("could not greet: %v", err3)
	}
}

//travelling realiza la simulacion de entrega de paquetes
//segun el camion dado
func travelling(r *pb2.Packet, r2 *pb2.Packet, id int32, send int32) {

	if r2 == nil {
		log.Printf("%s: Viajando a entregar un paquete\n", textTruck(id))
		trys, deli := doDelivery(send, id, r)
		log.Printf("%s Volviendo a la central\n", textTruck(id))

		sendToLogistica(r, deli, trys, id)
	} else {
		if r.GetValue() > r2.GetValue() {
			log.Printf("%s voy a entregar %d ya que cuesta %d y el otro %d\n", textTruck(id), r.GetId(), r.GetValue(), r2.GetValue())
			trys, deli := doDelivery(send, id, r)
			log.Printf("%s ahora voy a intentar emtregar %d\n", textTruck(id), r2.GetId())
			trys2, deli2 := doDelivery(send, id, r2)
			log.Printf("%s Volviendo a la central\n", textTruck(id))
			sendToLogistica(r, deli, trys, id)

			sendToLogistica(r2, deli2, trys2, id)
		} else {
			log.Printf("%s: voy a entregar %d ya que cuesta %d y el otro %d\n", textTruck(id), r2.GetId(), r2.GetValue(), r.GetValue())
			trys, deli := doDelivery(send, id, r2)
			log.Printf("%s: ahora voy a intentar emtregar %d\n", textTruck(id), r.GetId())
			trys2, deli2 := doDelivery(send, id, r)
			log.Printf("%s: Volviendo a la central\n", textTruck(id))

			sendToLogistica(r2, deli, trys, id)

			sendToLogistica(r, deli2, trys2, id)
		}
	}
}

//doDelivery permite ver si se entrega un paquete o no
//intento 3 veces maximo, dependiendo de las condiciones
//del paquete
func doDelivery(send int32, id int32, pack *pb2.Packet) (int32, int32) {
	var trys int32 = 1

	time.Sleep(time.Duration(send) * time.Second)
	if try(id) {
		return trys, 1
	}

	if (pack.GetTypes() == "1" || pack.GetTypes() == "0") && pack.GetValue() < 10*trys {
		return trys, 0
	}
	trys = trys + 1
	if try(id) {
		return trys, 1
	}

	if (pack.GetTypes() == "1" || pack.GetTypes() == "0") && pack.GetValue() < 10*trys {
		return trys, 0
	}
	trys = trys + 1
	if try(id) {
		return trys, 1
	}

	return trys, 0
}

//try realiza 1 intento de entrega
func try(id int32) bool {
	prob := rand.Intn(100)

	if prob < 80 {
		log.Printf("%s: yo te entregue el paquete!\n", textTruck(id))
		return true
	} else {
		log.Printf("%s: NO HABIA NADA EN EL LUGAR FALLE ENTREGA!\n", textTruck(id))

		return false
	}
}

//Prints function
func initTruck() (wait, send int32) {
	log.Printf("Cuanto se demora un camion en entregar un paquete")
	fmt.Scanf("%d", &send)
	log.Printf("Cuanto espera un camion un segundo paquete")
	fmt.Scanf("%d", &wait)
	return
}

func textTruck(id int32) string {
	if id == 0 {
		return "Camion N"
	}
	return "Camion R " + strconv.Itoa(int(id))
}

func main() {
	initCSV("0")
	initCSV("1")
	initCSV("2")

	wait, send := initTruck()
	go Truck1(wait, send, 1)
	go Truck1(wait, send, 2)
	Truck2(wait, send)
}
