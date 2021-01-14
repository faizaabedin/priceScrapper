package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
)

func main() {
	c := colly.NewCollector()
	c.OnXML("/html/body/div[3]/div[1]/div/div[3]/div[2]/div[3]/div[2]/div[1]/span/strong", func(e *colly.XMLElement) {
		match, _ := regexp.MatchString("\\d+.\\d+", e.Text)
		if match{
			re := regexp.MustCompile("\\d+.\\d+")
			price := re.FindStringSubmatch(e.Text)
			fmt.Println(price[0])
		}
	})

	c.Visit("https://www.canadacomputers.com/product_info.php?cPath=6_1937&item_id=105266")
}

