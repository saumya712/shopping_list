package operation

import (
	"INNVENTORY/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Searchbyid(products *[]model.Product) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the id you wanna search:")
	idstr, _ := reader.ReadString('\n')
	idstr = strings.TrimSpace(idstr)
	id, _ := strconv.ParseInt(idstr, 10, 64)

	for _, p := range *products {
		if p.Id == id {
			fmt.Println("product found")
			fmt.Printf("ID: %d,NAME: %s,PRICE: %d,QUANTITY: %d", p.Id, p.Name, p.Price, p.Quantity)
			return
		}
	}
	fmt.Println("product not found")
}
