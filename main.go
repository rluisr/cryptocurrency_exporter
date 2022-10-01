package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net"
	"net/http"
	"time"
)

var (
	config     *Config
	httpClient *http.Client
	version    string
)

const (
	NameSpace = "cryptocurrency"
)

func main() {
	var err error

	httpClient = &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          128,
			MaxIdleConnsPerHost:   100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		Timeout: 5 * time.Second,
	}

	config, err = NewConfig()
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", config.ListenPort)
	log.Printf("Start exporter on %s/metrics", addr)

	http.HandleFunc("/", top)
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		probeHandler(w, r)
	})

	log.Fatal(http.ListenAndServe(addr, nil))
}

func probeHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	// target: http://hoge?coin=BTC/USDT
	// Need parse target for getting a coin
	coin := params.Get("coin")
	if coin == "" {
		http.Error(w, "parameter \"coin\" is missing", http.StatusBadRequest)
		return
	}

	registry := prometheus.NewRegistry()
	_ = Probe(r.Context(), coin, registry)

	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}

func top(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "version: %s", version)
}
