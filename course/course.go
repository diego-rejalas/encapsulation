package course

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIds []uint
	Classes map[uint]string
}

func (c *Course) ChangePrice(price float64) {
	c.Price = price
}

// Receptor de clases
func (c *Course) PrintClasses() {
	text := "Las clases son: "

	for _, class := range c.Classes {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}
