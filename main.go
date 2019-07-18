package main

import (
	"github.com/jenningsloy318/redfish_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"net/http"
)

var (
	configFile = kingpin.Flag(
		"config.file",
		"Path to configuration file.",
	).String()
	listenAddress = kingpin.Flag(
		"web.listen-address",
		"Address to listen on for web interface and telemetry.",
	).Default(":9610").String()

	reloadCh chan chan error
)

// define new http handleer
func metricsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		registry := prometheus.NewRegistry()

		redfishHosts := loadFilerFromFile(*configFile)
		for _, redfishHost := range redfishHosts {
			collector := collector.New(redfishHost.Host, redfishHost.Username, redfishHost.Password)
			registry.MustRegister(collector)
			gatherers := prometheus.Gatherers{
				prometheus.DefaultGatherer,
				registry,
			}
			// Delegate http serving to Prometheus client library, which will call collector.Collect.
			h := promhttp.HandlerFor(gatherers, promhttp.HandlerOpts{})
			h.ServeHTTP(w, r)
		}
	}
}

var Vsersion string
var BuildRevision string
var BuildBranch string
var BuildTime string
var BuildHost string
func init() {
	log.Infof("redfish_exporter version %s, build reversion %s, build branch %s, build at %s on host %s",Vsersion,BuildRevision,BuildBranch,BuildTime,BuildHost)
}

func main() {
	log.AddFlags(kingpin.CommandLine)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	log.Infoln("Starting redfish_exporter")

	http.Handle("/metrics", metricsHandler()) // Regular metrics endpoint for local Redfish metrics.

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
            <head>
            <title>Redfish Exporter</title>
            </head>
            <body>
			<p><a href="/metrics">Local metrics</a></p>
            </body>
            </html>`))
	})

	log.Infof("Listening on %s", *listenAddress)
	err := http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}