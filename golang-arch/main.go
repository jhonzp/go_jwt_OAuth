package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	// p1 := person{
	// 	First: "Jhon",
	// }

	// p2 := person{
	// 	First: "Rocio",
	// }

	// xp := []person{p1, p2}
	// bs, err := json.Marshal(xp)

	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println("Print Json:", string(bs))

	// xp2 := []person{}
	// err = json.Unmarshal(bs, &xp2)

	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println("back into go data Structure", xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8082", nil)

}

// curl localhost:8082/encode
func foo(w http.ResponseWriter, r *http.Request) {
	p := person{First: "Jhon Encode"}
	err := json.NewEncoder(w).Encode(p)

	if err != nil {
		log.Println("Encoded bad data!", err)
	}

}

//  curl -XGET -H "content-type: application/json" -d '{"First": "Jhon Decode"}'  'localhost:8082/encode'
func bar(w http.ResponseWriter, r *http.Request) {
	var p person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Println("Decoded bad data!!", err)
	}

	log.Println("Person ", p)
}
