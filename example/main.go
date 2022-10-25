package main

import (
	"github.com/common-fate/clio"
	"github.com/common-fate/clio/clierr"
)

func main() {
	// set the level to 'debug' to demonstrate all messages being printed.
	clio.SetLevelFromString("debug")

	clio.Infof("this is an example of calling clio.Infof with no argument %s")

	clio.Info("hello %s from clio.Info", "world")
	clio.Infof("hello %s from clio.Infof", "world")
	clio.Infow("hello %s from clio.Infow", "key", "value")

	clio.Success("hello %s from clio.Success", "world")
	clio.Successf("hello %s from clio.Successf", "world")
	clio.Successw("hello %s from clio.Successw", "key", "value")

	clio.Warn("hello %s from clio.Warn", "world")
	clio.Warnf("hello %s from clio.Warnf", "world")
	clio.Warnw("hello %s from clio.Warnw", "key", "value")

	clio.Error("hello %s from clio.Error", "world")
	clio.Errorf("hello %s from clio.Errorf", "world")
	clio.Errorw("hello %s from clio.Errorw", "key", "value")

	clio.Debug("hello %s from clio.Debug", "world")
	clio.Debugf("hello %s from clio.Debugf", "world")
	clio.Debugw("hello %s from clio.Debugw", "key", "value")

	clierr.New("example error", clierr.Info("some more details")).PrintCLIError()
}
