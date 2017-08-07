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
	fmt.Println(req.Form["text"][0])
	fmt.Println("text:", req.Form["text"])
	names := strings.Split(req.Form["text"][0], " ")
	fmt.Printf("%T", names)
	fmt.Println(names)

	if len(names) < 4 {
		fmt.Println("There are not enough players!")

	} else {
		fmt.Println("Let the games begin!")

	}

}

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
