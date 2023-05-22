package interfaces

import (
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/interfaces/restful"
)

type ContentService interface {
	restful.Service[post_entities.Post]
	Configure(repository ContentRepository)
}
