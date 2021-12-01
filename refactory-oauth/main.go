package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func main() {
	key := "Secretsssss"
	maxAge := 86400 * 30
	isProd := false

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New("302445806692-0d09fl4kut3a3mqf5um6otid6v1f0ci9.apps.googleusercontent.com", "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", "http://localhost:3000/auth/google/callback", "email", "profile"),
	)

	p := pat.New()
	p.Get("/auth/{provider}/callback", func(writer http.ResponseWriter, request *http.Request) {

		user, err := gothic.CompleteUserAuth(writer, request)
		if err != nil {
			fmt.Fprintln(writer, err)
			return
		}

		t, _ := template.ParseFiles("templates/success.html")
		t.Execute(writer, user)
	})

	p.Get("/auth/{provider}", func(writer http.ResponseWriter, request *http.Request) {
		gothic.BeginAuthHandler(writer, request)
	})

	p.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(writer, false)
	})

	log.Println("Listening on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", p))
}
