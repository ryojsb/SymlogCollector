package LogCollect

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/ryojsb/paramigo"
)

func GetSRDFinf(host string, port string, user string, password string, date string, sid string, path1 string, path2 string, wg *sync.WaitGroup) {
	log.Println("Start the process Collecting SRDF information.")
	cmd := "for i in $(BinPATHsymaccess -sid SID list -type storage | sed '1,6d'); do BinPATHsymrdf -sid SID -sg $i query; done > FilePATHDATE_SID_symrdf_query.txt"
	cmd = strings.ReplaceAll(cmd, "BinPATH", path1)
	cmd = strings.ReplaceAll(cmd, "FilePATH", path2)
	cmd = strings.Replace(cmd, "DATE", date, 1)
	cmd = strings.ReplaceAll(cmd, "SID", sid)
	paramigo.InnerCommand(host, port, user, password, cmd)
	log.Println("Finish up the process Collecting SRDF information.")
	wg.Done()
}

func GetTFinf(host string, port string, user string, password string, date string, sid string, path1 string, path2 string, wg *sync.WaitGroup) {
	log.Println("Start the process Collecting Timefinder information.")
	cmd := "for i in $(BinPATHsymaccess -sid SID list -type storage | sed '1,6d'); do BinPATHsymsnapvx -sid SID -sg $i list; done > FilePATHDATE_SID_symsnapvx_list.txt"
	cmd = strings.ReplaceAll(cmd, "BinPATH", path1)
	cmd = strings.ReplaceAll(cmd, "FilePATH", path2)
	cmd = strings.Replace(cmd, "DATE", date, 1)
	cmd = strings.ReplaceAll(cmd, "SID", sid)
	paramigo.InnerCommand(host, port, user, password, cmd)
	log.Println("Finish up the process Collecting Timefinder information.")
	wg.Done()
}

func DrLogCollect(host, port, user, password, sid, path1, path2 string) {
	t := time.Now()
	formattedTime := t.Format(TimeFormat)
	date := strings.ReplaceAll(formattedTime, "-", "")

	var wg sync.WaitGroup
	wg.Add(2)
	go GetSRDFinf(host, port, user, password, date, sid, path1, path2, &wg)
	go GetTFinf(host, port, user, password, date, sid, path1, path2, &wg)
	wg.Wait()
}
