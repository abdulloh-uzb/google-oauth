package controller

import (
	"context"
	"fmt"
	"google-oauth/config"
	"io/ioutil"
	"net/http"
)

func Home(res http.ResponseWriter, req *http.Request) {
	html := `<a href="/google/login">Google orqali kirish</a>`
	fmt.Println("home")
	fmt.Fprintln(res, html)
}

func GoogleLogin(res http.ResponseWriter, req *http.Request) {
	conf := config.SetupConfig()
	url := conf.AuthCodeURL("state")
	http.Redirect(res, req, url, http.StatusSeeOther)
}

func GoogleCallback(res http.ResponseWriter, req *http.Request) {
	state := req.URL.Query()["state"][0]

	if state != "state" {
		fmt.Fprintln(res, "states dont match")
		return
	}

	code := req.URL.Query()["code"][0]
	conf := config.SetupConfig()

	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		fmt.Fprintln(res, "code token exchange failed")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Fprintln(res, "error while get user from google")
	}

	user, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(res, "code token exchange failed")

	}
	fmt.Println("token: ", token.AccessToken)
	// fmt.Println("response body: ", resp.Body)
	// fmt.Println("response :", resp)
	fmt.Fprintln(res, string(user))
}
