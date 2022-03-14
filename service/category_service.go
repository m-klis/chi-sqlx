package service

import (
	"context"
	"rest_api/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) (*web.CategoryResponse, error)
	Update(ctx context.Context, request web.CategoryUpdateRequest) (*web.CategoryResponse, error)
	FindById(ctx context.Context, categoryId int) (*web.CategoryResponse, error)
	Delete(ctx context.Context, categoryId int) (string, error)
	// FindAll(ctx context.Context, request web.GetParamRequest, wg *sync.WaitGroup) ([]*web.CategoryResponse, error)
	FindData(ctx context.Context, request web.GetParamRequest) ([]*web.CategoryResponse, *web.PaginateMetaData, error)
}
