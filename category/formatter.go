package category

type CategoryFormatter struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func FormatCategory(category Category) (categoryResponse CategoryFormatter) {
	categoryResponse = CategoryFormatter{
		ID:   category.ID,
		Name: category.Name,
	}
	return
}
