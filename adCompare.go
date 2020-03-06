package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

//Ad type description //TODO
type Ad struct {
	URL   string `xml:"url"`
	ID    int    `xml:"id"`
	Title string `xml:"title"`
}

func (a *Ad) filesProcess(files []string, adTag string) ([]string, []string) {

	/*
		Analizamos los archivos y obtenemos los tags
		para procesarlas mas tarde y encontrar coincidencias.
	*/

	var urls = []string{}
	var titles = []string{}

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

					//Tag values we get
					urls = append(urls, a.URL)
					titles = append(titles, a.Title)

					adsprocessed++
				}
			default:
			}
		}
	}

	fmt.Println("Ads processed", adsprocessed)
	return urls, titles
}

//TagCompare func description //TODO
func (a *Ad) TagCompare(tagList []string) {

	/*
		Hacemos la comparacion de la lista de tags
		para encontrar coincidencias

		en los parámetros se le pasa cualquier tipo de lista para hacer la comparacion
	*/

	total := map[string]int{}
	var sumaTotal int

	for _, tag := range tagList {

		if total[tag] > 0 {
			total[tag]++
		} else {
			total[tag] = 1
		}
	}

	for tag, veces := range total {

		if veces > 1 {

			fmt.Println(tag, "Duplicado", veces, "veces")
			sumaTotal += veces
		}
	}

	fmt.Println("Total duplicados", sumaTotal)

}
