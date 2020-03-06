package main

import (
	"log"
	"time"
)

func main() {

	start := time.Now()

	ad := &Ad{}
	files := []string{
		"file1.xml",
		"file2.xml",
	}

	var urls []string = ad.filesProcess(files, "ad")

	ad.TagCompare(urls)

	elapsed := time.Since(start)
	log.Println("Exec Time", elapsed)

}
