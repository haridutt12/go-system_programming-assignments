package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	filename := filepath.Base(os.Args[0])
	// defined logging facility
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, filename)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	// defined logging priority
	sysLog.Crit("crit : logging in go")

	sysLog, err = syslog.New(syslog.LOG_ALERT|syslog.LOG_LOCAL7, "u can define multiple syslog.new() in same programme")
	if err != nil {
		fmt.Println(err)
		log.Fatal(sysLog)
	}

	sysLog.Emerg("emer : logging in go")

	fmt.Fprintf(sysLog, "log.print : logging in go!")

}
