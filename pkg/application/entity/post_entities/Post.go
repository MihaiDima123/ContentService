package post_entities

type Post struct {
	ID             int64
	UserId         int
	UserEmail      string
	CategoryId     int
	PostTemplateId int
	TenantId       int
}
