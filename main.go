package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "something went wrong", http.StatusBadRequest) // equivalent to the two lines below
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("something went wrong"))
			return
		}

		fmt.Fprintf(rw, "Hello %s", d)
	})
	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe("127.0.0.1:5000", nil)
}
