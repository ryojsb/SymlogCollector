package main

import (
	"flag"
	"log"
	"time"

	"./LogCollect"
)

var (
	way      string
	host     string
	port     string
	user     string
	password string
	sid      string
	path1    string
	path2    string
)

func main() {
	flag.StringVar(&way, "w", "", "Purpose")
	flag.StringVar(&host, "host", "", "IP Address")
	flag.StringVar(&port, "port", "", "Port")
	flag.StringVar(&user, "u", "", "user")
	flag.StringVar(&password, "p", "", "Password")
	flag.StringVar(&sid, "sid", "", "Symmetrix S/N")
	flag.StringVar(&path1, "path1", "", "PATH for bin")
	flag.StringVar(&path2, "path2", "", "PATH for log file")
	flag.Parse()

	log.Println("Start the process Collecting information processes.")
	time.Sleep(2 * time.Second)
	log.Println("Start the Parallel Processing.")
	time.Sleep(2 * time.Second)

	if way == "cfg" {
		LogCollect.CfgLogCollect(host, port, user, password, sid, path1, path2)
	} else if way == "dr" {
		LogCollect.DrLogCollect(host, port, user, password, sid, path1, path2)
	}
}
