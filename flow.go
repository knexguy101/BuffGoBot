package main

import (
	"fmt"
	"github.com/knexguy101/BuffGo/buff/search"
	structs "github.com/knexguy101/BuffGo/models/client"
	"log"
	"strconv"
	"time"
)

func flow(itemId string, maxPrice float64, client *structs.HttpClient) {
	var (
		err error
		searchRes []search.SearchItem
	)
	for range time.Tick(5 * time.Second) {

		searchRes, err = search.Item(search.NewSearchFilter(itemId, search.PRICEASC, 1, true), client)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, v := range searchRes {

			fixedPrice, err := strconv.ParseFloat(v.Price, 64)
			if err != nil {
				continue
			}

			fmt.Println(fixedPrice)

			if fixedPrice <= maxPrice {
				//buy
				fmt.Println("[ITEM FOUND]", v.ID, v.Price, fmt.Sprintf("https://buff.163.com/goods/%d?from=market#tab=selling&page_num=1&sort_by=price.asc", v.GoodsID))
			}
		}

	}
}
