package main

import (
	"fmt"
	"github.com/shomali11/slacker"
	"os"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4052568962279-4060599505462-EbrMf52PFXNy2JoJZUw5ThL3")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A042BP9PQTB-4052608843463-c0d0aee2dbb22ab80bc579bddea87661f32e000c3dd57f805378f2738e62d1d0")
}
