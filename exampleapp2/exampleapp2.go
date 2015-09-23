package main

import "gopkg.in/hlandau/configurable.v0"
import "gopkg.in/hlandau/configurable.v0/cstruct"
import "gopkg.in/hlandau/configurable.v0/adaptflag"
import "gopkg.in/hlandau/configurable.v0/adaptconf"
import flag "github.com/ogier/pflag"
import "fmt"

type Config struct {
	Bind           string `default:":53" usage:"Address to bind to (e.g. 0.0.0.0:53)"`
	PublicKey      string `default:"" usage:"Path to the DNSKEY KSK public key file"`
	PrivateKey     string `default:"" usage:"Path to the KSK's corresponding private key file"`
	ZonePublicKey  string `default:"" usage:"Path to the DNSKEY ZSK public key file; if one is not specified, a temporary one is generated on startup and used only for the duration of that process"`
	ZonePrivateKey string `default:"" usage:"Path to the ZSK's corresponding private key file"`

	NamecoinRPCUsername string `default:"" usage:"Namecoin RPC username"`
	NamecoinRPCPassword string `default:"" usage:"Namecoin RPC password"`
	NamecoinRPCAddress  string `default:"localhost:8336" usage:"Namecoin RPC server address"`
	CacheMaxEntries     int    `default:"100" usage:"Maximum name cache entries"`
	SelfName            string `default:"" usage:"The FQDN of this nameserver. If empty, a psuedo-hostname is generated."`
	SelfIP              string `default:"127.127.127.127" usage:"The canonical IP address for this service"`

	HTTPListenAddr string `default:"" usage:"Address for webserver to listen at (default: disabled)"`

	CanonicalSuffix      string `default:"bit" usage:"Suffix to advertise via HTTP"`
	CanonicalNameservers string `default:"" usage:"Comma-separated list of nameservers to use for NS records. If blank, SelfName (or autogenerated psuedo-hostname) is used."`
	canonicalNameservers []string
	Hostmaster           string `default:"" usage:"Hostmaster e. mail address"`
	VanityIPs            string `default:"" usage:"Comma separated list of IP addresses to place in A/AAAA records at the zone apex (default: don't add any records)"`
}

func main() {
	tgt := &Config{}
	c := cstruct.MustNew(tgt, "example")
	configurable.Register(c)

	adaptflag.AdaptWithFunc(func(info adaptflag.Info) {
		flag.Var(info.Value, info.Name, info.Usage)
	})

	flag.Parse()

	err := adaptconf.Load("exampleapp2")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Printf("%#v\n", tgt)
}
