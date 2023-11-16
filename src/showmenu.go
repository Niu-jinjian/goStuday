package main

import "fmt"

var TheMenu []map[string]float64

func SetMenu(themenu *[]map[string]float64, dishes string, price float64) []map[string]float64 {
	if *themenu == nil {
		*themenu = make([]map[string]float64, 0)
	}
	*themenu = append(*themenu, map[string]float64{dishes: price})
	return *themenu
}

func main() {
	var menu []map[string]float64
	SetMenu(&menu, "锅包肉", 68.8)
	SetMenu(&menu, "水煮鱼", 39.8)
	SetMenu(&menu, "地三鲜", 18.8)

	for _, showmenu := range menu {
		for dishes, price := range showmenu {
			fmt.Printf("菜品: %s, 价格: %.2f \n", dishes, price)
		}
	}
}
