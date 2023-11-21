package api

import "net/http"

type api struct {
	addr string
	r    *http.ServeMux
}

func New(addr string, r *http.ServeMux) *api {
	return &api{addr: addr, r: r}
}
func (api *api) FillEndpoints() {
	api.r.HandleFunc("/api/figureMax", figureMax)
	api.r.HandleFunc("/api/figureSum", figureSum)
}
func (api *api) ListenAndServe() error {
	return http.ListenAndServe(api.addr, api.r)
}
