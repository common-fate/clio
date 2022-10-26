package clio_test

import (
	"os"

	"github.com/common-fate/clio"
)

// Here's the basics on how to use clio for logging.
func Example_usage() {
	clio.SetLevelFromString("debug")

	clio.SetWriter(os.Stdout) // print to stdout just to show logs in the example.

	// you can print basic logs like this
	clio.Info("here's an info message")

	// add separate messages together with the 'ln' variant, e.g. clio.Errorln
	clio.Errorln("hello", "world")

	// format messages with the 'f' variant (same as how fmt.Printf works), e.g. clio.Debugf
	clio.Debugf("hello %s", "world")

	// add key-values pairs with the 'w' variant, e.g. clio.Infow
	clio.Infow("calling an API", "url", "http://example.com")

	// Output: [i] here's an info message
	// [✘] hello world
	// [DEBUG] hello world
	// [i] calling an API  	url:http://example.com
}
