package main

import "fmt"

var langPrefixes = map[string]string{
	"Spanish": "Hola ",
	"French":  "Bonjour ",
	"English": "Hello ",
}

func main() {
	fmt.Println(Hello("Gyen", "Spanish"))
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return getPrefixOf(lang) + name + "!"
}

func getPrefixOf(lang string) (prefix string) {
	prefix = langPrefixes[lang]
	if prefix == "" {
		prefix = langPrefixes["English"]
	}

	return
}
