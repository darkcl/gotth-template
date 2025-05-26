package landing

import (
	"gotth/common"
	"net/http"
)

type LandingController struct {
	common.Controller
}

func NewLandingController() *LandingController {
	return &LandingController{}
}

func (l *LandingController) SetupRoute(router *http.ServeMux) {
	router.HandleFunc("GET /", l.LandingPage)
}

func (l *LandingController) LandingPage(w http.ResponseWriter, r *http.Request) {
	page := LandingPage()
	w.WriteHeader(http.StatusOK)
	page.Render(r.Context(), w)
}
