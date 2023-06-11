package interfaces

import (
	postDto "contentservice/pkg/application/dto/post"
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/interfaces/restful"
)

type ContentService interface {
	restful.Service[post_entities.Post, postDto.PostDTO]
	Configure(repository ContentRepository)
}
