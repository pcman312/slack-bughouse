package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func test(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	req.ParseForm()

	names := strings.Split(req.Form["text"][0], " ")

	if len(names) < 4 {
		fmt.Println("There are not enough players!")

	} else {
		//  Let the games begin!
		t := time.Now()
		rand.Seed(int64(t.Nanosecond())) // no shuffling without this line

		for i := len(names) - 1; i > 0; i-- {
			j := rand.Intn(i)
			names[i], names[j] = names[j], names[i]
		}

		// build json to send back to slack

		message := `{
        "response_type": "in_channel",
        "text": "*Would you like to play a game?*",
        "mrkdwn": true,
        "attachments": [
            {
                "text": "{{.}}",
                "mrkdwn_in": ["text"]
            }
            ]
        }`

		results := "*Team 1 White:* _" + names[0] + "_\n*Team 1 Black*: _" + names[1] + "_\n\n*Team 2 White:* _" + names[2] + "_\n*Team 2 Black:* _" + names[3] + "_"

		tmpl, err := template.New("message").Parse(message)
		if err != nil {
			panic(err)
		}

		var tpl bytes.Buffer

		err = tmpl.Execute(&tpl, results)
		if err != nil {
			panic(err)
		}

		io.WriteString(rw, tpl.String())
	}

}

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
