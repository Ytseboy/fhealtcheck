package healthcheck

import (
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

var (
	GetConnTime time.Time
)

type transport struct {
	current *http.Request
}

type TracingClient struct {
	http.Client
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	return http.DefaultTransport.RoundTrip(req)
}

func (t *transport) GetConInfo(hostPort string) {
	log.Println("get con info")
}

func GetRequest(url string, t *transport) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	trace := &httptrace.ClientTrace{
		GetConn: t.GetConInfo,
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	return req, err
}

func Get(client TracingClient, url string) *http.Response {
	t := &transport{}
	req, err := GetRequest(url, t)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}

// func (t *TraceClient) GotConInfo(nfo httptrace.GotConnInfo) {
// 	log.Println("got con info")
// 	t.GotConnTime = time.Now()
// }

// func (t *TraceClient) RoundTrip(req *http.Request) (*http.Response, error) {
// 	t.RTCounter++
// 	return http.DefaultTransport.RoundTrip(req)
// }
