package core

import (
	"github.com/jackc/pgx"
)

type Options struct {
	Connection pgx.ConnPool
}

type RestController interface {
}
