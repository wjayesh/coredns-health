package main

import (
	"flag"
)

var (
	pathToCfg   = flag.String("path", "", "the path to the kubeconfig file")
	podsAllowed = flag.String("allowPods", "false", "allow creation of lightweight pods in cluster")
	udpPort     = flag.String("port", "53", "the udp port for the dns server")
)

func main() {
	flag.Parse()
	var e engine
	e.Init(*pathToCfg)

	prefs := make(map[string]string)
	prefs["podsAllowed"] = *podsAllowed
	prefs["port"] = *udpPort
	e.SetOptions(prefs)

	e.Start()
}