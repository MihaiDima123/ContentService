package post_template_entities

type PostTemplateSlotProperty struct {
	ID                 int
	name               string
	value              string
	propertyTypeId     int
	postTemplateSlotId int
	tenantId           int
}
