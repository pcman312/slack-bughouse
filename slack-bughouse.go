package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type SlackMessage struct {
	Response_type string // always "in_channel"
	Text          string // The generated list of players
}

func test(rw http.ResponseWriter, req *http.Request) {

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
		/*
					{
			            "response_type": "in_channel",
			            "text": "It's 80 degrees right now.",
			            "attachments": [
			                {
			                    "text":"Partly cloudy today and tomorrow"
			                }
			            ]
			        }
		*/

		results := "Team 1 White: " + names[0] + "\nTeam 1 Black: " + names[1] + "\nTeam 2 White: " + names[2] + "\nTeam 2 Black: " + names[3]

		message := SlackMessage{
			Response_type: "in_channel",
			Text:          results,
		}

		b, _ := json.Marshal(message)

		//fmt.Println(string(b))

		//json := "{\n'response_type': 'in_channel',\n'text': 'Its game time!',\n'attachments': [\n{\n'text':'" + results + "'\n}\n]\n}"
		//            Team 1 White: " + names[0] + "\nTeam 1 Black: " + names[1] + "\nTeam 2 White: " + names[2] + "\nTeam 2 Black: " + names[3]
		io.WriteString(rw, string(b))
	}

}

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
