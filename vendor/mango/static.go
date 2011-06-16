package mango

import (
	"path"
	"io/ioutil"
	"os"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	} else if !info.IsRegular() {
		return false
	}

	return true
}

func readFile(filename string) (string, os.Error) {
	body, err := ioutil.ReadFile(filename)
	return string(body), err
}

func Static(directory string) Middleware {
	return func(env Env, app App) (Status, Headers, Body) {
		// See if we can serve a file
		file := path.Join(directory, env.Request().URL.Path)
		if fileExists(file) && (env.Request().Method == "GET" || env.Request().Method == "HEAD") {
			if body, err := readFile(file); err == nil {
				return 200, Headers{}, Body(body)
			} else {
				panic(err)
			}
		}

		// No file found, pass on to app
		return app(env)
	}
}
