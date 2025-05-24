package landing

import (
	"net/http"
)

type LandingController struct {
}

func NewLandingController(router *http.ServeMux) *LandingController {
	controller := &LandingController{}
	router.HandleFunc("GET /", controller.LandingPage)
	return controller
}

func (l *LandingController) LandingPage(w http.ResponseWriter, r *http.Request) {
	page := LandingPage()
	w.WriteHeader(http.StatusOK)
	page.Render(r.Context(), w)
}
