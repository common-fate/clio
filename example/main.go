package main

import (
	"github.com/common-fate/clio"
	"github.com/common-fate/clio/clierr"
)

func main() {
	// set the level to 'debug' to demonstrate all messages being printed.
	clio.SetLevelFromString("debug")

	// uncomment to test file logging
	// clio.SetFileLogging(cliolog.FileLoggerConfig{
	// 	Filename: "log",
	// })

	clio.Infof("this is an example of calling clio.Infof with no argument %s")

	clio.Log("hello %s from clio.Log", "world")
	clio.Logf("hello %s from clio.Logf", "world")
	clio.Logln("hello %s from clio.Logln", "world")

	clio.Info("hello %s from clio.Info", "world")
	clio.Infof("hello %s from clio.Infof", "world")
	clio.Infoln("hello %s from clio.Infoln", "world")
	clio.Infow("hello %s from clio.Infow", "key", "value")

	clio.Success("hello %s from clio.Success", "world")
	clio.Successf("hello %s from clio.Successf", "world")
	clio.Successln("hello %s from clio.Successln", "world")
	clio.Successw("hello %s from clio.Successw", "key", "value")

	clio.Warn("hello %s from clio.Warn", "world")
	clio.Warnf("hello %s from clio.Warnf", "world")
	clio.Warnln("hello %s from clio.Warnln", "world")
	clio.Warnw("hello %s from clio.Warnw", "key", "value")

	clio.Error("hello %s from clio.Error", "world")
	clio.Errorf("hello %s from clio.Errorf", "world")
	clio.Errorln("hello %s from clio.Errorln", "world")
	clio.Errorw("hello %s from clio.Errorw", "key", "value")

	clio.Debug("hello %s from clio.Debug", "world")
	clio.Debugf("hello %s from clio.Debugf", "world")
	clio.Debugln("hello %s from clio.Debugln", "world")
	clio.Debugw("hello %s from clio.Debugw", "key", "value")

	clierr.New("example error", clierr.Info("some more details")).PrintCLIError()
}
