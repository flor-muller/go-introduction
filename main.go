package main

import (
	"fmt"
	"muller-florencia/internal/tickets"
)

func main() {

	destination := "Brazil"
	hour := "noche"

	totalTickets := tickets.GetFileData()
	fmt.Println("El total de tickets vendidos del dia es:", totalTickets)

	destintationCh := make(chan int)
	hourCh := make(chan int)
	percentageCh := make(chan float64)

	go func(destination string) {
		totalDestination, err := tickets.GetTotalTickets(destination)
		if err != nil {
			fmt.Println("El siguiente ERROR ha ocurrido al calcular el total de tickets por destino.", err)
		}
		destintationCh <- totalDestination
	}(destination)

	go func(hour string) {
		totalPeriod, err := tickets.GetCountByPeriod(hour)
		if err != nil {
			fmt.Println("El siguiente ERROR ha ocurrido al calcular el total de tickets por periodo:", err)
		}
		hourCh <- totalPeriod
	}(hour)

	go func(destination string) {
		percentage, err := tickets.AverageDestination(destination, totalTickets)
		if err != nil {
			fmt.Println("El siguiente ERROR ha ocurrido al calcular el porcentaje de tickets por destino:", err)
		}
		percentageCh <- percentage
	}(destination)

	fmt.Printf("El total de tickets vendidos con destino a %s es: %d \n", destination, <-destintationCh)
	fmt.Printf("El total de tickets vendidos con hora de vuelo en el turno de %s es: %d \n", hour, <-hourCh)
	fmt.Printf("El porcentaje de tickets vendidos con destino a %s es de: %.1f%% \n", destination, <-percentageCh)
}
