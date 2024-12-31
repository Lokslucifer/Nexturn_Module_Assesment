package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

var inventory []Product

func AddProduct(id int, name string, priceInput string, stock int) error {
	price, err := strToFloat64(priceInput)
	if err != nil {
		return err
	}

	for _, product := range inventory {
		if product.ID == id {
			return errors.New("Product ID already exists")
		}
	}

	newProduct := Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}
	inventory = append(inventory, newProduct)
	return nil
}

func UpdateStock(id int, stock int) error {
	if stock < 0 {
		return errors.New("Stock cannot be negative")
	}

	for i, product := range inventory {
		if product.ID == id {
			inventory[i].Stock = stock
			return nil
		}
	}
	return errors.New("Product not found")
}

func SearchProduct(query string) (*Product, error) {
	for _, product := range inventory {
		if product.Name == query || fmt.Sprintf("%d", product.ID) == query {
			return &product, nil
		}
	}
	return nil, errors.New("Product not found")
}

func DisplayInventory() {
	if len(inventory) == 0 {
		fmt.Println("No products in inventory.")
		return
	}
	fmt.Println("ID | Name         | Price  | Stock")
	for _, product := range inventory {
		fmt.Printf("%d  | %-12s | Rs.%.2f | %d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

func strToFloat64(input string) (float64, error) {
	var price float64
	_, err := fmt.Sscanf(input, "%f", &price)
	if err != nil {
		return 0, errors.New("Invalid price input")
	}
	return price, nil
}

func main() {
	for {
		var choice int
		fmt.Println("\nMenu:")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Stock")
		fmt.Println("3. Search Product")
		fmt.Println("4. Display Inventory")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scanf("%d", &choice)
		fmt.Scanln()

		switch choice {
		case 1:
			var id, stock int
			var name, priceInput string
			fmt.Print("Enter Product ID: ")
			fmt.Scanf("%d", &id)
			fmt.Scanln()
			fmt.Print("Enter Product Name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter Product Price: ")
			fmt.Scanln(&priceInput)
			fmt.Print("Enter Product Stock: ")
			fmt.Scanf("%d", &stock)
			fmt.Scanln()

			err := AddProduct(id, name, priceInput, stock)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Product added successfully!")
			}

		case 2:
			var id, stock int
			fmt.Print("Enter Product ID: ")
			fmt.Scanf("%d", &id)
			fmt.Scanln()
			fmt.Print("Enter New Stock: ")
			fmt.Scanf("%d", &stock)
			fmt.Scanln()

			err := UpdateStock(id, stock)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Stock updated successfully!")
			}

		case 3:
			var query string
			fmt.Print("Enter Product ID or Name to search: ")
			fmt.Scanln(&query)
			fmt.Scanln()

			product, err := SearchProduct(query)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Product Found: ID: %d, Name: %s, Price: Rs.%.2f, Stock: %d\n",
					product.ID, product.Name, product.Price, product.Stock)
			}

		case 4:
			DisplayInventory()

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
