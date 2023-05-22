package interfaces

import (
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/interfaces/restful"
)

type ContentRepository interface {
	restful.Repository[post_entities.Post]
}
