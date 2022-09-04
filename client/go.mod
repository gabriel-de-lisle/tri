module github.com/gabriel-de-lisle/tri/client

go 1.19

replace github.com/gabriel-de-lisle/tri/client/cmd => ./cmd

replace github.com/gabriel-de-lisle/tri/protocol => ../protocol

require (
	github.com/gabriel-de-lisle/tri/protocol v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.5.0
	google.golang.org/grpc v1.49.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.0.0-20220826154423-83b083e8dc8b // indirect
	golang.org/x/sys v0.0.0-20220829200755-d48e67d00261 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220902135211-223410557253 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
