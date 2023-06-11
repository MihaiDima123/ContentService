package categoryDto

import (
	"contentservice/pkg/application/entity/post_entities"
)

type CategoryDTO struct {
	Name     string `json:"name"`
	TenantId int    `json:"tenantId"`
}

func New(category *post_entities.Category) *CategoryDTO {
	c := new(CategoryDTO)
	c.Name = category.Name
	c.TenantId = category.TenantId
	return c
}
