module github.com/ZolaraProject/user-api

go 1.23.2

require (
	github.com/ZolaraProject/library v0.1.1-rc05
	github.com/ZolaraProject/pki-vault-service v0.1.0-rc45
	github.com/gorilla/mux v1.8.1
	golang.org/x/crypto v0.32.0
	google.golang.org/grpc v1.70.0
)

require (
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250127172529-29210b9bc287 // indirect
	google.golang.org/protobuf v1.36.4 // indirect
)

replace github.com/ZolaraProject/pki-vault-service => ../pki-vault-service
