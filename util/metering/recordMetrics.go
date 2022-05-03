package metering

import (
	"fmt"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	requestsTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Subsystem: MeteringConfig.ServiceName,
			Name:      "http_requests_total",
			Help:      "Total number of http requests.",
		})

	requestsPerStatus = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: MeteringConfig.ServiceName,
			Name:      "http_requests_per_status",
			Help:      "Number of http requests per status(success, redirection, error).",
		},
		[]string{"status"},
	)

	requestsPerMethod = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: MeteringConfig.ServiceName,
			Name:      "http_requests_per_method",
			Help:      "Number of http requests per method(GET, PUT,...).",
		},
		[]string{"method"},
	)

	requestsPerOperationPerStatus = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: MeteringConfig.ServiceName,
			Name:      "http_requests_per_operation_per_status",
			Help:      "Number of http requests splitted per operation and per status (identity, offers, company...)(success, error).",
		},
		[]string{"operation", "status"},
	)

	requestsPerOperationPerMethod = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: MeteringConfig.ServiceName,
			Name:      "http_requests_per_operation_per_method",
			Help:      "Number of http requests splitted per operation and per method (identity, offers, company...)(GET, PUT,...).",
		},
		[]string{"operation", "method"},
	)

	requestsDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem: MeteringConfig.ServiceName,
			Name:      "http_requests_response_time_seconds_per_operation",
			Help:      "Duration of http requests (seconds) splitted per operation.",
			Buckets:   prometheus.LinearBuckets(0.5, 0.5, 8),
		},
		[]string{"operation"},
	)

	requestsMaxDuration = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: MeteringConfig.ServiceName,
			Name:      "http_requests_maximum_response_time_seconds_per_operation",
			Help:      "Maximum duration of http requests (seconds) splitted per operation.",
		},
		[]string{"operation"},
	)

	requestsAvgDuration = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: MeteringConfig.ServiceName,
			Name:      "http_requests_average_response_time_seconds_per_operation",
			Help:      "Average duration of http requests (seconds) splitted per operation.",
		},
		[]string{"operation"},
	)
)

//HTTP metrics
type httpRelatedInfo struct {
	operation  string
	method     string
	start      time.Time
	elapsed    time.Duration
	statusCode int
}

func (x httpRelatedInfo) computeBasicHTTPRelatedMetrics() {
	//log.Printf("record basic related metrics \n")

	var statusClass string
	switch {
	case x.statusCode >= 100 && x.statusCode < 200:
		statusClass = " 1XX informational response"
	case x.statusCode >= 200 && x.statusCode < 300:
		statusClass = "2XX success"
	case x.statusCode >= 300 && x.statusCode < 400:
		statusClass = "3XX redirection"
	case x.statusCode >= 400 && x.statusCode < 500:
		statusClass = "4XX client error"
	case x.statusCode >= 500 && x.statusCode < 600:
		statusClass = "5XX server error"
	default:
		statusClass = fmt.Sprintf("%d unknown status code", x.statusCode)
	}

	requestsTotal.Inc()
	requestsPerStatus.WithLabelValues(statusClass).Inc()
	requestsPerMethod.WithLabelValues(x.method).Inc()
	requestsPerOperationPerStatus.WithLabelValues(x.operation, statusClass).Inc()
	requestsPerOperationPerMethod.WithLabelValues(x.operation, x.method).Inc()
	requestsDuration.WithLabelValues(x.operation).Observe(x.elapsed.Seconds())

}

type computedHTTPStatistics struct {
	nrOfRequests int
	avgDuration  float64
	maxDuration  float64
}

var globalHTTPStatistics = struct {
	mutex      sync.Mutex
	statistics map[string]computedHTTPStatistics
}{
	statistics: map[string]computedHTTPStatistics{},
}

func (x httpRelatedInfo) computeEnhancedHTTPRelatedMetrics() {
	//log.Printf("record enhanced related metrics \n")

	newDuration := x.elapsed.Seconds()

	globalHTTPStatistics.mutex.Lock()

	if v, ok := globalHTTPStatistics.statistics[x.operation]; ok {
		var newAvg, newMax float64
		var newNr int

		//cumulative average formula ->  new_average = ((old_average * (n-1))  + new_value) / n
		newAvg = ((v.avgDuration * float64(v.nrOfRequests)) + newDuration) / (float64(v.nrOfRequests) + 1)

		if v.maxDuration < newDuration {
			newMax = newDuration
		} else {
			newMax = v.maxDuration
		}
		//number of requests (reset if close to upper limit)
		if v.nrOfRequests >= 9223372036850000000 {
			//reset
			newNr = 1
		} else {
			newNr = v.nrOfRequests + 1
		}

		globalHTTPStatistics.statistics[x.operation] = computedHTTPStatistics{
			nrOfRequests: newNr,
			avgDuration:  newAvg,
			maxDuration:  newMax,
		}

	} else {
		globalHTTPStatistics.statistics[x.operation] = computedHTTPStatistics{
			nrOfRequests: 1,
			avgDuration:  newDuration,
			maxDuration:  newDuration,
		}
	}
	globalHTTPStatistics.mutex.Unlock()

	requestsMaxDuration.WithLabelValues(x.operation).Set(globalHTTPStatistics.statistics[x.operation].maxDuration)
	requestsAvgDuration.WithLabelValues(x.operation).Set(globalHTTPStatistics.statistics[x.operation].avgDuration)
}
