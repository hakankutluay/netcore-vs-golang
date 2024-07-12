package main

import (
	"fmt"
	"go-test/models"
	"html"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
)

func main() {
	url := "http://" + os.Getenv("HOST") + ":5002/data"

	client := &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialTimeout(addr, time.Second*10)
		},
		MaxConnsPerHost:     100_000,
		MaxIdleConnDuration: time.Second * 10,
	}

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		var dst []byte
		_, bdy, err := client.Get(dst, url)
		if err != nil {
			serverError(w, err.Error())
			return
		}

		var obj models.Response
		if err := easyjson.Unmarshal(bdy, &obj); err != nil {
			serverError(w, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")

		// marshall with easyjson the object to the response
		if _, _, err := easyjson.MarshalToHTTPResponseWriter(&obj, w); err != nil {
			serverError(w, err.Error())
			return
		}

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	addr := ":5001"
	fmt.Println("listening on " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func serverError(w http.ResponseWriter, msg string) {
	http.Error(w, msg, http.StatusInternalServerError)
}
