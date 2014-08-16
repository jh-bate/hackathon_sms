package api

import (
	"encoding/json"
	twilio "github.com/carlosdp/twiliogo"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	NO_PARAMS   = map[string]string{}
	FAKE_CONFIG = Config{}
	//basics setup
	rtr    = mux.NewRouter()
	tc     = new(twilio.MockClient)
	dinojr = InitApi(FAKE_CONFIG, tc)
	//mocked twilio msg
	testMessage = twilio.Message{
		Sid:         "testsid",
		DateCreated: "2013-05-11",
		DateUpdated: "2013-05-11",
		DateSent:    "2013-05-11",
		AccountSid:  "AC3TestAccount",
		From:        "+15555555555",
		To:          "+16666666666",
		Body:        "TestBody",
		NumSegments: "1",
		Status:      "queued",
		Direction:   "outbound-api",
		Price:       "4",
		PriceUnit:   "dollars",
		ApiVersion:  "2008-04-01",
		Uri:         "/2010-04-01/Accounts/AC3TestAccount/Messages/testsid.json",
	}
	testMessages = twilio.MessageList{
		Messages: []twilio.Message{testMessage},
	}
)

func TestSendMsg_StatusOK(t *testing.T) {

	dinojr.SetHandlers("", rtr)
	//server login
	request, _ := http.NewRequest("POST", "/", nil)
	response := httptest.NewRecorder()
	//mock
	messageJson, _ := json.Marshal(testMessage)
	params := url.Values{}
	params.Set("From", "6666666666")
	params.Set("To", "5555555555")
	params.Set("Body", "TestBody")
	tc.On("post", params, tc.RootUrl()+"/Messages.json").Return(messageJson, nil)

	//api call
	dinojr.SendMsg(response, request, map[string]string{"userid": "1234"})

	if response.Code != http.StatusOK {
		t.Fatalf("Non-expected status [%v]  expected [%v]", response.Code, http.StatusOK)
	}
}

func TestSendMsg_StatusBadRequest_NoUserIdParam(t *testing.T) {

	dinojr.SetHandlers("", rtr)
	//server login
	request, _ := http.NewRequest("POST", "/", nil)
	response := httptest.NewRecorder()

	dinojr.SendMsg(response, request, NO_PARAMS)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("Non-expected status [%v]  expected [%v]", response.Code, http.StatusBadRequest)
	}
}

func TestLoadMsgs_StatusOK(t *testing.T) {

	dinojr.SetHandlers("", rtr)
	//server login
	request, _ := http.NewRequest("POST", "/", nil)
	response := httptest.NewRecorder()

	//mock
	messagesJson, _ := json.Marshal(testMessages)

	tc.On("get", url.Values{}, tc.RootUrl()+"/SMS/Messages.json").Return(messagesJson, nil)

	dinojr.LoadMsgs(response, request, map[string]string{"userid": "1234"})

	if response.Code != http.StatusCreated {
		t.Fatalf("Non-expected status [%v]  expected [%v]", response.Code, http.StatusCreated)
	}
}

func TestLoadMsgs_StatusBadRequest_NoUserIdParam(t *testing.T) {

	dinojr.SetHandlers("", rtr)
	//server login
	request, _ := http.NewRequest("POST", "/", nil)
	response := httptest.NewRecorder()

	dinojr.LoadMsgs(response, request, NO_PARAMS)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("Non-expected status [%v]  expected [%v]", response.Code, http.StatusBadRequest)
	}
}

func TestCalcBolus_StatusOK(t *testing.T) {

	dinojr.SetHandlers("", rtr)
	//server login
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	dinojr.CalcBolus(response, request, map[string]string{"userid": "1234"})

	if response.Code != http.StatusOK {
		t.Fatalf("Non-expected status [%v]  expected [%v]", response.Code, http.StatusOK)
	}
}
