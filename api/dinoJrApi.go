package api

import (
	"./../clients"
	"github.com/gorilla/mux"
	"net/http"
)

type (
	Api struct {
		platform clients.Platform
	}

	varsHandler func(http.ResponseWriter, *http.Request, map[string]string)
)

const (
	TP_SESSION_TOKEN      = "x-tidepool-session-token"
	STATUS_NO_USR_DETAILS = "No user id was given"
)

func InitApi(pf clients.Platform) *Api {
	return &Api{
		platform: pf,
	}
}

func (a *Api) SetHandlers(prefix string, rtr *mux.Router) {
	rtr.Handle("/bolus/{userid}", varsHandler(a.GetBolus)).Methods("GET")
	rtr.Handle("/iob/{userid}", varsHandler(a.GetIOB)).Methods("GET")
}

func (h varsHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	h(res, req, vars)
}

func (a *Api) GetBolus(res http.ResponseWriter, req *http.Request, vars map[string]string) {

	id := vars["userid"]
	if id == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(STATUS_NO_USR_DETAILS))
		return
	}
	res.WriteHeader(http.StatusNotImplemented)
	return
}

func (a *Api) GetIOB(res http.ResponseWriter, req *http.Request, vars map[string]string) {

	id := vars["userid"]
	if id == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(STATUS_NO_USR_DETAILS))
		return
	}
	res.WriteHeader(http.StatusNotImplemented)
	return
}
