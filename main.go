package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/FlorentinoJunior/proyectoGo.git/entities"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Recoge el nombre del estudiante
	fmt.Print("Ingrese el nombre del estudiante: ")
	nombreEstudiante, _ := reader.ReadString('\n')

	// Recoge el nombre y autor del libro
	fmt.Print("Ingrese el título del libro: ")
	tituloLibro, _ := reader.ReadString('\n')

	fmt.Print("Ingrese el autor del libro: ")
	autorLibro, _ := reader.ReadString('\n')

	// Crea un nuevo estudiante
	estudiante1 := entities.NewStudent(nombreEstudiante)

	// Crea un nuevo libro
	libro1 := entities.NewBook(tituloLibro, autorLibro)

	// El estudiante toma prestado el libro
	estudiante1.BorrowBook(libro1)

	// Imprime el estado del libro después de ser prestado
	fmt.Printf("Estado del libro después de ser prestado: %s\n", libro1.Estado)

	// El estudiante devuelve el libro
	estudiante1.ReturnBook(libro1)

	// Imprime el estado del libro después de ser devuelto
	fmt.Printf("Estado del libro después de ser devuelto: %s\n", libro1.Estado)
}
