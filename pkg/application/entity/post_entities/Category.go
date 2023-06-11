package post_entities

import (
	categoryDto "contentservice/pkg/application/dto/category"
)

type Category struct {
	ID       int64  `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	TenantId int    `json:"tenant_id" gorm:"column:tenant_id"`
}

func NewCategory(dto *categoryDto.CategoryDTO) *Category {
	c := new(Category)
	c.Name = dto.Name
	c.TenantId = dto.TenantId
	return c
}

func ToCategoryDTO(category *Category) *categoryDto.CategoryDTO {
	c := new(categoryDto.CategoryDTO)
	c.Name = category.Name
	c.TenantId = category.TenantId
	return c
}
