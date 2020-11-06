package main

import (
	"encoding/json"
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

func foo(w http.ResponseWriter, r *http.Request) {
	p := person{First: "Jhon Encode"}
	err := json.NewEncoder(w).Encode(p)

	if err != nil {
		log.Println("Encoded bad data!", err)
	}

}

func bar(w http.ResponseWriter, r *http.Request) {

}
