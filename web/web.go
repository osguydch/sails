package web

import (
	"github.com/Unknwon/macaron"

	"github.com/containerops/sails/middleware"
	"github.com/containerops/sails/router"
)

func SetSailsMacaron(m *macaron.Macaron) {
	//Setting Middleware
	middleware.SetMiddlewares(m)
	//Setting Router
	router.SetRouters(m)
}
