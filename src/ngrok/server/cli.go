package server

import (
	"flag"
)

type Options struct {
	httpAddr     string
	httpsAddr    string
	srcHttpAddr  string
	srcHttpsAddr string
	tunnelAddr   string
	domain       string
	tlsCrt       string
	tlsKey       string
	logto        string
	loglevel     string
}

func parseArgs() *Options {
	httpAddr := flag.String("httpAddr", ":80", "Public address for HTTP connections, empty string to disable")
	httpsAddr := flag.String("httpsAddr", ":443", "Public address listening for HTTPS connections, emptry string to disable")
	srcHttpAddr := flag.String("srcHttpAddr", "", "Source public address for HTTP connections proxy, empty string to disable")
	srcHttpsAddr := flag.String("srcHttpsAddr", "", "Source public address listening for HTTPS connections proxy, emptry string to disable")
	tunnelAddr := flag.String("tunnelAddr", ":4443", "Public address listening for ngrok client")
	domain := flag.String("domain", "ngrok.com", "Domain where the tunnels are hosted")
	tlsCrt := flag.String("tlsCrt", "", "Path to a TLS certificate file")
	tlsKey := flag.String("tlsKey", "", "Path to a TLS key file")
	logto := flag.String("log", "stdout", "Write log messages to this file. 'stdout' and 'none' have special meanings")
	loglevel := flag.String("log-level", "DEBUG", "The level of messages to log. One of: DEBUG, INFO, WARNING, ERROR")
	flag.Parse()

	return &Options{
		httpAddr:     *httpAddr,
		httpsAddr:    *httpsAddr,
		srcHttpAddr:  *srcHttpAddr,
		srcHttpsAddr: *srcHttpsAddr,
		tunnelAddr:   *tunnelAddr,
		domain:       *domain,
		tlsCrt:       *tlsCrt,
		tlsKey:       *tlsKey,
		logto:        *logto,
		loglevel:     *loglevel,
	}
}
