package operation

import (
	"INNVENTORY/model"
	"fmt"
)

func Listall(products *[]model.Product) {
	if len(*products) == 0 {
		fmt.Println("sorry the list empty!!!!")
		return
	}
	fmt.Println("your list is:")
	for _, p := range *products {
		fmt.Printf("ID: %d,NAME: %s,PRICE: %d,QUANTITY: %d", p.Id, p.Name, p.Price, p.Quantity)
	}
}
