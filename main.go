package main

import (
	"example/public"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/siuyin/dflt"
	"github.com/starfederation/datastar-go/datastar"
)

func main() {
	port := dflt.EnvString("PORT", "8080")
	log.Printf("PORT=%s\n", port)

	http.HandleFunc("/endpoint", endpointHandler)
	http.HandleFunc("/hal", halHandler)
	http.HandleFunc("/reset", resetHandler)
	http.Handle("/", http.FileServerFS(public.Content))

	log.Fatal(http.ListenAndServe(":"+port, nil))
	fmt.Println("vim-go")
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	sse.PatchElements(`<div id="hal">I’m sorry, Dave. I’m afraid I can’t do that.</div>`)
	time.Sleep(time.Second)
	sse.PatchElements(`<div id="hal">Waiting for an order...</div>`)
}

func halHandler(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	sse.PatchSignals([]byte(`{hal:"Broadcast: Affirmative, Dave. I read you.", fooBar: "Affirmative, Dave. I read you. "}`))
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	sse.PatchElements(`<div id="hal"></div>`)
	sse.PatchSignals([]byte(`{hal:"...", fooBar: ""}`))
}
