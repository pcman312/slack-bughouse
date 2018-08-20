package handlers

import (
	"io"
	"net/http"
)

func HandleRoot(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "OK")
}
