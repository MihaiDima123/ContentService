package repository

import (
	"contentservice/pkg/interfaces/restful"
)

type ContentRepository interface {
	Test()
	Configure(configuration restful.RepositoryConfiguration)
}
