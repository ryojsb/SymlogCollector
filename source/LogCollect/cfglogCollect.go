package LogCollect

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/ryojsb/paramigo"
)

const (
	TimeFormat = "2006-01-02"
)

func GetMVinf(host string, port string, user string, password string, date string, sid string, path1 string, path2 string, wg *sync.WaitGroup) {
	log.Println("Start the process Collecting MV information.")
	cmd := "BinPATHsymaccess -sid SID list view -detail > FilePATHDATE_SID_symaccess_list_view_detail.txt"
	cmd = strings.ReplaceAll(cmd, "BinPATH", path1)
	cmd = strings.ReplaceAll(cmd, "FilePATH", path2)
	cmd = strings.Replace(cmd, "DATE", date, 1)
	cmd = strings.ReplaceAll(cmd, "SID", sid)
	paramigo.InnerCommand(host, port, user, password, cmd)
	log.Println("Finish up the process Collecting MV information.")
	wg.Done()
}

func GetIGinf(host string, port string, user string, password string, date string, sid string, path1 string, path2 string, wg *sync.WaitGroup) {
	log.Println("Start the process Collecting IG information.")
	cmd := "for i in $(BinPATHsymaccess -sid SID list -type init | grep IG); do BinPATHsymaccess -sid SID -type initiator show $i -detail; done > FilePATHDATE_SID_symaccess_list_initiator.txt"
	cmd = strings.ReplaceAll(cmd, "BinPATH", path1)
	cmd = strings.ReplaceAll(cmd, "FilePATH", path2)
	cmd = strings.Replace(cmd, "DATE", date, 1)
	cmd = strings.ReplaceAll(cmd, "SID", sid)
	paramigo.InnerCommand(host, port, user, password, cmd)
	log.Println("Finish up the process Collecting IG information.")
	wg.Done()
}

func GetSGinf(host string, port string, user string, password string, date string, sid string, path1 string, path2 string, wg *sync.WaitGroup) {
	log.Println("Start the process Collecting SG information.")
	cmd := "BinPATHsymsg -sid SID list -detail > FilePATHDATE_SID_symsg_list_detail.txt"
	cmd = strings.ReplaceAll(cmd, "BinPATH", path1)
	cmd = strings.ReplaceAll(cmd, "FilePATH", path2)
	cmd = strings.Replace(cmd, "DATE", date, 1)
	cmd = strings.ReplaceAll(cmd, "SID", sid)
	paramigo.InnerCommand(host, port, user, password, cmd)
	log.Println("Finish up the process Collecting SG information.")
	wg.Done()
}

func GetEachSymdevinf(host string, port string, user string, password string, date string, sid string, path1 string, path2 string, wg *sync.WaitGroup) {
	log.Println("Start the process Collecting symdev information.")
	cmd := "for i in $(BinPATHsymdev -sid SID list -all | grep 0 | awk '{print $1}' | sed 1d); do BinPATHsymdev -sid SID show $i ;done > FilePATHDATE_SID_symdev_list.txt"
	cmd = strings.ReplaceAll(cmd, "BinPATH", path1)
	cmd = strings.ReplaceAll(cmd, "FilePATH", path2)
	cmd = strings.Replace(cmd, "DATE", date, 1)
	cmd = strings.ReplaceAll(cmd, "SID", sid)
	paramigo.InnerCommand(host, port, user, password, cmd)
	log.Println("Finish up the process Collecting symdev information.")
	wg.Done()
}

func GetCascadeSGinf(host string, port string, user string, password string, date string, sid string, path1 string, path2 string, wg *sync.WaitGroup) {
	log.Println("Start the process Collecting Cascade SG information.")
	cmd := "for i in $(BinPATHsymaccess -sid SID list -type storage -detail | sed '1,6d' | grep ' C' | awk '{print $1}' | grep -v '(S)'); do BinPATHsymsg -sid SID show $i ; done > FilePATHDATE_SID_symsg_cascade.txt"
	cmd = strings.ReplaceAll(cmd, "BinPATH", path1)
	cmd = strings.ReplaceAll(cmd, "FilePATH", path2)
	cmd = strings.Replace(cmd, "DATE", date, 1)
	cmd = strings.ReplaceAll(cmd, "SID", sid)
	paramigo.InnerCommand(host, port, user, password, cmd)
	log.Println("Finish up the process Collecting Cascade SG information.")
	wg.Done()
}

func CfgLogCollect(host, port, user, password, sid, path1, path2 string) {
	t := time.Now()
	formattedTime := t.Format(TimeFormat)
	date := strings.ReplaceAll(formattedTime, "-", "")

	var wg sync.WaitGroup
	wg.Add(5)
	go GetMVinf(host, port, user, password, date, sid, path1, path2, &wg)
	go GetIGinf(host, port, user, password, date, sid, path1, path2, &wg)
	go GetSGinf(host, port, user, password, date, sid, path1, path2, &wg)
	go GetEachSymdevinf(host, port, user, password, date, sid, path1, path2, &wg)
	go GetCascadeSGinf(host, port, user, password, date, sid, path1, path2, &wg)
	wg.Wait()
}
