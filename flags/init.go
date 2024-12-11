package flags

import (
	"flag"
)

type Startup struct {
	ConfigFile string
	RunType    RunTypeArr
	logFile    string
}

var This Startup

func init() {
	flag.StringVar(&This.ConfigFile, "c", "", "required: path to .xocket.yaml")
	flag.StringVar(&This.logFile, "log", "", "path to save the output, default=stdout")
	flag.Var(&This.RunType, "run", "<debug | all>, default=all")
}
