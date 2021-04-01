package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/naiiytom/auth_connector/config"
	"github.com/naiiytom/auth_connector/internal/user"

	"github.com/Nerzal/gocloak/v8"
)

var (
	realm         = config.KeycloakRealm
	client_name   = config.KeycloakClientID
	client_secret = config.KeycloakClientSecret
	client        = gocloak.NewClient(config.KeycloakHost)
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
