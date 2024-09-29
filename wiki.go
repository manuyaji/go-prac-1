package main

import (
	"fmt"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	return nil
}

func main() {
	fmt.Println("Om Ganapataye Namah!")
}
