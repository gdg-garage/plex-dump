package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HostPort struct {
	host string
	port uint16
}

type Server struct {
	addr HostPort
}

func NewServer(addr HostPort) *Server {
	return &Server{
		addr: addr,
	}
}

func makeURL(hp HostPort, uri string) string {
	return fmt.Sprintf("http://%s:%d%s", hp.host, hp.port, uri)
}

func (s *Server) Get(uri string) ([]byte, error) {
	resp, err := http.Get(makeURL(s.addr, uri))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
