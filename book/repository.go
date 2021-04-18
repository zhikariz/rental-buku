package book

import "gorm.io/gorm"

type Repository interface {
	Save(book Book) (Book, error)
	Update(book Book) (Book, error)
	FindById(id int) (Book, error)
	FindAll() ([]Book, error)
	Delete(book Book) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(book Book) (Book, error) {
	err := r.db.Create(&book).Error

	if err != nil {
		return book, err
	}
	return book, nil
}

func (r *repository) Update(book Book) (Book, error) {
	err := r.db.Save(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (r *repository) FindById(id int) (Book, error) {
	var book Book
	err := r.db.Where("id = ?", id).Find(&book).Error

	if err != nil {
		return book, err
	}
	return book, nil
}

func (r *repository) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	if err != nil {
		return books, err
	}
	return books, nil
}

func (r *repository) Delete(book Book) (bool, error) {
	err := r.db.Delete(&book).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
