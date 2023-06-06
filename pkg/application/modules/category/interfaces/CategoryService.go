package interfaces

import (
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/interfaces/restful"
)

type CategoryService interface {
	restful.Service[post_entities.Category]
}
