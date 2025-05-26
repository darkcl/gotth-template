package common

import "net/http"

type Controller interface {
	SetupRoute(router *http.ServeMux)
}
