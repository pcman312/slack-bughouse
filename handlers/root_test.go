package handlers

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleRoot(t *testing.T) {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/?text=player1%20player2%20player3%20player4", strings.NewReader(""))
	HandleRoot(recorder, req)

	rawBody, err := ioutil.ReadAll(recorder.Body)
	Ok(t, err)

	if string(rawBody) != "OK" {
		t.Fatalf("body expected to be [OK] but was %s", rawBody)
	}
}
