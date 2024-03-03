package tickets

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	madrugada = "madrugada"
	manana    = "manana"
	tarde     = "tarde"
	noche     = "noche"
)

type Ticket struct {
	Id          int
	Name        string
	Email       string
	Destination string
	Hour        string
	Price       float64
}

// Slice con total de tickets diarios
var TicketsAll []Ticket

// Funcion lectura archivo, almacenado de informacion y retorno de total
func GetFileData() int {

	fileName := "./tickets.csv"
	//Manejo panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	//Lectura archivo
	res, err := os.ReadFile(fileName)
	if err != nil {
		panic("\n ALERTA! El archivo indicado no fue encontrado o está dañado: " + fileName + " No podran realizarse los calculos.\n")
	}
	//Extraccion de datos y almacenamiento en slice de estructuras Ticket
	data := strings.Split(string(res), "\n")
	for i := 0; i < len(data)-1; i++ {
		ticketData := strings.Split(string(data[i]), ",")
		id, err := strconv.Atoi(ticketData[0])
		if err != nil {
			fmt.Printf("Error al intentar parsear el id del ticket con email: %s Revisar ticket \n", ticketData[2])
		}
		price, err := strconv.ParseFloat(ticketData[5], 64)
		if err != nil {
			fmt.Printf("Error al intentar parsear el precio del ticket con email: %s Revisar ticket \n", ticketData[2])
		}
		ticket := Ticket{
			Id:          id,
			Name:        ticketData[1],
			Email:       ticketData[2],
			Destination: ticketData[3],
			Hour:        ticketData[4],
			Price:       price,
		}
		TicketsAll = append(TicketsAll, ticket)
	}
	totalTickets := len(TicketsAll)
	return totalTickets
}

// Funcion calculo total de tickets por destino
func GetTotalTickets(destination string) (int, error) {
	var total int
	for _, ticket := range TicketsAll {
		if ticket.Destination == destination {
			total += 1
		}
	}
	if total == 0 {
		return 0, errors.New("No hay pasajes al destino indicado o El destino indicado es invalido.")
	}
	return total, nil
}

// Funcion derivadora para calculo tickets por horario vuelo
func GetCountByPeriod(time string) (int, error) {
	switch time {
	case madrugada:
		total, err := CountByPeriod(0, 6)
		return total, err
	case manana:
		total, err := CountByPeriod(7, 12)
		return total, err
	case tarde:
		total, err := CountByPeriod(13, 19)
		return total, err
	case noche:
		total, err := CountByPeriod(20, 23)
		return total, err
	default:
		return 0, errors.New("El periodo indicado no existe.")
	}

}

// Funcion calculo tickets por horario vuelo
func CountByPeriod(min int, max int) (int, error) {
	var total int
	for _, ticket := range TicketsAll {
		hourString := strings.Split(string(ticket.Hour), ":")
		hour, err := strconv.Atoi(hourString[0])
		if err != nil {
			return 0, errors.New("Error al intentar parsear la hora del ticket con email: " + ticket.Email + ". Revisar ticket")
		}
		if hour >= min && hour <= max {
			total += 1
		}
	}
	return total, nil
}

// Funcion calculo porcentaje tickets destino respecto a total tickets
// (Mantuve la firma indicada en el enunciado, sin embargo no seria necesario recibir el total de tickets por parametro.
// Se podria acceder directo a length de TotalTickets dentro de esta funcion)
func AverageDestination(destination string, total int) (float64, error) {
	totalDestination, err := GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}
	if total == 0 {
		return 0, errors.New("Error en el numero de pasajes totales. Valor: 0")
	}
	percentage := float64(totalDestination) / float64(total) * 100
	return percentage, nil
}
