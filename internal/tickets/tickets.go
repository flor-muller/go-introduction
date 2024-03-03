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

// id, nombre, email, país de destino, hora del vuelo y precio.
// 1,Tait Mc Caughan,tmc0@scribd.com,Finland,17:11,785

type Ticket struct {
	Id          int
	Name        string
	Email       string
	Destination string
	Hour        string
	Price       float64
}

// Funcion lectura archivo
func GetFileData() []string {
	fileName := "./tickets.csv"
	res, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al intentar leer el archivo: ", fileName)
	}
	data := strings.Split(string(res), "\n")
	return data
}

// Funcion Calculo total de tickets diarios (todos los destinos)
func GetTotal() {

}

// ejemplo 1: Funcion calculo total de tickets por destino
func GetTotalTickets(destination string) (int, error) {
	data := GetFileData()
	var total int
	for i := 0; i < len(data)-1; i++ {
		ticket := strings.Split(string(data[i]), ",")
		if ticket[3] == destination {
			total += 1
		}
	}
	return total, nil
}

// ejemplo 2
func CountByPeriod(min int64, max int64) (int, error) {
	data := GetFileData()
	var total int
	for i := 0; i < len(data)-1; i++ {
		ticket := strings.Split(string(data[i]), ",")
		hourString := strings.Split(string(ticket[4]), ":")
		hour, err := strconv.ParseInt(hourString[0], 10, 64)
		if err != nil {
			errors.New("Error al intentar parsear la hora.")
		}
		if hour >= min && hour <= max {
			total += 1
		}
	}
	return total, nil
}

// ejemplo 2 enunciado
func GetCountByPeriod(time string) (int, error) {
	switch time {
	case madrugada:
		//0 a 6
		total, err := CountByPeriod(0, 6)
		return total, err
	case manana:
		//7 a 12
		total, err := CountByPeriod(7, 12)
		return total, err
	case tarde:
		//13 a 19
		total, err := CountByPeriod(13, 19)
		return total, err
	case noche:
		//20 a 23
		total, err := CountByPeriod(20, 23)
		return total, err
	default:
		return 0, errors.New("El periodo indicado no existe.")
	}

}

// ejemplo 3: Funcion calculo porcentaje tickets destino respecto a total tickets
func AverageDestination(destination string, total int) (float64, error) {
	return 1, nil
}
