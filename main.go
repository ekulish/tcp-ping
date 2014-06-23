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
	if err != nil {
		log.Fatal(err)
		return
	}
	c := make(chan string)

	//for {
	go func() {
		conn, err := net.DialTimeout("tcp", "172.17.0.191:9000", time.Duration(10*time.Second))
		if err != nil {
			// log error
			c <- fmt.Sprint("Error: ", err)
			return
		}
		// log success
		c <- fmt.Sprint("Connected to ", conn.RemoteAddr())
		defer conn.Close()
	}()
	time.Sleep(time.Second * 60)
	//}

	printLog(<-c)

	defer logFile.Close()
}
