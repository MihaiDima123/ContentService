package postDto

import (
	postContentDto "contentservice/pkg/application/dto/post-content"
)

type PostDTO struct {
	UserId       int                             `json:"userId"`
	UserEmail    string                          `json:"userEmail"`
	Contents     []postContentDto.PostContentDTO `json:"contents"`
	TemplateName string                          `json:"templateName"`
	CategoryName string                          `json:"categoryName"`
	TenantId     int                             `json:"tenantId"`
}

func New() *PostDTO {
	return new(PostDTO)
}
