package api

import (
	twilio "github.com/carlosdp/twiliogo"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type (
	Api struct {
		config Config
		client twilio.Client
	}

	Config struct {
		userPh           string `json:"userph"`
		TwilioAccountSid string `json:"twilioAccountSid"`
		TwilioAuthToken  string `json:"twilioAuthToken"`
	}

	varsHandler func(http.ResponseWriter, *http.Request, map[string]string)
)

func InitApi(cfg Config, tc twilio.Client) *Api {

	return &Api{
		config: cfg,
		client: tc,
	}
}

func (a *Api) SetHandlers(prefix string, rtr *mux.Router) {
	rtr.Handle("/sms/send/{userid}", varsHandler(a.SendMsg)).Methods("POST")
	rtr.Handle("/sms/load/{userid}", varsHandler(a.LoadMsgs)).Methods("POST")
	rtr.Handle("/calc/bolus/{userid}", varsHandler(a.CalcBolus)).Methods("GET")
}

func (h varsHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	h(res, req, vars)
}

func (a *Api) insertMessages(messages interface{}) {

}

func (a *Api) SendMsg(res http.ResponseWriter, req *http.Request, vars map[string]string) {

	if vars["userid"] != "" {

		message, err := twilio.NewMessage(a.client, "6666666666", "5555555555", twilio.Body("TestBody"))

		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			log.Printf("yay sent!! %v ", message)
			res.WriteHeader(http.StatusOK)
			return
		}
	}
	res.WriteHeader(http.StatusBadRequest)
	return
}

func (a *Api) LoadMsgs(res http.ResponseWriter, req *http.Request, vars map[string]string) {

	if vars["userid"] != "" {

		messages, err := twilio.GetMessageList(a.client)

		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			log.Printf("yay loaded!! %v ", messages)
			res.WriteHeader(http.StatusCreated)
			return
		}
	}
	res.WriteHeader(http.StatusBadRequest)
	return
}

func (a *Api) CalcBolus(res http.ResponseWriter, req *http.Request, vars map[string]string) {

	res.WriteHeader(http.StatusOK)
	return
}
