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

/*
TO DO LIST
1. Cara penampilan stock**
2. Tunjuin 'Tidak ada product' di Display function (sudah)
3. Experiation Date nya  masih tahun(sudah)
4. go back in every choices (sementara)
5. limit tahunnya(sudah)
6. urusin spasinya njing (aman)

*/

// declare array list
var listProduct []Product

// display main function
func main() {
	var choose int
	display() //call display function
	fmt.Scan(&choose)
	for choose != 6 {
		//call function depend on choose
		if choose == 1 {
			add_product()
		} else if choose == 2 {
			display_product()
		} else if choose == 3 {
			remove_product()
		} else if choose == 4 {
			edit_product()
		} else if choose == 5 {
			check_expired_date()
		} else {
			fmt.Println("invalid input")
			continue
		}
		//sort the input or display it
		sortProduct()
		display() //call display function
		fmt.Scan(&choose)
	}

}

// to display welcome
func display() {
	fmt.Println("                    Welcome to Sembako Super App                    ")
	fmt.Println("1. Add Product ")
	fmt.Println("2. Display Product")
	fmt.Println("3. Remove Product")
	fmt.Println("4. Edit Product")
	fmt.Println("5. Check Product Expired Date")
	fmt.Println("6. exit Program")
	fmt.Println("What option would you like to choose ?")
}

func add_product() {
	//declare input variable
	var input Product

	//input the product
	fmt.Print("Enter your Product name : ")
	fmt.Scan(&input.name)
	fmt.Print("Enter Product type : ")
	fmt.Scan(&input.prod_type)
	input.exp_date = inputExpiredDate()
	fmt.Print("Stock : ")
	fmt.Scan(&input.stock)

	input.id = len(listProduct) + 1

	//add the input to the array listProduct
	listProduct = append(listProduct, input)

	fmt.Println("Product is succesfully stored")
	fmt.Println(" ")

}

func display_product() {
	//to display a bunch of product
	if len(listProduct) != 0 {
		for i := 0; i < len(listProduct); i++ {
			fmt.Println("Id : ", listProduct[i].id, "|", "Name : ", listProduct[i].name, "|", "Expired Date : ", listProduct[i].exp_date, "|", "Product Type : ", listProduct[i].prod_type, "|", "Stock : ", listProduct[i].stock, "|")
		}
		fmt.Println(" ")
	} else {
		fmt.Println("There is no product to display")
		fmt.Println(" ")
	}
}

func remove_product() {
	var choose int
	var duplicate []Product
	fmt.Print("What product ID would you like to remove? : ")
	fmt.Scan(&choose)

	idx := -1
	low := 0
	high := len(listProduct) - 1

	for low <= high {
		mid := (low + high) / 2

		if listProduct[mid].id == choose {
			idx = mid
			break
		} else if listProduct[mid].id < choose {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if idx != -1 {
		for i := 0; i < len(listProduct); i++ {
			if listProduct[i].id != idx+1 {
				duplicate = append(duplicate, listProduct[i])
			}
		}

		listProduct = duplicate
		fmt.Print("Product successfully removed")
		fmt.Println(" ")
	} else {
		fmt.Println("ID not found")
		fmt.Println(" ")
	}
}

func edit_product() {
	//declare chosen_id variable
	var chosen_id, choice, val int
	fmt.Print("choose product ID to edit : ")
	//scan product ID that want to be edited
	fmt.Scan(&chosen_id)
	//choose what product attribute want to be edited
	fmt.Println("what would you like to edit : ")
	fmt.Println("1. Edit Product name")
	fmt.Println("2. Edit Product Expired Date")
	fmt.Println("3. Edit product type")
	fmt.Println("4. add or remove stock")
	fmt.Println("5. Edit all")
	fmt.Println("6. Go back to main menu")
	fmt.Print("Your Choose : ")
	fmt.Scan(&choice)

	//depend on what n you scan
	switch choice {
	case 1:
		fmt.Print("Enter the new product name : ")
		fmt.Scan(&listProduct[chosen_id-1].name)
		fmt.Println("New product name already replaced")
		fmt.Println(" ")
	case 2:
		fmt.Print("Enter the new product expired date : ")
		fmt.Scan(&listProduct[chosen_id-1].exp_date)
		fmt.Println("New product expired date already replaced")
		fmt.Println(" ")
	case 3:
		fmt.Print("Enter the new product type : ")
		fmt.Scan(&listProduct[chosen_id-1].prod_type)
		fmt.Println("New product type date already replaced")
		fmt.Println(" ")
	case 4:
		fmt.Println("If you want to add stock type +(value)\n if you want to remove stock type -(value)")
		fmt.Print("Your Choice : ")
		fmt.Scan(&val)
		listProduct[chosen_id-1].stock += val
		fmt.Println("Stock has been updated")
		fmt.Println(" ")
	case 5:
		fmt.Print("Enter the new product name : ")
		fmt.Scan(&listProduct[chosen_id-1].name)
		fmt.Print("Enter the new product expired date : ")
		fmt.Scan(&listProduct[chosen_id-1].exp_date)
		fmt.Print("Enter the new product type : ")
		fmt.Scan(&listProduct[chosen_id-1].prod_type)
		fmt.Println("If you want to add stock type +(value)\n if you want to remove stock type -(value)")
		fmt.Print("Your Choice : ")
		fmt.Scan(&val)
		listProduct[chosen_id-1].stock += val
		fmt.Println("All product attribute already replaced")
		fmt.Println(" ")
	case 6:
		main()
	default:
		fmt.Println("Invalid input")
		fmt.Println(" ")
	}
}

func inputExpiredDate() string {
	var dateInput string
	var err error

	for {
		fmt.Print("Expired Date (YYYY-MM-DD): ")
		fmt.Scan(&dateInput)

		//validating input
		_, err = time.Parse("2006-01-02", dateInput)
		if err == nil {
			return dateInput // return if its valid
		}

		fmt.Println("Wrong format, try again!.")
	}
}

func check_expired_date() {
	if len(listProduct) == 0 {
		fmt.Println("No product can be check.")
		return
	}

	for _, product := range listProduct {
		expiredDate, err := time.Parse("2006-01-02", product.exp_date)
		if err != nil {
			fmt.Println("Wrong date format for product:", product.name)
			continue
		}

		now := time.Now()

		if now.After(expiredDate) {
			fmt.Printf("Product %s: has expired\n", product.name)
		} else {
			remaining := time.Until(expiredDate)
			days := int(remaining.Hours()) / 24
			hours := int(remaining.Hours()) % 24
			minutes := int(remaining.Minutes()) % 60

			fmt.Printf("Product %s: %d day %d hour %d minute \n", product.name, days, hours, minutes)
		}
	}
}

func sortProduct() {
	n := len(listProduct)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if listProduct[j].id > listProduct[j+1].id {
				// Tukar posisi jika id lebih besar
				listProduct[j], listProduct[j+1] = listProduct[j+1], listProduct[j]
			}
		}
	}
}
