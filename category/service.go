package category

type Service interface {
	CreateCategory(input CreateCategoryInput) (Category, error)
	UpdateCategory(id int, input UpdateCategoryInput) (Category, error)
	DeleteCategory(input CategoryUriInput) (bool, error)
	GetCategoryById(id int) (Category, error)
	GetCategories() ([]Category, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateCategory(input CreateCategoryInput) (Category, error) {
	category := Category{
		Name: input.Name,
	}

	newCategory, err := s.repository.Save(category)

	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

func (s *service) UpdateCategory(id int, input UpdateCategoryInput) (Category, error) {
	category, err := s.repository.FindById(id)
	if err != nil {
		return category, err
	}

	category.Name = input.Name

	updatedCategory, err := s.repository.Update(category)

	if err != nil {
		return updatedCategory, err
	}
	return updatedCategory, nil

}

func (s *service) DeleteCategory(input CategoryUriInput) (bool, error) {
	category, err := s.repository.FindById(input.ID)

	if err != nil {
		return false, err
	}

	isDeleted, err := s.repository.Delete(category)

	if err != nil {
		return false, err
	}

	return isDeleted, nil

}

func (s *service) GetCategoryById(id int) (Category, error) {
	category, err := s.repository.FindById(id)
	if err != nil {
		return category, err
	}
	return category, nil

}

func (s *service) GetCategories() ([]Category, error) {
	categories, err := s.repository.FindAll()
	if err != nil {
		return categories, err
	}
	return categories, nil
}
