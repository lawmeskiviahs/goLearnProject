module examples.com/hello

go 1.18

replace example.com/greetings => ../greetings

replace example.com/maths => ../maths

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	example.com/maths v0.0.0-00010101000000-000000000000
	example/localTest v0.0.0-00010101000000-000000000000
)

replace example/localTest => ../localTest
