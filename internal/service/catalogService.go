package service

import (
	"e-commerce/config"
	"e-commerce/internal/auth"
	"e-commerce/internal/domain"
	"e-commerce/internal/dto/catalogDto"
	"e-commerce/internal/repository"
)

type CatalogService struct {
	CatalogRepo  repository.CatalogRepository
	TokenService auth.TokenService
	AppConfig    config.AppConfig
}

func (cs *CatalogService) CreateCategory(input catalogDto.CreateCategoryDto) (domain.Category, error) {

	category, err := cs.CatalogRepo.CreateCategory(domain.Category{
		Name:         input.Name,
		ParentId:     input.ParentId,
		ImageUrl:     input.ImageUrl,
		DisplayOrder: input.DisplayOrder,
	})
	if err != nil {
		return domain.Category{}, err
	}

	return category, nil
}

func (cs *CatalogService) UpdateCategory(cid uint, input catalogDto.UpdateCategoryDto) (domain.Category, error) {

	category, err := cs.CatalogRepo.UpdateCategory(cid, domain.Category{
		Name:         input.Name,
		ParentId:     input.ParentId,
		ImageUrl:     input.ImageUrl,
		DisplayOrder: input.DisplayOrder,
	})
	if err != nil {
		return domain.Category{}, err
	}

	return category, nil
}

func (cs *CatalogService) GetCategoryById(cid uint) (domain.Category, error) {
	return cs.CatalogRepo.FindCategoryById(cid)
}

func (cs *CatalogService) GetAllCategories() ([]domain.Category, error) {
	return cs.CatalogRepo.FindAllCategories()
}

func (cs *CatalogService) DeleteCategory(cid uint) error {
	return cs.CatalogRepo.DeleteCategory(cid)
}
