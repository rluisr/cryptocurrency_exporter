package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"io"
	"net/http"
	"os"
	"strings"
)

func Probe(ctx context.Context, coin string, registry *prometheus.Registry) bool {
	registry.MustRegister(rateGauge, volumeGauge, capGauge, deltaHourGauge, deltaDayGauge, deltaWeekGauge, deltaMonthGauge, deltaQuarterGauge, deltaYearGauge)

	coinResponse, err := request(ctx, coin)
	if err != nil {
		writeError(err)
		return false
	}

	rateGauge.WithLabelValues(coin).Set(coinResponse.Rate)
	volumeGauge.WithLabelValues(coin).Set(coinResponse.Volume)
	capGauge.WithLabelValues(coin).Set(coinResponse.Cap)

	deltaHour, isOK := coinResponse.Delta.Hour.(float64)
	if isOK {
		deltaHourGauge.WithLabelValues(coin).Set(deltaHour)
	}
	deltaDay, isOK := coinResponse.Delta.Day.(float64)
	if isOK {
		deltaDayGauge.WithLabelValues(coin).Set(deltaDay)
	}
	deltaWeek, isOK := coinResponse.Delta.Week.(float64)
	if isOK {
		deltaWeekGauge.WithLabelValues(coin).Set(deltaWeek)
	}
	deltaMonth, isOK := coinResponse.Delta.Month.(float64)
	if isOK {
		deltaMonthGauge.WithLabelValues(coin).Set(deltaMonth)
	}
	deltaYear, isOK := coinResponse.Delta.Year.(float64)
	if isOK {
		deltaYearGauge.WithLabelValues(coin).Set(deltaYear)
	}

	return true
}

func request(ctx context.Context, coin string) (*Response, error) {
	data := strings.Split(coin, "/")
	code := data[0]
	currency := data[1]

	body := Request{
		Currency: currency,
		Code:     code,
		Meta:     false,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	payload := strings.NewReader(string(b))

	reqURL := fmt.Sprint("https://api.livecoinwatch.com/coins/single")
	req, _ := http.NewRequestWithContext(ctx, "POST", reqURL, payload)
	req.Header.Add("x-api-key", config.APIKey)
	req.Header.Add("content-type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch coin status body: %s", string(b))
	}

	var coinResponse Response
	err = json.Unmarshal(b, &coinResponse)

	return &coinResponse, err
}

func writeError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "[%s_exporter ERROR] %s\n", NameSpace, err.Error())
}
