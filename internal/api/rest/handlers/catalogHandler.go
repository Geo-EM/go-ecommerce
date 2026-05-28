package handlers

import (
	"e-commerce/internal/api/rest"
	"e-commerce/internal/api/rest/response"
	"e-commerce/internal/dto/catalogDto"
	"e-commerce/internal/repository"
	"e-commerce/internal/service"
	"e-commerce/internal/utils"

	"github.com/gofiber/fiber/v3"
)

type CatalogHandler struct {
	catalogService service.CatalogService
}

func SetupCatalogRoutes(rh *rest.RestHandler) {
	app := rh.App

	handler := CatalogHandler{
		catalogService: service.CatalogService{
			CatalogRepo:  repository.NewCatalogRepository(rh.DB),
			TokenService: *rh.TokenService,
			AppConfig:    *rh.AppConfig,
		},
	}

	routes := app.Group("/catalog")
	// Public endpoints
	// Products
	// routes.Get("/products", handler.getProducts)
	// routes.Get("/products/:id", handler.getProductById)

	// Categories
	routes.Get("/categories", handler.getCategories)
	routes.Get("/categories/:id", handler.getCategoryById)

	sellerRoutes := routes.Group("/seller", rh.TokenService.AuthorizeSeller)
	// Private endpoints
	// Products
	// sellerRoutes.Post("/products", handler.createProduct)
	// sellerRoutes.Patch("/products/:id", handler.updateProductById)
	// sellerRoutes.Delete("/products/:id", handler.deleteProductById)

	// Categories
	sellerRoutes.Post("/categories", handler.createCategory)
	sellerRoutes.Patch("/categories/:id", handler.updateCategory)
	sellerRoutes.Delete("/categories/:id", handler.deleteCategory)

}

// Products

// func (ch CatalogHandler) getProducts(ctx fiber.Ctx) error {
// 	products, err := ch.catalogService.GetProducts()
// 	if err != nil {
// 		return response.BadRequest(ctx, err.Error())
// 	}

// 	return response.OK(ctx, fiber.Map{"products": products}, "Fetched successfully")
// }

// func (ch CatalogHandler) getProductById(ctx fiber.Ctx) error {
// 	id := ctx.Params("id")

// 	product, err := ch.catalogService.GetProductById(id)
// 	if err != nil {
// 		return response.BadRequest(ctx, err.Error())
// 	}

// 	return response.OK(ctx, fiber.Map{"product": product}, "Fetched successfully")
// }

// func (ch CatalogHandler) createProduct(ctx fiber.Ctx) error {
// 	userClaims, err := uh.userService.TokenService.GetCurrentUser(ctx)
// 	if err != nil {
// 		return response.Unauthorized(ctx, err.Error())
// 	}

// 	input := catalogDto.CreateProductDto{}

// 	if err := ctx.Bind().Body(&input); err != nil {
// 		return response.BadRequest(ctx, "Invalid input")
// 	}

// 	product, err := ch.catalogService.CreateProduct(input)
// 	if err != nil {
// 		return response.BadRequest(ctx, err.Error())
// 	}

// 	return response.Created(ctx, fiber.Map{"product": product}, "Created successfully")
// }

// func (ch CatalogHandler) updateProductById(ctx fiber.Ctx) error {
// 	input := catalogDto.UpdateProductDto{}

// 	if err := ctx.Bind().Body(&input); err != nil {
// 		return response.BadRequest(ctx, "Invalid input")
// 	}

// 	product, err := ch.catalogService.UpdateProductById(input)
// 	if err != nil {
// 		return response.BadRequest(ctx, err.Error())
// 	}

// 	return response.OK(ctx, fiber.Map{"product": product}, "Updated successfully")
// }

// func (ch CatalogHandler) deleteProductById(ctx fiber.Ctx) error {
// 	id := ctx.Params("id")

// 	if err := ch.catalogService.DeleteProductById(id); err != nil {
// 		return response.BadRequest(ctx, err.Error())
// 	}

// 	return response.OK(ctx, nil, "Deleted successfully")
// }

// Categories

func (ch CatalogHandler) getCategories(ctx fiber.Ctx) error {
	categories, err := ch.catalogService.GetAllCategories()
	if err != nil {
		return response.NotFound(ctx, "Categories not found")
	}

	return response.OK(ctx, fiber.Map{"categories": categories}, "Fetched successfully")
}

func (ch CatalogHandler) getCategoryById(ctx fiber.Ctx) error {
	idParam := ctx.Params("id")

	id, err := utils.ParseStringToUint(idParam)
	if err != nil {
		return response.NotFound(ctx, "Category not found")
	}

	category, err := ch.catalogService.GetCategoryById(id)
	if err != nil {
		return response.NotFound(ctx, "Category not found")
	}

	return response.OK(ctx, fiber.Map{"category": category}, "Fetched successfully")
}

func (ch CatalogHandler) createCategory(ctx fiber.Ctx) error {
	input := catalogDto.CreateCategoryDto{}

	if err := ctx.Bind().Body(&input); err != nil {
		return response.BadRequest(ctx, "Invalid input")
	}

	category, err := ch.catalogService.CreateCategory(input)
	if err != nil {
		return response.BadRequest(ctx, err.Error())
	}

	return response.Created(ctx, fiber.Map{"category": category}, "Created successfully")
}

func (ch CatalogHandler) updateCategory(ctx fiber.Ctx) error {
	idParam := ctx.Params("id")

	id, err := utils.ParseStringToUint(idParam)
	if err != nil {
		return response.NotFound(ctx, "Category not found")
	}

	input := catalogDto.UpdateCategoryDto{}
	if err := ctx.Bind().Body(&input); err != nil {
		return response.BadRequest(ctx, "Invalid input")
	}

	category, err := ch.catalogService.UpdateCategory(id, input)
	if err != nil {
		return response.BadRequest(ctx, err.Error())
	}

	return response.OK(ctx, fiber.Map{"category": category}, "Updated successfully")
}

func (ch CatalogHandler) deleteCategory(ctx fiber.Ctx) error {
	idParam := ctx.Params("id")

	id, err := utils.ParseStringToUint(idParam)
	if err != nil {
		return response.NotFound(ctx, "Category not found")
	}

	if err := ch.catalogService.DeleteCategory(id); err != nil {
		return response.BadRequest(ctx, err.Error())
	}

	return response.OK(ctx, nil, "Deleted successfully")
}
