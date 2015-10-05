package main

import (
 "github.com/gdg-garage/plex-dump/Parser"
)


func GetLibraries(s *Server) ([]byte, error) {
	return s.Get("/library/sections/")
}

func main() {
	//parse hosts from stdin
	servers := []HostPort{
		HostPort{host: "127.0.0.1", port:32400},
	}

	for _, host := range servers {
		GetLibraries(NewServer(host))
	}
}