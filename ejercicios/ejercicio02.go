package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error

func TableMultiplicar() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(("ingrese numero 1: "))
		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El dato ingresado es incorrecto" + err.Error())
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i += 1 {
		multi := numero1 * i
		fmt.Printf("\n %d x %d = %d", numero1, i, multi)
	}

}
