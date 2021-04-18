package book

type BookFormatter struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Year        int    `json:"year"`
	CategoryID  uint   `json:"category_id"`
}

func FormatBook(book Book) (bookResponse BookFormatter) {
	bookResponse = BookFormatter{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
		Year:        book.Year,
		CategoryID:  book.CategoryID,
	}
	return
}
