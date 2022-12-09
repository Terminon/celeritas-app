package middleware

import (
	"myapp/data"

	"github.com/Terminon/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
