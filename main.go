package main

import (
	"fmt"
	"time"
)

// declare struct variable
type Product struct {
	id        int
	name      string
	exp_date  string
	prod_type string
	stock     int
}

func main() {
	// IS: No input
	// FS: Runs the program loop and processes user input to manage product list
	var choose int
	var listProduct []Product
	display() //call display function
	fmt.Scan(&choose)
	if choose == 0 {
		display()
		fmt.Scan(&choose)
	}
	for choose != 6 {

		if choose == 0 {
			continue
		}
		switch choose {
		//call function depend on choose
		case 1:
			add_product(&listProduct)
		case 2:
			display_product(listProduct)
		case 3:
			remove_product(&listProduct)
		case 4:
			edit_product(&listProduct)
		case 5:
			check_expired_date(listProduct)
		default:
			fmt.Println("Invalid input")
		}
		//sort the input or display it
		sortProduct(&listProduct)
		display() //call display function
		fmt.Scan(&choose)
	}
}

// to display welcome
func display() {
	// IS: No input
	// FS: Displays the main menu
	fmt.Println("\n==============================")
	fmt.Println("     Sembako Super App")
	fmt.Println("==============================")
	fmt.Println("1. Add Product")
	fmt.Println("2. Display Product")
	fmt.Println("3. Remove Product")
	fmt.Println("4. Edit Product")
	fmt.Println("5. Check Product Expired Date")
	fmt.Println("6. Exit Program")
	fmt.Print("Choose an option (0 to redisplay menu): ")
}

func add_product(listProduct *[]Product) {
	// IS: User inputs name, type, expiration date, and stock
	// FS: Adds the new product to the listProduct
	var input Product
	fmt.Print("Enter product name : ")
	fmt.Scan(&input.name)
	fmt.Print("Enter product type : ")
	fmt.Scan(&input.prod_type)
	input.exp_date = inputExpiredDate()
	fmt.Print("Enter stock : ")
	fmt.Scan(&input.stock)

	input.id = len(*listProduct) + 1

	//add the input to the array listProduct
	*listProduct = append(*listProduct, input)

	fmt.Println("Product added successfully.\n")
}

//to display a bunch of product
func display_product(listProduct []Product) {
	// IS: Receives the list of products
	// FS: Displays all product details or shows message if list is empty
	if len(listProduct) == 0 {
		fmt.Println("No products to display.\n")
		return
	}
	fmt.Println("\nList of Products:")
	for i := 0; i < len(listProduct); i++ {
		fmt.Printf("ID: %d | Name: %s | Type: %s | Exp: %s | Stock: %d\n",
			listProduct[i].id, listProduct[i].name, listProduct[i].prod_type, listProduct[i].exp_date, listProduct[i].stock)
	}
	fmt.Println()
}

