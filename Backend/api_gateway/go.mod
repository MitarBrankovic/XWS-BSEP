module dislinkt/api_gateway

go 1.18

replace dislinkt/common => ../common

require (
	dislinkt/common v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	github.com/joho/godotenv v1.4.0
	google.golang.org/grpc v1.46.2
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220505152158-f39f71e6c8f3 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
