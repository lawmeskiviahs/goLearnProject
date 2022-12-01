module examples.com/hello

go 1.18

replace example.com/greetings => ../greetings

replace example.com/maths => ../maths

require github.com/gorilla/mux v1.8.0

replace example/localTest => ../localTest
