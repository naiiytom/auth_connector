# Keycloak Authentication Middleware

A Golang issuer backend for for communicating with Keycloak Authentication service

## Deployment

Change environment variables in `env-file`

Deploy using docker using command to build docker image

    docker build -t auth-middleware .

Then run command

    docker run -itd --name auth-middleware -p 5000:5000 --restart=unless-stopped auth-middleware

## Environment variables

|          Var           |        Default         |             Description            |
|------------------------|------------------------|------------------------------------|
| KEYCLOAK_HOST          | http://localhost:8080  | Keycloak address                   |
| KEYCLOAK_REALM         | default                | Realm name                         |
| KEYCLOAK_CLIENT_ID     | default-client         | Client name                        |
| KEYCLOAK_CLIENT_SECRET | `blank`                | Client secret (optional) if not set|
