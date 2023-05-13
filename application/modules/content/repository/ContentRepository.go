package repository

import (
	"contentservice/application/core"
)

type ContentRepository interface {
	Test()
	Configure(configuration core.RepositoryConfiguration)
}
