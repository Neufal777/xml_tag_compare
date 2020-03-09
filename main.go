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
		//"file2.xml",
		//"file3.xml",
		//you can add as many as you want
	}

	//process files & get the map[string]int with all the duplicates
	res := ad.filesProcess(files, "ad")

	//show results
	ad.showDuplicates(res)

	//calculate exec time
	elapsed := time.Since(start)

	//print excution time
	log.Println("Exec Time", elapsed)

}
