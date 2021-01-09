module github.com/koltyakov/spvault

go 1.16

replace github.com/koltyakov/spvault => ./

require (
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.4
	github.com/koltyakov/gosip v0.0.0-20210106125641-a1603f5f707a
	github.com/koltyakov/gosip-sandbox v0.0.0-20210109230414-0568abc3ad49
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20210108203827-ffc7fda8c3d7 // indirect
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)
