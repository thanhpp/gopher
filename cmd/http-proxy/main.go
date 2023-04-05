package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	p := NewProxies("https://bsc-dataseed.binance.org",
		"https://bsc-dataseed1.binance.org",
		"https://bsc-dataseed2.binance.org",
		"https://bsc-dataseed3.binance.org",
		"https://bsc-dataseed4.binance.org",
		"https://bsc-dataseed1.defibit.io",
		"https://bsc-dataseed1.ninicoin.io",
		"https://bsc.nodereal.io")
	s := &Server{p}
	// s := &Server{NewProxies("https://ipinfo.io/")}

	err := http.ListenAndServe("127.0.0.1:12342", s)
	if err != nil {
		panic(err)
	}
}

type Server struct {
	p *proxies
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	px, rpc := s.p.Next()
	log.Printf("from: %s\nto: %s\nmethod: %s\nuse: %s", req.RemoteAddr, req.URL.String(), req.Method, rpc)

	data, _ := io.ReadAll(req.Body)
	req.Body.Close()
	req.Body = io.NopCloser(bytes.NewReader(data))

	log.Printf("req: %+v\nbody: %s", req, string(data))

	px.ServeHTTP(w, req)
}

type proxies struct {
	proxies []*httputil.ReverseProxy
	rpcs    []string
	idx     int
}

func NewProxies(rpcs ...string) *proxies {
	p := new(proxies)
	p.rpcs = rpcs

	for i := range rpcs {
		u, err := url.Parse(rpcs[i])
		if err != nil {
			panic(err)
		}
		log.Printf("%+v", u.Host)
		px := httputil.NewSingleHostReverseProxy(u)
		directorFn := px.Director
		px.Director = func(r *http.Request) {
			directorFn(r)
			r.Host = u.Host
		}
		p.proxies = append(p.proxies, px)
	}

	return p
}

func (p *proxies) Next() (*httputil.ReverseProxy, string) {
	p.idx = (p.idx + 1) % len(p.proxies)

	return p.proxies[p.idx], p.rpcs[p.idx]
}
