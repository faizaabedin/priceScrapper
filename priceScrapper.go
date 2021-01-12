package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnXML("/html/body/div[3]/div[1]/div/div[3]/div[2]/div[3]/div[2]/div[1]/span/strong", func(e *colly.XMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit("https://www.canadacomputers.com/product_info.php?cPath=6_1937&item_id=105266")
}

