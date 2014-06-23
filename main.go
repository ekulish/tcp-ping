// main
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var logFile *os.File

func printLog(message string) {
	fmt.Fprintln(logFile, time.Now().Format("2006-01-02 15:04:05"), message)
}

func main() {
	//conn, err := tls.Dial("tcp", "172.17.0.191:9000", &config)
	var err error
	logFile, err = os.Create("ping_log.txt")
	defer logFile.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	c := make(chan string)

	for {
		go func() {
			//conn, err := net.DialTimeout("tcp", "172.17.0.191:9000", time.Duration(10*time.Second))
			//198.199.122.230
			conn, err := net.DialTimeout("tcp", "198.199.122.230:22", time.Duration(10*time.Second))
			defer conn.Close()

			if err != nil {
				// log error
				c <- fmt.Sprint("Error: ", err)
				return
			}
			// log success
			c <- fmt.Sprint("Connected to ", conn.RemoteAddr())
		}()

		printLog(<-c)

		time.Sleep(time.Second * 60)
	}
}
