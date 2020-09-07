module hello

go 1.15

replace example.com/greetings => ../greetings

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	github.com/go-kit/kit v0.10.0 // indirect
)
