package main

import (
	"log"
	"time"

	"github.com/CompareFiles/domain"
)

func main() {

	//calculate exec time
	start := time.Now()

	files := []string{
		"file1.xml",
		"file2.xml",
		//you can add as many as you want
	}

	//process files & get the map[string]int with all the duplicates
	res := domain.FilesProcess(files, "ad")
	dups := domain.CheckDuplicates(res)

	//print all duplicates
	domain.ShowDuplicates(dups)

	//calculate exec time
	elapsed := time.Since(start)

	//print excution time
	log.Println("Exec Time", elapsed)

}
