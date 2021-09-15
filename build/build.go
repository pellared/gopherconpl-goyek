package main

import (
	"github.com/goyek/goyek"
)

func main() {
	flow := &goyek.Flow{}

	msg := flow.RegisterStringParam(goyek.StringParam{
		Name:    "msg",
		Usage:   "message the world",
		Default: "Hello world!",
	})

	hello := flow.Register(goyek.Task{
		Name:  "hello",
		Usage: "demonstration",
		Action: func(tf *goyek.TF) {
			tf.Log(msg.Get(tf))
			// tf.Error("Something is wrong")
		},
		Params: goyek.Params{msg},
	})

	test := flow.Register(goyek.Task{
		Name:  "test",
		Usage: "go test",
		Action: func(tf *goyek.TF) {
			if err := tf.Cmd("go", "test").Run(); err != nil {
				tf.Error(err)
			}
		},
		Deps: goyek.Deps{hello},
	})

	flow.DefaultTask = test

	flow.Main()
}
