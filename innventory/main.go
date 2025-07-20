package main

import (
	"INNVENTORY/model"
	"INNVENTORY/operation"
	"INNVENTORY/selectmsg"
	"fmt"
)

func main() {

	var opt int
	var product []model.Product
	for {
		selectmsg.ShowMenu()
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			operation.Addproduct(&product)
		case 2:
			operation.Listall(&product)
		case 3:
			operation.Update(&product)
		case 4:
			operation.Searchbyid(&product)
		case 5:
			return
		default:
			fmt.Println("invalid choice")

		}
	}
}
