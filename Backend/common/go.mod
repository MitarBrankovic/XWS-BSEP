module dislinkt/common

go 1.18

replace dislinkt/common => ../common

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	google.golang.org/genproto v0.0.0-20220505152158-f39f71e6c8f3
	google.golang.org/grpc v1.46.2
	google.golang.org/protobuf v1.28.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
)
