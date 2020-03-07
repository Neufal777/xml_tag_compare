package main

import (
	"log"
	"time"
)

func main() {

	//calculate exec time
	start := time.Now()

	ad := &Ad{}

	files := []string{
		"file1.xml",
		"file2.xml",
		//you can add as many as you want
	}

	//process files
	ad.filesProcess(files, "ad")

	//calculate exec time
	elapsed := time.Since(start)

	//print excution time
	log.Println("Exec Time", elapsed)

}
