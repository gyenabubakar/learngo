package main

type Book struct {
	title  string;
	author string;
}

func main() {
	var books = [...]Book{
		{title: "Oliver Twist", author: "Charles Dickens"},
		{title: "No Sweetness Here", author: "Lawrence Darmani"},
		{title: "The Grief Child", author: "Lawrence Darmani"},
		{title: "End of the Tunnel", author: "Jane Doe"},
		{title: "Adventures of Cleopas", author: "Jane Doe"},
	}

	// fmt.Printf("Books: %#v\n\n", books)

	for _, value := range books {
		println(value.title+", by", value.author)
	}
}
