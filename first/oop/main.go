package main

import "fmt"

func main() {
	cleopas := book{
		title:  "Adventures of Cleopas",
		author: "Vivian Gil",
		price:  15,
	}
	grief := book{
		title:  "The Grief Child",
		author: "Lawrence Darmani",
		price:  15,
	}
	oliver := book{
		title:  "Oliver Twist",
		author: "Charles Dickens",
		price:  15,
	}

	snake := game{
		title: "Nokia Snake",
		price: 20,
	}

	var products list

	products = append(products, &cleopas, &grief, &oliver, &snake)

	fmt.Printf("%p\n\n", &products)

	products.print()
}
