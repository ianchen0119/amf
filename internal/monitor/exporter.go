package monitor

import (
	"net/http"

	"github.com/free5gc/amf/internal/logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type regisCollector struct {
	regisDetail *prometheus.Desc
}

func NewRegistrationCollector() *regisCollector {
	return &regisCollector{
		regisDetail: prometheus.NewDesc("registration_metric",
			"Amount of successful registration & received regis req",
			nil, nil),
	}
}

func (collector *regisCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.regisDetail
}

func (collector *regisCollector) Collect(ch chan<- prometheus.Metric) {
	regisSuccess := GetAmountOfSuccessRegistration()
	regisTry := GetAmountOfReceivedRegistration()
	prometheus.MustNewConstMetric(collector.regisDetail, prometheus.CounterValue, regisSuccess/regisTry)
}

func (collector *regisCollector) Start() {
	prometheus.MustRegister(collector)
	http.Handle("/metrics", promhttp.Handler())
	logger.HttpLog.Infoln("Monitor service is listening on 3000 port")
	logger.HttpLog.Errorln(http.ListenAndServe(":3000", nil))
}
