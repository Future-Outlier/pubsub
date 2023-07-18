package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
			return
		}

		log.Println("Received message: " + string(body))
		newBody := &Person{}
		err = proto.Unmarshal(body, newBody)
		log.Println("Received name: " + string(newBody.Name))
		log.Println("Received age: ", newBody.Age)
		log.Println("Body message len:", len(body), "string len:", strconv.Itoa(len(string(body))))
		// log.Println("Body message len : "+len(body), "string len: "+len(string(body)))
	})

	log.Fatal(http.ListenAndServe(":8900", nil))
}
