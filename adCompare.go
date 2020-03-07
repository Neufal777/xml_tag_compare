package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

//Ad type description
type Ad struct {
	URL string `xml:"url"`
}

func (a *Ad) filesProcess(files []string, adTag string) {

	/*
		We analyze the files and get the tags
		to process them later and find matches.
	*/
	fmt.Println("Processing..")

	var urls = []string{}
	total := map[string]int{}
	var totalSum int
	var urlprocessed int = 0
	var inElement string

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
	fmt.Println("url processed", urlprocessed)
	fmt.Println("Comparing Urls...")

	//check if the URL's are duplicated
	for _, tag := range urls {

		if total[tag] > 0 {
			total[tag]++
		} else {
			total[tag] = 1
		}
	}

	for tag, veces := range total {
		if veces > 1 {
			fmt.Println(tag, "Duplicated", veces, "Times")
			totalSum += veces
		}
	}

	fmt.Println("Total duplicates", totalSum)

}
