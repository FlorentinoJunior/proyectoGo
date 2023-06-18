package entities

type Student struct {
	Name string
}

func NewStudent(name string) *Student {
	return &Student{
		Name: name,
	}
}

func (s *Student) BorrowBook(b *Book) {
	b.Borrow()
}

func (s *Student) ReturnBook(b *Book) {
	b.Return()
}
