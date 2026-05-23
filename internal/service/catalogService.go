package service

import (
	"e-commerce/config"
	"e-commerce/internal/auth"
	"e-commerce/internal/repository"
)

type CatalogService struct {
	CatalogRepo  repository.CatalogRepository
	TokenService auth.TokenService
	AppConfig    config.AppConfig
}
