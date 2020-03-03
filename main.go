package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Item struct {
	Item []Ad `xml:"ad"`
}

type Ad struct {
	Url string `xml:"url"`
}

func UrlsCompare(urlSlice []string) {

	/*Comparar las urls de los distintos enlaces
	y obtener enlaces duplicados entre los archivos*/

	total := map[string]int{}

	for _, urls := range urlSlice {

		if total[urls] > 0 {
			total[urls] += 1

		} else {

			total[urls] = 1
		}
	}

	for url, veces := range total {

		if veces > 1 {

			fmt.Println("La url:", url, "esta repetida", veces, "veces")
		}
	}
}

func main() {

	files := []string{
		//"jobs_165.xml",
		//"jobs_7459.xml",
		//"jobs_8218.xml",
		//"jobs_8957.xml",
	}

	urls := []string{}

	for _, file := range files {

		data, _ := ioutil.ReadFile(file)

		ad := &Ad{}

		xml.Unmarshal(data, &ad)
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
				if inElement == "ad" {
					decoder.DecodeElement(&ad, &se)
					urls = append(urls, ad.Url)
				}
			default:
			}
		}

	}

	UrlsCompare(urls)
}
