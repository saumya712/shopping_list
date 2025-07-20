package selectmsg

import "fmt"

func ShowMenu() {
	fmt.Println("\n========= INVENTORY MENU =========")
	fmt.Println("1. Add Product")
	fmt.Println("2. List Products")
	fmt.Println("3. Update Product")
	fmt.Println("4. Search Product")
	fmt.Println("5. Exit")
	fmt.Print("Enter your choice: ")
}
