package main

import _ "gopkg.in/hlandau/configurable.v0/exampleapp/examplelib"
import "gopkg.in/hlandau/configurable.v0/adaptflag"

//import "flag"
import flag "github.com/ogier/pflag"

func main() {
	adaptflag.AdaptWithFunc(func(info adaptflag.Info) {
		flag.Var(info.Value, info.Name, info.Usage)
	})

	flag.Parse()
}
