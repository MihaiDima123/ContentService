package postTemplateSlotPropertyDto

import postTemplateSlotPropertyTypeDto "contentservice/pkg/application/dto/post-template-slot-property-type"

type PostTemplateSlotPropertyDTO struct {
	Id           int                                                             `json:"id"`
	Name         string                                                          `json:"name"`
	Value        string                                                          `json:"value"`
	PropertyType postTemplateSlotPropertyTypeDto.PostTemplateSlotPropertyTypeDTO `json:"propertyType"`
}

func New() *PostTemplateSlotPropertyDTO {
	return new(PostTemplateSlotPropertyDTO)
}
