package post_template_entities

type PostTemplate struct {
	ID       int    `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	TenantId int    `json:"tenant_id" gorm:"column:tenant_id"`
}
