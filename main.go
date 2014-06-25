// main
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

// define flags
var flagHost string
var flagPort int
var flagTimeout int
var flagPeriod int
var flagFile string

var logFile *os.File

func printLog(message string) {
	if flagFile != "" {
		fmt.Fprintln(logFile, time.Now().Format("2006-01-02 15:04:05"), message)
	} else {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), message)
	}
}

func init() {
	flag.StringVar(&flagHost, "h", "127.0.0.1", "Host name or IP address")
	flag.IntVar(&flagPort, "p", 80, "Host port")
	flag.IntVar(&flagTimeout, "t", 10, "Connection timeout (in seconds)")
	flag.IntVar(&flagPeriod, "r", 60, "Dial period (in seconds)")
	flag.StringVar(&flagFile, "f", "", "File name which is used instead of printing results")
}

func main() {
	var err error

	flag.Parse()
	// create/open file log, if file name presents as a command line argument
	if flagFile != "" {
		logFile, err = os.Create(flagFile)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer logFile.Close()
	}

	c := make(chan string)
	host := fmt.Sprintf("%v:%v", flagHost, flagPort)

	// execute dail in separate routine to the defined host:port until program is killed
	for {
		go func() {
			conn, err := net.DialTimeout("tcp", host, time.Duration(flagTimeout)*time.Second)
			if err != nil {
				// log error
				c <- fmt.Sprint("Error: ", err)
				return
			}
			defer conn.Close()

			// log success
			c <- fmt.Sprint("Connected to ", conn.RemoteAddr())
		}()

		printLog(<-c)
		time.Sleep(time.Duration(flagPeriod) * time.Second)
	}
}
