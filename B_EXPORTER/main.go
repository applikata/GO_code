package main

import (
	"exporter/src"
	"fmt"
)

func main() {

	conf := src.ReadConf() //Get configuration exporter

	logFiles := src.GetFiles(conf.Exporter.Log_Path, conf.Exporter.Prefix) //Get files for metrics collection

	for i, f := range logFiles {
		fmt.Println("index = ", i, " files = ", f)
	}

	counter := src.FindErr(logFiles, conf)
	fmt.Println(counter)
}
