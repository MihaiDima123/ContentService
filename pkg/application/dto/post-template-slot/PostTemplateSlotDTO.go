package postTemplateSlotDto

import postTemplateSlotPropertyDto "contentservice/pkg/application/dto/post-template-slot-property"

type PostTemplateSlotDTO struct {
	Id         int                                                       `json:"id"`
	Name       string                                                    `json:"name"`
	Properties []postTemplateSlotPropertyDto.PostTemplateSlotPropertyDTO `json:"properties"`
	Children   []PostTemplateSlotDTO                                     `json:"children"`
	TenantId   int                                                       `json:"tenantId"`
}

func New() *PostTemplateSlotDTO {
	return new(PostTemplateSlotDTO)
}
