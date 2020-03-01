package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

type Ad struct {
	Url string `xml:"url"`
}

func main() {

	files := []string{
		"file1.xml",
		"file2.xml",
		"file3.xml",
	}

	urls := []string{}

	for _, file := range files {

		data, _ := ioutil.ReadFile(file)

		ad := &Ad{}

		_ = xml.Unmarshal([]byte(data), &ad)

		urls = append(urls, ad.Url+"&"+file)

	}

	//print slice of ALL oobtained ulrs
	fmt.Println(urls)

	urlsList1 := []string{"url1"}
	urlsList2 := []string{"url2"}
	urlsList3 := []string{"url3"}

	for _, url := range urls {

		contiene1 := strings.Contains(url, "&file1.xml")
		contiene2 := strings.Contains(url, "&file2.xml")
		contiene3 := strings.Contains(url, "&file3.xml")

		if contiene1 == true {
			urlsList1 = append(urlsList1, url)
		} else if contiene2 == true {
			urlsList2 = append(urlsList2, url)
		} else if contiene3 == true {
			urlsList3 = append(urlsList3, url)

		}
	}

}
