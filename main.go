package main

import (
	"log"
	"time"
)

func main() {

	start := time.Now()

	ad := &Ad{}
	files := []string{
		//"file1.xml",
		"file2.xml",
		//"file3.xml",
		//etc..
	}

	var urls, titles []string = ad.filesProcess(files, "ad")

	ad.TagCompare(urls)
	ad.TagCompare(titles)

	elapsed := time.Since(start)
	log.Println("Exec Time", elapsed)

}
