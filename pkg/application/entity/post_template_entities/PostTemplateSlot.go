package post_template_entities

type PostTemplateSlot struct {
	ID       int    `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	ParentId int    `json:"parent_id" gorm:"column:parent_id"`
	TenantId int    `json:"tenant_id" gorm:"column:tenant_id"`
}
