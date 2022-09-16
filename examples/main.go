package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	doc, err := goquery.NewDocument("https://www.okx.com/docs-v5/zh/")
	if err != nil {
		log.Fatal(err)
	}
	//> result 
	rex:=regexp.MustCompile(`\w+`)
	
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

				key:=rex.FindString(trs[0])
				if key==""{
					return
				}
				//fmt.Println(key)
				if _, ok := m[key]; !ok {
					m[key] = trs
				}
				
			})
		})
	})
	//num := 0
	for i, v := range m {
		fmt.Printf("%s ----> %v\n", i, v)
		// num++
		// if num == 10 {
		// 	break
		// }
	}
}
