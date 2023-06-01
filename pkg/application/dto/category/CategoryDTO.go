package categoryDto

type CategoryDTO struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	TenantId string `json:"tenantId"`
}

func New() *CategoryDTO {
	return new(CategoryDTO)
}
