package main

import (
	"strconv"
	"io/ioutil"
	"net"
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

func (s *Server) Get(uri string) ([]byte, error) {
	resp, err := http.Get(net.JoinHostPort(s.addr.host, strconv.FormatUint(uint64(s.addr.port), 10)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}