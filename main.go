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
		//you can add as many as you want
	}

	files2 := []string{
		"file2.xml",
		//you can add as many as you want
	}

	//process files
	go ad.filesProcess(files, "ad")
	go ad.filesProcess(files2, "ad")

	//calculate exec time
	elapsed := time.Since(start)

	//print excution time
	log.Println("Exec Time", elapsed)

}
