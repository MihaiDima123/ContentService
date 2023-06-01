package postContentDto

import postTemplateSlotDto "contentservice/pkg/application/dto/post-template-slot"

type PostContentDTO struct {
	Id       int                                     `json:"id"`
	Content  string                                  `json:"content"`
	Slot     postTemplateSlotDto.PostTemplateSlotDTO `json:"slot"`
	TenantId int                                     `json:"tenantId"`
}

func New() *PostContentDTO {
	return new(PostContentDTO)
}
