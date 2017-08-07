package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type test_struct struct {
	Test string
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
	fmt.Println("text:", req.Form["text"])
	names := strings.Split(req.Form["text"], " ")
	fmt.Println(names)

}

func randomizer(names string) {

}

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
