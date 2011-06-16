package main

import (
	"github.com/paulbellamy/mango"
	"flag"
)

var port *string = flag.String("port", "3000", "server port")

func Hello(env mango.Env) (mango.Status, mango.Headers, mango.Body) {
	return 200, mango.Headers{}, mango.Body("Hi Go!")
}

func main() {
	flag.Parse()
	stack := new(mango.Stack)
	stack.Address = ":" + *port
	stack.Run(Hello)
}
