package main

import (
	"fmt"
	"sync"
	"github.com/FlorentinoJunior/proyectoGo.git/entities"
)

func main() {
	// Crear estudiantes y libros
	estudiante1 := entities.NewStudent("Estudiante 1")
	estudiante2 := entities.NewStudent("Estudiante 2")
	libro1 := entities.NewBook("Libro 1", "Autor 1")
	libro2 := entities.NewBook("Libro 2", "Autor 2")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		// Estudiante 1 toma prestado el libro 1
		estudiante1.BorrowBook(libro1)
		fmt.Printf("%s ha tomado prestado %s\n", estudiante1.Name, libro1.Title)

		// Estudiante 1 devuelve el libro 1
		estudiante1.ReturnBook(libro1)
		fmt.Printf("%s ha devuelto %s\n", estudiante1.Name, libro1.Title)
	}()

	go func() {
		defer wg.Done()
		// Estudiante 2 toma prestado el libro 2
		estudiante2.BorrowBook(libro2)
		fmt.Printf("%s ha tomado prestado %s\n", estudiante2.Name, libro2.Title)

		// Estudiante 2 devuelve el libro 2
		estudiante2.ReturnBook(libro2)
		fmt.Printf("%s ha devuelto %s\n", estudiante2.Name, libro2.Title)
	}()

	wg.Wait()
}
