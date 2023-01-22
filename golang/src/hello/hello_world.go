package main

import "fmt"

const englishHelloGreeting = "Hello, "
const englishHelloDefault = "World!"

func Hello(name string) string {
	if name != "" {
		return fmt.Sprintf("%s%s", englishHelloGreeting, name)
	} else {
		return fmt.Sprintf("%s%s", englishHelloGreeting, englishHelloDefault)
	}
}

func main() {
	name := "Paul"
	fmt.Println(Hello(name))
}
