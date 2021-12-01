package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"oauth/database"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:3000/auth/google/callback",
		ClientID:     "302445806692-0d09fl4kut3a3mqf5um6otid6v1f0ci9.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	// TODO: randomize it
	randomState = "randomak"
)

func main() {

	http.HandleFunc("/", handlerHome)
	http.HandleFunc("/login", handlerLogin)
	http.HandleFunc("/auth/google/callback", handlerCallback)
	http.HandleFunc("/show", handlerShow)
	http.ListenAndServe(":3000", nil)
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	var html = `<html><body><a href="/login">Google Log In</a></body></html>`
	fmt.Fprint(w, html)
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handlerCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != randomState {
		fmt.Println("state is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		fmt.Printf("could not get token: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("could not create get request: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("could not parse response: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// fmt.Fprintf(w, "Response: %s", content)

	user := &database.User{}

	err = json.Unmarshal(content, user)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/refactory")
	if err != nil {
		panic(err)
	}
	userRepository := database.NewRepository(db)
	ctx := context.Background()

	result, err := userRepository.Save(ctx, *user)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	var html = `<html><body>Success add to database <a href="/show">Show Data</a></body></html>`
	fmt.Fprint(w, html)

}

func handlerShow(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/refactory")
	if err != nil {
		panic(err)
	}
	userRepository := database.NewRepository(db)
	ctx := context.Background()

	users, err := userRepository.GetAllData(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, users)
}
