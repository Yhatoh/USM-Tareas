package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	pb "tarea1/tunel"

	"google.golang.org/grpc"
)

const (
	address = "10.6.40.240:9009"
)

//Structs
type Products struct {
	value                                int32
	id, product, name, dest, prioritario string
}

//Global
var c pb.GreeterClient

//Core function
//Recibe el tipo de tienda, el nombre de esta y la frecuencia en la cual se desean hacer los pedidos. Lee y envia los productos como pedidos a logistica
func order(types int, name string, frec int) {
	var file string
	if types == 1 {
		file = "pymes2.csv"
	} else {
		file = "retail.csv"
	}
	f, err := os.Open(file)
	if err != nil {
		log.Printf("No se puedo abrir archivo: %v", err)
	}
	defer f.Close()
	csvR := csv.NewReader(f)
	var prods []Products
	for {
		row, err := csvR.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		if name == row[3] {
			if types == 1 {
				value, _ := strconv.Atoi(row[2])
				prods = append(prods, Products{int32(value),
					row[0],
					row[1],
					row[3],
					row[4],
					row[5]})
			} else {
				value, _ := strconv.Atoi(row[2])
				prods = append(prods, Products{int32(value),
					row[0],
					row[1],
					row[3],
					row[4],
					"r"})
			}
		}
	}
	for {
		time.Sleep(time.Duration(frec) * time.Second)

		for _, prod := range prods {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			if types == 1 {
				r, err := c.Order(ctx, &pb.Product{Id: prod.id, Value: prod.value, Product: prod.product, Type: "pyme", Name: prod.name, Dest: prod.dest, Prioritario: prod.prioritario})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
				log.Printf("El producto tiene el siguiente id de seguimiento pa que lo sigas: %d", r.GetId())
			} else {
				r, err := c.Order(ctx, &pb.Product{Id: prod.id, Value: prod.value, Product: prod.product, Type: "retail", Name: prod.name, Dest: prod.dest, Prioritario: prod.prioritario})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
				log.Printf("El producto tiene el siguiente id de seguimiento pa que lo sigas: %d", r.GetId())
			}
		}
	}
}

//Recibe el tipo de la tienda y el nombre. Muestra el estado del producto
func tracing(types int, name string) {
	for {
		action, idS := askAction()
		if action == 2 {
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := c.Trace(ctx, &pb.Tracing{Id: int32(idS), Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("El producto tiene el siguiente estado: %s", r.GetEstado())

	}
	log.Printf("Hasta la proxima!")
}

//Prints Function
func askAction() (action, idS int) {
	log.Printf("Que quieres hacer?")
	log.Printf("[1] Seguimiento")
	log.Printf("[2] Salir")
	fmt.Scanf("%d", &action)
	if action == 2 {
		idS = -1
		return
	}
	log.Printf("Dame id de producto a seguir")
	fmt.Scanf("%d", &idS)
	return
}

//Recibe las direcciones de memoria de el tipo, frecuencia y nombre. Recibe por input los datos y los asigna respectivamente
func initStore() (types int, frec int, name string) {
	log.Printf("Eres Pyme o Retail?")
	log.Printf("[1] Pyme")
	log.Printf("[2] Retail")
	fmt.Scanf("%d", &types)
	log.Printf("Como se llama tu tienda?")
	fmt.Scanf("%s", &name)
	log.Printf("Cada cuanto quieres tu pedido?")
	fmt.Scanf("%d", &frec)
	return
}

func main() {
	types, frec, name := initStore()

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se conecto: %v", err)
	}
	defer conn.Close()
	c = pb.NewGreeterClient(conn)
	if types == 1 {
		go order(types, name, frec)
		tracing(types, name)
	} else {
		log.Printf("Lo siento eres retail no puedes hacer seguimiendo, saludos")
		order(types, name, frec)
	}
}
