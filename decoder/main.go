package main

import (
	"flag"
	"log"
	"runtime"
	"strings"
	"time"
	"strconv"
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
)

var (
	in_addr  = flag.String("in", "", "incoming socket")
	cpu      = flag.Int("cpu", 1, "how much cores to use")
)

func worker(source <-chan RawData, decoded chan<- string, timers chan<- time.Duration) {
	for {
		select {
		case pb := <-source:
			start := time.Now()
			metrics, err := Decode(pb.Timestamp, pb.Data)
			if err != nil {
				log.Printf("[Decoder] Error decoding protobuf packet: %v", err)
				continue
			}
			for _, m := range metrics {
				decoded <- m
			}
			timers <- time.Since(start)
		}
	}
}

func main() {
	flag.Parse()

	log.Printf("Using %d/%d CPU\n", *cpu, runtime.NumCPU())
	runtime.GOMAXPROCS(*cpu)

	listener := NewListener(in_addr)
        db, err := sql.Open("mysql", "username:password@/database")
            if err != nil {
                log.Printf("[Decoder] Cannot connct to MySQL: %v", err)
            }
	defer db.Close()
        stmtIns, err := db.Prepare("INSERT INTO phpload VALUES( ?, ?, ?, ?, ?, '1' ) ON DUPLICATE KEY UPDATE cpu = cpu + VALUES(cpu), cnt = cnt + VALUES(cnt)")
         if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
	defer stmtIns.Close()
	var decoding_time time.Duration
	var decoded_count, sent_count int64
	decoded := make(chan string, 10000)
	timers := make(chan time.Duration, 1000)
	ticker := time.NewTicker(time.Second)

	// Let's go!
	for i := 0; i < *cpu; i++ {
		go worker(listener.RawMetrics, decoded, timers)
	}
	go listener.Start()

	for {
		select {
		case <-ticker.C:
			log.Printf("[Decoder] Decoded %v (%v), sent %v", decoded_count, decoding_time, sent_count)
			sent_count = 0
			decoded_count = 0
			decoding_time = 0

		case metric := <-decoded:
			sent_count += 1
			data := strings.Split(metric, " ")
			unixtime := data[1]
			curtime, _ := strconv.Atoi(unixtime)
			cpu := data[4]
			scriptname := data[6]
			date:= time.Unix(int64(curtime), 0).Format("2006-1-2:15:04:05")
			tempdate := strings.Split(date, ":")
			currentdate := tempdate[0]
			currenthour := tempdate[1]
			//we have 3 different hosting-environment:
			//1. own design
			//2. ISP-hosting
			//3. Cpanel-hosting
			//All users in /home, or /var/www - for ISP-hosting
			//p0 - its our system-user and we exclude him from global-stat
			//this is my first go-coding and my modification like a shit:)
			if strings.HasPrefix(scriptname, "/home/p") && !strings.Contains(scriptname, "/public_html/") {
				tempscriptname := strings.Split(scriptname, "/")
				tempuser := tempscriptname[2]
				user := strings.TrimLeft(tempuser, "p")
				newscriptname := strings.TrimLeft(scriptname, "/home/")
				newscriptname2 := strings.TrimLeft(newscriptname, tempuser)
				newscriptname3 := strings.TrimLeft(newscriptname2, "www/")
				if strings.HasPrefix(scriptname, "/home/p0/") {
					log.Printf("[Decoder] p0 detected, doesnt send p0 to database")
				} else {
					stmtIns.Exec(currentdate, currenthour, user, newscriptname3, cpu)
				}
			} else if strings.HasPrefix(scriptname, "/var/www/p") {
                tempscriptname := strings.Split(scriptname, "/")
                tempuser := tempscriptname[3]
                user := strings.TrimLeft(tempuser, "p")
                newscriptname := strings.TrimLeft(scriptname, "/var/www/")
                newscriptname2 := strings.TrimLeft(newscriptname, tempuser)
				newscriptname3 := strings.TrimLeft(newscriptname2, "data/")
                if strings.HasPrefix(scriptname, "/var/www/p0/") {
                      log.Printf("[Decoder] p0 detected, doesnt send p0 to database")
                } else {
                      stmtIns.Exec(currentdate, currenthour, user, newscriptname3, cpu)
                }
			} else if strings.HasPrefix(scriptname, "/home/p") && strings.Contains(scriptname, "/public_html/") {
                tempscriptname := strings.Split(scriptname, "/")
                tempuser := tempscriptname[3]
                user := strings.TrimLeft(tempuser, "p")
                newscriptname := strings.TrimLeft(scriptname, "/home/")
                newscriptname2 := strings.TrimLeft(newscriptname, tempuser)
                if strings.HasPrefix(scriptname, "/home/p0/") {
                       log.Printf("[Decoder] p0 detected, doesnt send p0 to database")
                } else {
                       stmtIns.Exec(currentdate, currenthour, user, newscriptname2, cpu)
                }
			}
			//log.Printf("[Decoder] time: %v cpu: %v script: %v", currenthour, cpu, scriptname)

		case t := <-timers:
			decoding_time += t
			decoded_count += 1
		}
	}
}
