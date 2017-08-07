package main

import (
	"fmt"
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

	/*/ Parse form values
	fmt.Println("token:", req.Form["token"])
	fmt.Println("team_id:", req.Form["team_id"])
	fmt.Println("team_domain:", req.Form["team_domain"])
	fmt.Println("user_name:", req.Form["user_name"])
	fmt.Println("text:", req.Form["text"])
	*/
	fmt.Println(req.Form["text"][0])
	fmt.Println("text:", req.Form["text"])
	names := strings.Split(req.Form["text"][0], " ")
	fmt.Printf("%T", names)
	fmt.Println(names)

	if len(names) < 4 {
		fmt.Println("There are not enough players!")

	} else {
		//fmt.Println("Let the games begin!")
		t := time.Now()
		rand.Seed(int64(t.Nanosecond())) // no shuffling without this line

		for i := len(names) - 1; i > 0; i-- {
			j := rand.Intn(i)
			names[i], names[j] = names[j], names[i]
		}
		fmt.Println("Team 1 White: ", names[1], "\nTeam 1 Black: ", names[2], "\nTeam 2 Black: ", names[3], "\nTeam 2 Black: ", names[3])
	}

}

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
