package config

import (
	"fmt"
	"os"
)

var (
	KeycloakHost         = getEnv("KEYCLOAK_HOST", "http://localhost:8000")
	KeycloakRealm        = getEnv("KEYCLOAK_REALM", "master")
	KeycloakClientID     = getEnv("KEYCLOAK_CLIENT_ID", "default")
	KeycloakClientSecret = getEnv("KEYCLOAK_CLIENT_SECRET", "")
)

func getEnv(key string, defaultValue string) string {
	val, exist := os.LookupEnv(key)
	if !exist {
		val = defaultValue
	}
	fmt.Println(val)
	return val
}
