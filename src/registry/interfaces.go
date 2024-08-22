package interfaces

import "net/http"

type Controller interface {
	SetupRoutes(handler *http.ServeMux)
}
