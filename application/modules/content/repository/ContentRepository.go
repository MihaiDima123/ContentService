package repository

import (
	"contentservice/application/modules"
)

type ContentRepository interface {
	Test()
	Configure(configuration modules.RepositoryConfiguration)
}
