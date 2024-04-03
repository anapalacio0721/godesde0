package teclado

import (
	"bufio" //io imput y output para teclado
	"fmt"   //mostrar algo
	"os"    //sistema operativo
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumero() { //es un metodo cuando no rcibe nada
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(("ingrese numero 1: "))
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("El dato ingresado es incorrecto" + err.Error())
		}
	}
	fmt.Println(("ingrese numero 2: "))
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("El dato ingresado es incorrecto" + err.Error())
		}
	}
	fmt.Println(("ingrese leyenda: "))
	if scanner.Scan() {
		leyenda = scanner.Text()

	}

	fmt.Println(leyenda, numero1*numero2)
}
