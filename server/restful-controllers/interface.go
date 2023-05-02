package restful_controllers

import (
	"github.com/jackc/pgx"
)

type Options struct {
	Connection pgx.ConnPool
}

type RestController interface {
	Configure(options Options)
}
