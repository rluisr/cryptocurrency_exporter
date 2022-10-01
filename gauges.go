package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	rateGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: NameSpace,
		Name:      "rate",
		Help:      "rate",
	},
		[]string{"coin"})
	volumeGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: NameSpace,
		Name:      "volume",
		Help:      "volume",
	}, []string{"coin"})
	capGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: NameSpace,
		Name:      "cap",
		Help:      "cap",
	}, []string{"coin"})
	deltaHourGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: NameSpace,
		Name:      "delta_hour",
		Help:      "delta hour",
	}, []string{"coin"})
	deltaDayGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: NameSpace,
		Name:      "delta_day",
		Help:      "delta day",
	}, []string{"coin"})
	deltaWeekGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: NameSpace,
		Name:      "delta_week",
		Help:      "delta week",
	}, []string{"coin"})
	deltaMonthGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: NameSpace,
		Name:      "delta_month",
		Help:      "delta month",
	}, []string{"coin"})
	deltaQuarterGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: NameSpace,
		Name:      "delta_quarter",
		Help:      "delta quarter",
	}, []string{"coin"})
	deltaYearGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: NameSpace,
		Name:      "delta_year",
		Help:      "delta year",
	}, []string{"coin"})
)
