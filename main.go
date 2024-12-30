package main

import "fmt"

type Product struct {
	id    int64
	title string
	price float64
}

func main() {
	// Working with arrays
	// Empty array => hobbies := []string{}
	hobbies := []string{"Music", "Coding", "Gaming"}
	fmt.Println(hobbies)
	// Playing with slices
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])
	fmt.Println(cap(hobbies))
	hobbies = append(hobbies, "Upskilling")
	fmt.Println(hobbies)

	// Working with struct arrays
	products := []Product{
		{
			id:    1,
			title: "Product 1",
			price: 100.00,
		},
		{
			id:    2,
			title: "Product 2",
			price: 200.00,
		},
	}
	thirdProduct := Product{
		id:    3,
		title: "Product 3",
		price: 300.00,
	}
	products = append(products, thirdProduct)
	fmt.Println(products)

	// Working with maps
	// Empty map => person := map[string]string{}
	person := map[string]string{
		"first_name": "John",
		"last_name":  "Doe",
		"email_id":   "john_doe@gmail.com",
	}
	fmt.Println(person)
}
