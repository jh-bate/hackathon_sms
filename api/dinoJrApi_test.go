package api

import (
	"./../clients"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	NO_PARAMS = map[string]string{}
	//basics setup
	rtr             = mux.NewRouter()
	mockplatform, _ = clients.NewMockClient("usrname", "pw")
	api             = InitApi(mockplatform)
)

func TestGetBolus(t *testing.T) {

	req, _ := http.NewRequest("GET", "/bolus", nil)
	res := httptest.NewRecorder()

	api.SetHandlers("", rtr)

	/*
	 * No userid given
	 */
	api.GetBolus(res, req, NO_PARAMS)

	if res.Code != http.StatusBadRequest {
		t.Fatalf("Resp given [%s] expected [%s] ", res.Code, http.StatusBadRequest)
	}

	body, _ := ioutil.ReadAll(res.Body)

	if string(body) != STATUS_NO_USR_DETAILS {
		t.Fatalf("Message given [%s] expected [%s] ", string(body), STATUS_NO_USR_DETAILS)
	}

}

func TestGetIOB(t *testing.T) {

	req, _ := http.NewRequest("GET", "/iob", nil)
	res := httptest.NewRecorder()

	api.SetHandlers("", rtr)

	/*
	 * No userid given
	 */
	api.GetIOB(res, req, NO_PARAMS)

	if res.Code != http.StatusBadRequest {
		t.Fatalf("Resp given [%s] expected [%s] ", res.Code, http.StatusBadRequest)
	}

	body, _ := ioutil.ReadAll(res.Body)

	if string(body) != STATUS_NO_USR_DETAILS {
		t.Fatalf("Message given [%s] expected [%s] ", string(body), STATUS_NO_USR_DETAILS)
	}
}
