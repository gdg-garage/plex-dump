package main

import (
	"fmt"
	"github.com/gdg-garage/plex-dump/Parser"
)

func getLibraryDirectories(s *Server) []Parser.LibraryDirectory {
	data, err := s.Get("/library/sections/")
	if nil != err {
		fmt.Println("Error: ", err)
		return nil
	}
	return Parser.ParseLibraries(data)
}

func processHost(host HostPort) {
	server := NewServer(host)

	dirs := getLibraryDirectories(server)

	fmt.Println(dirs)
}

func main() {
	//parse hosts from stdin
	servers := []HostPort{
		HostPort{host: "127.0.0.1", port: 32400},
	}

	for _, host := range servers {
		processHost(host)
	}
}
