package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handles() {

	log.Println("My learning web server starting @ http://localhost:8080/intro")

	r := http.NewServeMux()
	r.HandleFunc("/api/", handleAPI)
	r.HandleFunc("/about", handleAbout)
	r.HandleFunc("/careers", handleCareers)
	r.HandleFunc("/intro", handleIntro)
	r.HandleFunc("/doc", handleDoc)
	r.HandleFunc("/", handleRoot)

	err := http.ListenAndServe(":8080", r) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you got the api")
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you got the about")
}

func handleCareers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you got the Careers")
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintf(w, "%v\n", t.Format("2006-January-02"))
	fmt.Fprintf(w, "Started Web Server....")
}

func handleIntro(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
    <head>
      <h2>My Header</h2>

    </head>
    <body>

        <button>Hello there</button>

    </body>

    `)
}

func handleDoc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
    <head>
      <h2>api documentation</h2>

    </head>
    <body>
        <li>/api</li>
        <li>/about</li>
        <li>/careers</li>
        <li>/intro</li>
        <li> / </li>"
    </body>
    `)
}
