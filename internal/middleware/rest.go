package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/naiiytom/auth_connector/internal/user"

	"github.com/Nerzal/gocloak/v8"
)

var (
	realm         = "aigen"
	client_name   = "aigen-client"
	client_secret = "8dfddd7c-00dc-44d9-b437-723c9c84118f"
	client        = gocloak.NewClient("http://192.168.1.99:8080")
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u user.LoginUser
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Println("decoded to user", u)

	ctx := context.Background()
	fmt.Println("username", u.UserName)
	token, err := client.Login(ctx, client_name, client_secret, realm, u.UserName, u.Password)

	if err != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
	} else {
		fmt.Println("token: ", token)
		json.NewEncoder(w).Encode(token)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u user.RefreshToken
	json.NewDecoder(r.Body).Decode(&u)
	ctx := context.Background()
	err := client.Logout(ctx, client_name, client_secret, realm, u.RefreshToken)

	if err != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
	} else {
		w.Write([]byte("OK"))
	}
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var t user.RefreshToken
	json.NewDecoder(r.Body).Decode(&t)
	ctx := context.Background()
	token, err := client.RefreshToken(ctx, t.RefreshToken, client_name, client_secret, realm)

	if err != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
	} else {
		json.NewEncoder(w).Encode(token)
	}
}
