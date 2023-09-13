package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Zapato struct {
	marca  string
	precio float64
	talla  int
}

func newZapato() (Zapato, error) {
	inputHandler := bufio.NewReader(os.Stdin)
	fmt.Print("Digite la marca del zapato: ")
	marca, err := inputHandler.ReadString('\n')
	if err != nil {
		return Zapato{}, err
	}
	marca = strings.TrimSuffix(marca, "\n")

	fmt.Print("Digite el precio del zapato: ")
	var precio float64
	_, err = fmt.Scanf("%f", &precio)
	if err != nil {
		return Zapato{}, err
	}

	fmt.Print("Digite la talla del zapato: ")

	var talla int
	_, err = fmt.Scan(&talla)
	if err != nil {
		return Zapato{}, err
	}

	if tallaDisponible(talla) {
		return Zapato{}, errors.New("Talla no disponible")
	}

	return Zapato{marca: marca, precio: precio, talla: talla}, nil
}

func tallaDisponible(talla int) bool {
	return (talla <= 44 || talla >= 34)
}

func main() {
	zapatos := []Zapato{}
	for i := 0; i < 5; i++ {
		zapato, err := newZapato()
		if err != nil {
			i--
			fmt.Println(err)
			continue // NOTE: código de mierda, ni lo voy a probar
		}
		zapatos = append(zapatos, zapato)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(zapatos[i])
	}
}
