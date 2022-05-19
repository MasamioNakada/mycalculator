package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Función que nos permite leer la entrada del usuario y devolver el valor que ha introducido
func ReadInput() string {
	// Lee la entrada de un usaurio
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// Devuelve el valor introducido por el usuario
	return scanner.Text()
}

// Función que nos permite recibir un string, transformarlo en un int y devolverlo
func parsear(input string) int {
	// Transformamos los strings en valores de tipo int
	operator, err := strconv.Atoi(input)
	if err != nil { // Manejando el error al parsear los datos de string a int
		fmt.Println(err)
	}

	// Devuelve en valor de tipo int
	return operator
}

/*
	STRUCT
*/
type calc struct{}

// Metodo del STRUCT
func (calc) operate(operators string, symbol string) int {
	// Convierte en una lista los valores que reciba en "operators" separados por el valor "symbol". Ejm: Recibe 2+2, lo transforma en [2 2]
	inputSplit := strings.Split(operators, symbol)

	// Parsea los strigns de la lista para recibir int
	operator1 := parsear(inputSplit[0])
	operator2 := parsear(inputSplit[1])

	// Evalúa el símbolo recibido para realizar la operación correspondiente
	switch symbol {
	case "+":
		// SUMA
		return operator1 + operator2
	case "-":
		// RESTA
		return operator1 - operator2
	case "*":
		// Multiplicación
		return operator1 * operator2
	case "/":
		// División
		return operator1 / operator2
	default:
		fmt.Println("El operador indicado no es correcto o no está soportado") // Muestra el mensaje por defecto
		return 0
	}
}
