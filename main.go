package main

import (
	"fmt"
	"muller-florencia/internal/tickets"
)

func main() {

	//Manejo de errores y panic
	//goroutines
	//estructuras

	//total, err := tickets.GetTotalTickets("Brazil")
	//arrayData := tickets.GetFileData()
	//fmt.Println(arrayData[0])
	destination := "Japan"
	total, err := tickets.GetTotalTickets(destination)
	if err != nil {
		fmt.Printf("Error al calcular el total de tickets por destino.")
	}
	fmt.Printf("Total de tickets para %s es %d \n", destination, total)
}
