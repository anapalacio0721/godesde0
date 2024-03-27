package main

import (
	"fmt"

	ejercicios "github.com/anapalacio0721/godesde0/ejercicios"
)

func main() {

	/*estado, texto := variables.ConviertoaTexto1(45)
	fmt.Println(estado)
	fmt.Println(texto)*/
	/*
		if os := runtime.GOOS; os == "Linux." || os == "OS X" {
			fmt.Printf("Esto no es Windows")
		} else {
			fmt.Println("Esto es Windows")

		}

		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Esto es Linux")
		case "darwin":
			fmt.Println("Esto es Darwin")
		default:
			fmt.Printf("%s \n", os)
		}
	*/
	numero, texto := ejercicios.ConvNumerico("500")
	fmt.Println(numero)
	fmt.Println(texto)

}
