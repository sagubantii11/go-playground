package main

import "fmt"

type Product struct {
	id    int64
	title string
	price float64
}

func main() {
	// Working with arrays
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
	products = append(products, Product{
		id:    3,
		title: "Product 3",
		price: 300.00,
	})
	fmt.Println(products)
}
