package main

import (
	"./mango"
	"flag"
)

var port *string = flag.String("port", "3000", "server port")

func Hello(env mango.Env) (mango.Status, mango.Headers, mango.Body) {
	return 200, mango.Headers{}, mango.Body("Hi from <a href='http://golang.org'>Go!</a><br />This is an experiment to get Go working on Heroku by <a href='http://twitter.com/qrush'>@qrush</a>.<br />Source is up on <a href='http://github.com/qrush/higo'>GitHub</a> for now until I get a blog post out.")
}

func main() {
	flag.Parse()
	stack := new(mango.Stack)
	stack.Address = ":" + *port
	stack.Run(Hello)
}
