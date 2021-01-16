module github.com/koltyakov/spvault

go 1.16

replace github.com/koltyakov/spvault => ./

require (
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.6 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.5
	github.com/koltyakov/gosip v0.0.0-20210116122858-f3adcc821533
	github.com/koltyakov/gosip-sandbox v0.0.0-20210109235807-2bbd70e6a6e0
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20210114201628-6edceaf6022f // indirect
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)
