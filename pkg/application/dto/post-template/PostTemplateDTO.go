package postTemplateDto

import postTemplateSlotDto "contentservice/pkg/application/dto/post-template-slot"

type PostTemplateDTO struct {
	Id       int                                       `json:"id"`
	Name     string                                    `json:"name"`
	Slots    []postTemplateSlotDto.PostTemplateSlotDTO `json:"slots"`
	TenantId int                                       `json:"tenantId"`
}

func New() *PostTemplateDTO {
	return new(PostTemplateDTO)
}
