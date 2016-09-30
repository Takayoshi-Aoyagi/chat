package main

import (
	"flag"
	"log"
	"net/http"
	"text/template"
	"os"
	"path/filepath"
	"sync"
	"github.com/Takayoshi-Aoyagi/tracer"
)

type templateHandler struct {
	once sync.Once
	filename string
	templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
			t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "IP Address")
	flag.Parse()
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// Start chat room
	go r.run()
	// Web server
	log.Println("Starting web server, port: ", *addr)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
