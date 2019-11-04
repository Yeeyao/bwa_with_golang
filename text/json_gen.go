package main

import (
	"encoding/json"
	"os"
)

type Server struct {
	ID          int    `json:"-"`
	ServerName  string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"`
	ServerIP    string `json:"serverIP,omitempty"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}

func main() {
	s := Server{
		ID:          3,
		ServerName:  `Go "1.0"`,
		ServerName2: `Go "1.0"`,
		ServerIP:    ``,
	}
	b, _ := json.Marshal(s)
	os.Stdout.Write(b)
}
