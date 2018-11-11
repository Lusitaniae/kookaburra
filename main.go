package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

var (
	SimpsonTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "homer_simpson_requests_total",
		Help: "The total number of requests for the Homer Simpson image.",
	})

	CovilhaTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "covilha_requests_total",
		Help: "The total number of requests for the time in Covilha.",
	})
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	message := `
<html>
Welcome to the Kookaburra application: <br>
- <a href="/homersimpson">Homer Simpson's Picture</a>.<br>
- <a href="/covilha">Time in Covilha, Portugal</a>.<br>
- <a href="/metrics">Application metrics</a>.<br>
</html>
`
	fmt.Fprint(w, message)
}

func Simpson(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./simpson.png")
	SimpsonTotal.Inc()
}

func Covilha(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tz, _ := time.LoadLocation("Europe/Lisbon")
	time := time.Now().In(tz)
	fmt.Fprint(w, "Time in Covilha: ")
	fmt.Fprintf(w, time.Format("2006-01-02 15:04:05"))
	CovilhaTotal.Inc()
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/homersimpson", Simpson)
	router.GET("/covilha", Covilha)
	router.Handler("GET", "/metrics", promhttp.Handler())

	log.Println("Server is listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))

}
