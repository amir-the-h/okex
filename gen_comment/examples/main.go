package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	f, err := os.Open("/home/james/go/src/github.com/okex/examples/okapi-zh.html")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	//> result
	rex := regexp.MustCompile(`\w+`)

	m := make(map[string][]string)
	table := doc.Find("table")
	table.Each(func(i int, s *goquery.Selection) {
		s.Find("tbody").Each(func(i int, s *goquery.Selection) {
			s.Find("tr").Each(func(i int, s *goquery.Selection) {
				trs := make([]string, 0)
				s.Find("td").Each(func(i int, s *goquery.Selection) {
					trs = append(trs, s.Text())
				})
				if len(trs) == 0 {
					return
				}

				key := rex.FindString(trs[0])
				if key == "" {
					return
				}
				//fmt.Println(key)
				if _, ok := m[key]; !ok {

					m[key] = trs
					return
				}
				if len(trs) > 1 {
					var insert []string
				loop:
					for _, tr := range trs[:1] {
						for _, v := range m[key] {
							if v == tr {
								continue loop
							}
						}
						insert = append(insert, tr)
					}
					if len(insert) > 0 {
						m[key] = append(m[key], insert...)
					}

				}

			})
		})
	})
	//num := 0
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("comment.json", j, 0644)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range m {
		fmt.Printf("%s ----> %v\n", i, v)
		// num++
		// if num == 10 {
		// 	break
		// }
	}
}
