package book

type Service interface {
	CreateBook(input CreateBookInput) (Book, error)
	UpdateBook(id int, input UpdateBookInput) (Book, error)
	DeleteBook(input BookUriInput) (bool, error)
	GetBookById(id int) (Book, error)
	GetBooks() ([]Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateBook(input CreateBookInput) (Book, error) {
	book := Book{
		Title:       input.Title,
		Description: input.Description,
		Author:      input.Author,
		Year:        input.Year,
		CategoryID:  input.CategoryID,
	}

	newBook, err := s.repository.Save(book)
	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (s *service) UpdateBook(id int, input UpdateBookInput) (Book, error) {
	book, err := s.repository.FindById(id)
	if err != nil {
		return book, err
	}

	book.Title = input.Title
	book.Description = input.Title
	book.Author = input.Author
	book.Year = input.Year
	book.CategoryID = input.CategoryID

	updatedBook, err := s.repository.Update(book)

	if err != nil {
		return updatedBook, err
	}

	return updatedBook, nil

}

func (s *service) DeleteBook(input BookUriInput) (bool, error) {
	book, err := s.repository.FindById(input.ID)

	if err != nil {
		return false, err
	}

	isDeleted, err := s.repository.Delete(book)

	if err != nil {
		return false, err
	}
	return isDeleted, nil
}

func (s *service) GetBookById(id int) (Book, error) {
	book, err := s.repository.FindById(id)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (s *service) GetBooks() ([]Book, error) {
	books, err := s.repository.FindAll()
	if err != nil {
		return books, err
	}
	return books, nil
}
