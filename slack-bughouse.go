package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type test_struct struct {
	Test string
}

func test(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))

	if req.Method == "POST" {
		req.ParseForm()
		token := req.Form.Get("token")
		if token != "" {
			fmt.Println("token is:", token)
		} else {
			fmt.Println("I see no token")
		}
		// logic part of log in
		fmt.Println("token:", req.Form["token"])
		fmt.Println("team_id:", req.Form["team_id"])

		fmt.Println("team_domain:", req.Form["team_domain"])
		fmt.Println("user_name:", req.Form["user_name"])
		fmt.Println("text:", req.Form["text"])

	}

}

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
