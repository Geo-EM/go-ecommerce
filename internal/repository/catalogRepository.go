package repository

import (
	"e-commerce/internal/domain"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CatalogRepository interface {
	CreateCategory(c domain.Category) (domain.Category, error)
	UpdateCategory(cid uint, c domain.Category) (domain.Category, error)

	FindCategoryById(cid uint) (domain.Category, error)
	FindAllCategories() ([]domain.Category, error)

	DeleteCategory(cid uint) error
}

type catalogRepository struct {
	db *gorm.DB
}

func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &catalogRepository{db: db}
}

//

func (catalogRepo catalogRepository) CreateCategory(input domain.Category) (domain.Category, error) {
	var newCategory domain.Category = input

	err := catalogRepo.db.Create(&newCategory).Error
	if err != nil {
		return domain.Category{}, errors.New("Failed to create category")
	}

	return newCategory, nil
}

func (catalogRepo catalogRepository) UpdateCategory(catId uint, input domain.Category) (domain.Category, error) {
	var existingCategory domain.Category
	err := catalogRepo.db.Model(&existingCategory).Clauses(clause.Returning{}).Where("id = ?", catId).Updates(input).Error
	if err != nil {
		return domain.Category{}, errors.New("Failed to update category")
	}

	return existingCategory, nil
}

func (catalogRepo catalogRepository) FindCategoryById(catId uint) (domain.Category, error) {
	var category domain.Category
	err := catalogRepo.db.First(&category, catId).Error
	if err != nil {
		return domain.Category{}, errors.New("Category not found")
	}

	return category, nil
}

func (catalogRepo catalogRepository) FindCategoryByEmail(email string) (domain.Category, error) {
	var category domain.Category
	err := catalogRepo.db.First(&category, "email = ?", email).Error
	if err != nil {
		return domain.Category{}, errors.New("Category not found")
	}

	return category, nil
}

func (catalogRepo catalogRepository) FindAllCategories() ([]domain.Category, error) {
	var categories []domain.Category
	err := catalogRepo.db.Find(&categories).Error
	if err != nil {
		return []domain.Category{}, errors.New("Failed to fetch categories")
	}

	return categories, nil
}

func (catalogRepo catalogRepository) DeleteCategory(catId uint) error {
	err := catalogRepo.db.Delete(&domain.Category{}, catId).Error
	if err != nil {
		return errors.New("Failed to delete category")
	}

	return nil
}
