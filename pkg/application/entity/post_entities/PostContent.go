package post_entities

type PostContent struct {
	ID                 int    `json:"id" gorm:"column:id"`
	Content            string `json:"content" gorm:"column:content"`
	PostId             int    `json:"postId" gorm:"column:post_id"`
	PostTemplateSlotId int    `json:"post_template_slot_id" gorm:"column:post_template_slot_id"`
	TenantId           int    `json:"tenant_id" gorm:"column:tenant_id"`
}
