package run

import "net/http"

type app struct {
	httpSrv *http.Server
}

func newApp() app {
	return app{}
}
