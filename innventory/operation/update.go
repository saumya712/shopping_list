package operation

import (
	"INNVENTORY/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Update(product *[]model.Product) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the id of the product u wanna update")
	idstr, _ := reader.ReadString('\n')
	idstr = strings.TrimSpace(idstr)
	id, _ := strconv.ParseInt(idstr, 10, 64)

	for i := range *product {
		if (*product)[i].Id == id {
			var opt, price, quantity int64
			var name string

			fmt.Println("what do u wanna update")
			fmt.Println("1. name")
			fmt.Println("2. price")
			fmt.Println("3. quantity")
			fmt.Scanln(&opt)
			switch opt {
			case 1:
				fmt.Println("enter the new name")
				fmt.Scan(&name)
				(*product)[i].Name = name
				return
			case 2:
				fmt.Println("enter the new price")
				fmt.Scan(&price)
				(*product)[i].Price = price
				return
			case 3:
				fmt.Println("enter the new quantity")
				fmt.Scan(&quantity)
				(*product)[i].Quantity = quantity
				return
			}
		}
	}
}
