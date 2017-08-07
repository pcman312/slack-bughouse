package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type SlackMessage struct {
	Username string
	text     string
	mrkdwn   bool
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

		results := "Team 1 White: " + names[0] + "\nTeam 1 Black: " + names[1] + "\nTeam 2 White: " + names[2] + "\nTeam 2 Black: " + names[3]
		io.WriteString(rw, results)
	}

}

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
