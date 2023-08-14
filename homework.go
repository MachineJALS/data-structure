package main

import (
	"bufio"
	"fmt"
	"os"
)

type zapato struct {
	marca  string
	precio float64
	talla  int
}

func new_zapato() zapato {
	inputHandler := bufio.NewReader(os.Stdin)
	fmt.Println("Digite la marca del zapato: ")
	marca, err := inputHandler.ReadString('\n')
	if err != nil {
		return zapato{}
	}
	fmt.Println("Digite el precio del zapato: ")
	precio, err := fmt.Scanf("%f", &precio)
	if err != nil {
		return zapato{}
	}

	fmt.Println("Digite la talla del zapato: ")
	talla := talla_disponible()

	z := zapato{marca: marca, precio: precio, talla: talla}
	return z
}

func talla_disponible() int {
	fmt.Println("Digite la talla del zapato: ")
	talla, _ := fmt.Scan()
	if talla <= 44 || talla >= 34 {
		return talla
	} else {
		talla_disponible()
	}
}

func main() {
	zapatos := [5]zapato{}
	for i := 0; i < 5; i++ {
		zapatos[i] = new_zapato()
	}
	for i := 0; i < 5; i++ {
		fmt.Println(zapatos[i])
	}
}
