package interfaces

import (
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/interfaces/customerrors"
	"contentservice/pkg/interfaces/restful"
)

type CategoryRepository interface {
	restful.Repository[post_entities.Category]
	GetCountByNameAndTenant(name string, tenantId int) (count int64, dbError customerrors.DbError)
}
