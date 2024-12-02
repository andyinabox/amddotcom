package main

import amdvanilla "github.com/amd-vanilla"

func main() {
	_, err := amdvanilla.LoadContent()

	if err != nil {
		panic(err)
	}
}
