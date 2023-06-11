package interfaces

import (
	categoryDto "contentservice/pkg/application/dto/category"
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/interfaces/restful"
)

type CategoryService interface {
	restful.Service[post_entities.Category, categoryDto.CategoryDTO]
}
