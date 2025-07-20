package operation

import (
	"INNVENTORY/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Addproduct(products *[]model.Product) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the product id:")
	idstr, _ := reader.ReadString('\n')
	idstr = strings.TrimSpace(idstr)
	id, _ := strconv.ParseInt(idstr, 10, 64)

	fmt.Println("enter the product name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("enter the product price:")
	pricestr, _ := reader.ReadString('\n')
	pricestr = strings.TrimSpace(pricestr)
	price, _ := strconv.ParseInt(pricestr, 10, 64)

	fmt.Println("enter the product quantity:")
	qtystr, _ := reader.ReadString('\n')
	qtystr = strings.TrimSpace(qtystr)
	quantity, _ := strconv.ParseInt(qtystr, 10, 64)

	newproduct := model.Product{
		Id:       id,
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
	*products = append(*products, newproduct)
	fmt.Println("product added")
}
