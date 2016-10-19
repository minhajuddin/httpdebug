package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println(string(bytes))
		fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")

		fmt.Fprintln(w, "<h1>HTTP Debug</h1>")
		fmt.Fprintln(w, "<pre>"+string(bytes)+"</pre>")

	})
	fmt.Println("Starting debug server at http://localhost:4444/ ..")
	http.ListenAndServe(":4444", nil)
}
