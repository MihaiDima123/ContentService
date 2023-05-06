package repository

import (
	"contentservice/application/modules/core"
)

type ContentRepository interface {
	Test()
	Configure(configuration core.RepositoryConfiguration)
}
