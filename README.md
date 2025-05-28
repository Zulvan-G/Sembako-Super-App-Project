# Sembako Super App – Final Project

This is a terminal based inventory management system written in Go. It helps manage a simple sembako (daily needs) product inventory.
---

## Features

- Add product with name, type, stock, and expiration date
- Display all products
- Remove product by ID
- Edit product details
- Check whether products are expired or how long until they expire
- Sort product list automatically by ID
- Use option `0` in menu to return or redisplay

---

## Final Project Specification

- Structured product input is handled in the `add_product()` function
- Terminal menu and user input navigation is controlled in `main()` and `display()`
- Deletion of product by ID is implemented in `remove_product()`, with binary search logic
- Editing of product details is managed by `edit_product()`, including automatic removal if stock becomes 0
- Expiration status is checked using `check_expired_date()`
- Sorting by ID is done in `sortProduct()`
- Date validation is enforced in `inputExpiredDate()`

---

## Code Overview

- `main()` – Controls overall program flow and user input loop
- `display()` – Prints the main menu interface
- `add_product()` – Adds a new product with validation and struct
- `display_product()` – Displays all product data or notifies if empty
- `remove_product()` – Removes a product by ID using binary search
- `edit_product()` – Edits product fields and validates stock removal
- `check_expired_date()` – Compares current time with expiration date and prints status
- `inputExpiredDate()` – Parses user input to ensure proper date format
- `sortProduct()` – Sorts the list of products in ascending order by ID

---
