package greetings

import( 
    "fmt"
    "errors"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // Return a greeting that embeds the name in a message.
    if name == "" {
        return "", errors.New("name cannot be empty") // this returns an empty string for the message var and an error in the err var which will be handled by the caller accordingly
    }
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil // this is the success return statement which returns nil error
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
    // formats is the slice of type string
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}