func remove_product(listProduct *[]Product) {
	// IS: User inputs product ID
	// FS: Removes product with matching ID using binary search logic
	var choose int
	var duplicate []Product
	fmt.Print("What product ID would you like to remove? : ")
	fmt.Scan(&choose)
	if choose == 0 {
		display()
		return
	}

	idx := -1
	low := 0
	high := len(*listProduct) - 1

	for low <= high {
		mid := (low + high) / 2
		if (*listProduct)[mid].id == choose {
			idx = mid
			break
		} else if (*listProduct)[mid].id < choose {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if idx != -1 {
		for i := 0; i < len(*listProduct); i++ {
			if (*listProduct)[i].id != choose {
				duplicate = append(duplicate, (*listProduct)[i])
			}
		}
		*listProduct = duplicate
		fmt.Print("Product successfully removed")
		fmt.Println(" ")
	} else {
		fmt.Println("ID not found")
		fmt.Println(" ")
	}
}

func edit_product(listProduct *[]Product) {
	// IS: User inputs a product ID and a field to edit
	// FS: Updates the selected product field(s) or removes it if stock reaches zero
	
	var chosen_id, choice, val int//declare chosen_id variable
	fmt.Print("Choose product ID to edit : ")//scan product ID that want to be edited
	fmt.Scan(&chosen_id)//choose what product attribute want to be edited
	if chosen_id == 0 {
		display()
		return
	}

	// Find index by matching ID
	index := -1
	for i := 0; i < len(*listProduct); i++ {
		if (*listProduct)[i].id == chosen_id {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Product ID not found.")
		return
	}

	fmt.Println("What would you like to edit : ")
	fmt.Println("1. Edit Product name")
	fmt.Println("2. Edit Product Expired Date")
	fmt.Println("3. Edit product type")
	fmt.Println("4. Add or remove stock")
	fmt.Println("5. Edit all")
	fmt.Println("6. Go back to main menu")
	fmt.Print("Your Choose : ")
	fmt.Scan(&choice)
	if choice == 0 {
		display()
		return
	}

	switch choice {
	case 1:
		fmt.Print("Enter the new product name : ")
		fmt.Scan(&(*listProduct)[index].name)
		fmt.Println("New product name already replaced")
		fmt.Println(" ")
	case 2:
		fmt.Print("Enter the new product expired date : ")
		fmt.Scan(&(*listProduct)[chosen_id-1].exp_date)
		fmt.Println("New product expired date already replaced")
		fmt.Println(" ")
	case 3:
		fmt.Print("Enter the new product type : ")
		fmt.Scan(&(*listProduct)[chosen_id-1].prod_type)
		fmt.Println("New product type already replaced")
		fmt.Println(" ")
	case 4:
		fmt.Println("If you want to add stock type +(value), \nIf you want to remove stock type -(value)")
		fmt.Print("Your Choice : ")
		fmt.Scan(&val)
		(*listProduct)[chosen_id-1].stock += val
		fmt.Println("Stock has been updated")
		fmt.Println(" ")
		if (*listProduct)[chosen_id-1].stock <= 0 {
			fmt.Println("Stock is 0, product will be automatically removed.")
			remove_product(listProduct)
		}
	case 5:
		fmt.Print("Enter the new product name : ")
		fmt.Scan(&(*listProduct)[chosen_id-1].name)
		fmt.Print("Enter the new product expired date : ")
		fmt.Scan(&(*listProduct)[chosen_id-1].exp_date)
		fmt.Print("Enter the new product type : ")
		fmt.Scan(&(*listProduct)[chosen_id-1].prod_type)
		fmt.Println("If you want to add stock type +(value)\nIf you want to remove stock type -(value)")
		fmt.Print("Your Choice : ")
		fmt.Scan(&val)
		(*listProduct)[chosen_id-1].stock += val
		fmt.Println("All product attribute already replaced")
		fmt.Println(" ")
		if (*listProduct)[chosen_id-1].stock <= 0 {
			fmt.Println("Stock is 0, product will be automatically removed.")
			remove_product(listProduct)
		}
	case 6:
		display()
		return
	default:
		fmt.Println("Invalid input")
		fmt.Println(" ")
	}
}

func inputExpiredDate() string {
	// IS: User inputs date as string
	// FS: Returns validated date in YYYY-MM-DD format
	var dateInput string
	for {
		fmt.Print("Enter expiration date (YYYY-MM-DD): ")
		fmt.Scan(&dateInput)
		_, err := time.Parse("2006-01-02", dateInput)
		if err == nil {
			return dateInput // return if its valid
		}
		fmt.Println("Invalid format. Please try again.")
	}
}

func check_expired_date(listProduct []Product) {
	// IS: Receives the list of products
	// FS: Checks and prints if each product is expired or how long until expiry
	if len(listProduct) == 0 {
		fmt.Println("No products to check. ")
		return
	}

	fmt.Println("Product Expiration Check:")
	for i := 0; i < len(listProduct); i++ {
		expDate, err := time.Parse("2006-01-02", listProduct[i].exp_date)
		if err != nil {
			fmt.Println("Invalid date for product:", listProduct[i].name)
			continue
		}
		if time.Now().After(expDate) {
			fmt.Printf("Product %s has expired.", listProduct[i].name)
		} else {
			dur := time.Until(expDate)
			days := int(dur.Hours()) / 24
			hours := int(dur.Hours()) % 24
			minutes := int(dur.Minutes()) % 60
			fmt.Printf("Product %s expires in %d days %d hours %d minutes.", listProduct[i].name, days, hours, minutes)
		}
	}
	fmt.Println()
}

func sortProduct(listProduct *[]Product) {
	// IS: Receives pointer to list of products
	// FS: Sorts product list in ascending order by ID
	n := len(*listProduct)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if (*listProduct)[j].id > (*listProduct)[j+1].id {
				// Tukar posisi jika id lebih besar
				(*listProduct)[j], (*listProduct)[j+1] = (*listProduct)[j+1], (*listProduct)[j]
			}
		}
	}
}
