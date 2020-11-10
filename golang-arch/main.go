package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string
}

func main() {

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.HandleFunc("/base64", base64test)
	//http.HandleFunc("/hash", hashpassword)
	http.ListenAndServe(":8082", nil)

}

// curl localhost:8082/encode
func foo(w http.ResponseWriter, r *http.Request) {
	sp := []person{{First: "Jhon Encode"}, {First: "Jhon2 Encode"}}
	err := json.NewEncoder(w).Encode(sp)

	if err != nil {
		log.Println("Encoded bad data!", err)
	}

}

//  curl -XGET -H "content-type: application/json" -d '[{"First": "Jhon Decode"},{"First": "Jhon Decode2"}]'  'localhost:8082/decode'
func bar(w http.ResponseWriter, r *http.Request) {
	var p []person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Println("Decoded bad data!!", err)
	}

	log.Println("Person ", p)
}

func base64test(w http.ResponseWriter, r *http.Request) {
	msg := "Hello, 世界"
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	log.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	log.Println(string(decoded))
}

func hashpassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating bcrypt hash from password: %w", err)
	}
	return bs, nil
}

func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid Password: %w", err)
	}
	return nil
}
