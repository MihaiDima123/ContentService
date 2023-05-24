package post_template_entities

type PostTemplateSlotProperty struct {
	ID                 int    `json:"id" gorm:"column:id"`
	Name               string `json:"name" gorm:"column:name"`
	Value              string `json:"value" gorm:"column:value"`
	PropertyTypeId     int    `json:"property_type_id" gorm:"column:property_type_id"`
	PostTemplateSlotId int    `json:"post_template_slot_id" gorm:"column:post_template_slot_id"`
	TenantId           int    `json:"tenant_id" gorm:"column:tenant_id"`
}
