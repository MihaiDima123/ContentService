package postTemplateSlotPropertyTypeDto

type PostTemplateSlotPropertyTypeDTO struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	TenantId int    `json:"tenantId"`
}

func New() *PostTemplateSlotPropertyTypeDTO {
	return new(PostTemplateSlotPropertyTypeDTO)
}
