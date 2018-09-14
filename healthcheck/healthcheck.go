package healthcheck

import (
	"log"
	"net/http"
	"net/http/httptrace"
)

type Getter interface {
}

type Doer interface {
}

type GetterDoer interface {
}

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
	log.Println("get conn info")
}

func (t *transport) GotConnInfo(info httptrace.GotConnInfo) {
	log.Println("got conn info")
}

func GetRequest(url string, t *transport) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	trace := &httptrace.ClientTrace{
		GetConn: t.GetConInfo,
		GotConn: t.GotConnInfo,
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
