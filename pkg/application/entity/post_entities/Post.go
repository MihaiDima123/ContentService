package post_entities

type Post struct {
	ID             int64  `json:"id" gorm:"column:id"`
	UserId         int    `json:"userId"  gorm:"column:user_id"`
	UserEmail      string `json:"userEmail"  gorm:"column:user_email"`
	CategoryId     int    `json:"categoryId"  gorm:"column:category_id"`
	PostTemplateId int    `json:"postTemplateId" gorm:"column:post_template_id"`
	TenantId       int    `json:"tenantId" gorm:"column:tenant_id"`
}
