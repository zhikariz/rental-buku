package category

type CreateCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

type CategoryUriInput struct {
	ID int `uri:"id" binding:"required"`
}
