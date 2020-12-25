package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/streadway/amqp"
)

type Complete struct {
	Id       string `json:"id"`
	Tipo     string `json:"tipo"`
	Valor    string `json:"valor"`
	Origen   string `json:"origen"`
	Destino  string `json:"destino"`
	Intentos string `json:"intentos"`
	Date     string `json:"date"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var total float64
var gastos float64
var ingresos float64

//Lee el mensaje de la cola de RabbitMQ y guarda los datos pertinentes en el archivo "data.csv"
func main() {
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

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	total = 0.0
	gastos = 0.0
	ingresos = 0.0
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		log.Printf("\nIngresos: %f\nPerdidas: %f\nTotal: %f\n", ingresos, gastos, total)
		os.Exit(1)
	}()

	go func() {
		data, err := os.OpenFile("data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		a := 0.0
		b := 0.0

		w := csv.NewWriter(data)
		defer data.Close()
		for d := range msgs {

			var sli Complete
			balance := 0.0
			entregado := "Si"

			json.Unmarshal(d.Body, &sli)

			if sli.Date == "0" {
				entregado = "No"
				if sli.Tipo == "prioritario" {
					a, _ = strconv.ParseFloat(sli.Valor, 64)
					b, _ = strconv.ParseFloat(sli.Intentos, 64)
					balance = balance + 0.3*a - (10*b - 10)
					gastos += (10*b - 10)
					ingresos += 0.3 * a
				} else if sli.Tipo == "retail" {
					a, _ = strconv.ParseFloat(sli.Valor, 64)
					b, _ = strconv.ParseFloat(sli.Intentos, 64)
					balance = balance + a - (10*b - 10)
					gastos += (10*b - 10)
					ingresos += a
				} else {
					b, _ = strconv.ParseFloat(sli.Intentos, 64)
					balance = balance - (10*b - 10)
					gastos += (10*b - 10)

				}
			} else {
				a, _ = strconv.ParseFloat(sli.Valor, 64)
				b, _ = strconv.ParseFloat(sli.Intentos, 64)
				balance = balance + a - (10*b - 10)
				gastos += (10*b - 10)
				ingresos += a
			}

			total = ingresos - gastos
			wr := make([]string, 6)
			wr[0] = sli.Id
			wr[1] = sli.Tipo
			wr[2] = sli.Valor
			wr[3] = sli.Intentos
			wr[4] = entregado
			wr[5] = fmt.Sprintf("%f", balance)

			_ = w.Write(wr)
			w.Flush()

		}
	}()
	log.Printf(" [*] Esperando mensajes. Para salir presione CTRL+C")
	<-forever

}
