package entities

type Book struct {
	Title  string
	Author string
	Estado string
}

func NewBook(title string, author string) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Estado: "Disponible",
	}
}

func (b *Book) Borrow() {
	b.Estado = "Prestado"
}

func (b *Book) Return() {
	b.Estado = "Disponible"
}
