package book

type CreateBookInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Year        int    `json:"year" binding:"required"`
	CategoryID  uint   `json:"category_id" binding:"required"`
}

type UpdateBookInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Year        int    `json:"year" binding:"required"`
	CategoryID  uint   `json:"category_id" binding:"required"`
}

type BookUriInput struct {
	ID int `uri:"id" binding:"required"`
}
