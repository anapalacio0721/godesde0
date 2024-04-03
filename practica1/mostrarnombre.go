package practica1

func NameComplete(firstname string, lastname string) (string, string) {
	fullname := firstname + " " + lastname
	nombreCompleto := "Hola " + fullname
	return fullname, nombreCompleto
}
