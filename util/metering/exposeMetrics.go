package metering

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func PrometheusHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		promhttp.Handler().ServeHTTP(w, r)
	})
}

//middleware related
type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newCustomResponseWriter(w http.ResponseWriter) *customResponseWriter {
	return &customResponseWriter{w, http.StatusOK}
}

func (crw *customResponseWriter) WriteHeader(code int) {
	crw.statusCode = code
	crw.ResponseWriter.WriteHeader(code)
}

var MeteringConfig = struct {
	ServiceName string
	Enabled     bool
}{}

func MeteringMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Printf("start - metering middleware - enabled:%v servicename:%s \n", MeteringConfig.Enabled, MeteringConfig.ServiceName)

		if MeteringConfig.Enabled {

			// identify operation (endpoint last segment)
			vars := mux.Vars(r)
			//log.Printf("path parameters -> %+v\n", vars)
			segs := strings.Split(r.URL.Path, "/")
			//log.Printf("path segments -> %+v\n", segs)
			op := ""
		xloop:
			for i := len(segs); i > 0; i-- {
				isVar := false
				// check if the segment is path variable
				for _, v := range vars {
					if segs[i-1] == v {
						isVar = true
					}
				}

				if !isVar {
					op = segs[i-1]
					break xloop
				}
			}

			m := httpRelatedInfo{
				operation: op,
				method:    r.Method,
				start:     time.Now(),
			}

			crw := newCustomResponseWriter(w)
			next.ServeHTTP(crw, r)

			m.elapsed = time.Since(m.start)
			m.statusCode = crw.statusCode

			m.computeBasicHTTPRelatedMetrics()
			m.computeEnhancedHTTPRelatedMetrics()
			//log.Printf("metering middleware - operation:%s  httpinfo:%+v \n", op, m)
		} else {
			next.ServeHTTP(w, r)
		}
		//log.Println("end - metering middleware")
	})
}
