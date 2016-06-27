package adaptflag_test

import "gopkg.in/hlandau/easyconfig.v1/adaptflag"

import (
	"flag"
	"github.com/ogier/pflag"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Example of how to use with "flag", "pflag", and "kingpin" (you only
// really need the line for whichever you are using).  Note that
// AdaptWithFunc will not do anything with Configurables which it has
// already adapted once, thus it is safe to call the function multiple
// times.
func Example() {
	adaptflag.AdaptWithFunc(func(info Info) {
		dpn := adaptflag.DottedPath(info.Path)
		if len(dpn) > 0 {
			dpn += "."
		}
		dpn += info.Name
		flag.Var(info.Value, dpn, info.Usage)
		pflag.Var(info.Value, dpn, info.Usage)
		kingpin.Flag(dpn, info.Usage).SetValue(info.Value)
	})
}
