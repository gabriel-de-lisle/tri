module github.com/gabriel-de-lisle/tri/cmd

go 1.19

require (
	github.com/gabriel-de-lisle/tri/todo v0.0.0-00010101000000-000000000000
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.5.0
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace github.com/gabriel-de-lisle/tri/todo => ../todo
