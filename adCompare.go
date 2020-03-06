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

func (a *Ad) filesProcess(files []string, adTag string) []string {

	/*
		We analyze the files and get the tags
		to process them later and find matches.
	*/

	var urls = []string{}

	adsprocessed := 0

	for _, file := range files {

		data, _ := ioutil.ReadFile(file)

		xml.Unmarshal(data, &a)

		r := bytes.NewReader(data)

		var inElement string
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
					adsprocessed++
				}
			}
		}
	}

	fmt.Println("Ads processed", adsprocessed)
	return urls
}

//TagCompare func description
func (a *Ad) tagCompare(tagList []string) {

	/*
		We make the comparison of the tag list
		to find matches in the parameters you pass
		any type of list to make the comparison
	*/

	total := map[string]int{}
	var totalSum int

	for _, tag := range tagList {

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
