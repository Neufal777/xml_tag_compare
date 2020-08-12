package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"time"
	//"github.com/xmlTagCompare/domain"
)

//Ad type description
type Ad struct {
	URL string `xml:"url"`
}

//FilesProcess process the files and finds the tag X
func FilesProcess(files []string, adTag string) []string {

	/*
		We analyze the files and get the tags
		to process them later and find matches.
	*/

	a := &Ad{}
	//Show processing message
	fmt.Println("Processing..")

	var (
		urlprocessed int
		inElement    string
		urls         []string
	)

	for _, file := range files {

		data, _ := ioutil.ReadFile(file)

		xml.Unmarshal(data, &a)

		r := bytes.NewReader(data)

		decoder := xml.NewDecoder(r)

		for {
			t, _ := decoder.Token()
			if t == nil {
				break
			}

			switch se := t.(type) {
			case xml.StartElement:
				inElement = se.Name.Local
				if inElement == adTag {
					decoder.DecodeElement(&a, &se)
					urls = append(urls, a.URL)
					urlprocessed++
				}
			}
		}
	}
	fmt.Println(adTag, "processed", urlprocessed)

	//return slice with all url dups
	return urls

}

//CheckDuplicates finds duplicates
func CheckDuplicates(allUrls []string) map[string]int {

	totalUrls := map[string]int{}

	//check duplicates in slice return map with all the dups
	for _, tagUrl := range allUrls {

		if totalUrls[tagUrl] > 0 {
			totalUrls[tagUrl]++
		} else {
			totalUrls[tagUrl] = 1
		}
	}

	return totalUrls

}

//ShowDuplicates show all th duplicates
func ShowDuplicates(tagmap map[string]int) {

	var (
		sumTotalDup int
	)

	for tag, times := range tagmap {
		if times > 1 {
			fmt.Println(tag, "Duplicated", times, "Times")
			sumTotalDup += times
		}
	}

	fmt.Println("Total duplicates: ", sumTotalDup)
}

func main() {

	//calculate exec time
	start := time.Now()

	files := []string{
		"1.xml",
		"2.xml",
		"3.xml",
		"4.xml",
		"5.xml",
		"6.xml",
		//you can add as many as you want
	}

	//process files & get the map[string]int with all the duplicates
	res := FilesProcess(files, "ad")
	dups := CheckDuplicates(res)

	//print all duplicates
	ShowDuplicates(dups)

	//calculate exec time
	elapsed := time.Since(start)

	//print excution time
	log.Println("Exec Time", elapsed)

}
