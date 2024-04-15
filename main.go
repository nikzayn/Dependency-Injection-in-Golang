package main

import "fmt"

// Interface for collection of methods, maninly to handle behaviour
type Category interface {
	checkPrices()
}

// Composition
type Collection struct {
	braSize  string
	color    string
	price    int
	category Category
}

// Injection occurs here
func newCollection(category Category) *Collection {
	return &Collection{
		category: category,
	}
}

// Behavior
func (c *Collection) doCheck() {
	c.price++
	if c.price == 10 {
		c.category.checkPrices()
	}
}

// We achive reusability here
type CheckPantiesCollection struct{}
type CheckLingerieCollection struct{}

func (pc *CheckPantiesCollection) checkPrices() {
	fmt.Println("Here's the details of Panties Collection....")
}

func (lc *CheckLingerieCollection) checkPrices() {
	fmt.Println("Here's the details of Lingerie Collection....")
}

func main() {
	pc := newCollection(&CheckPantiesCollection{})
	lc := newCollection(&CheckLingerieCollection{})
	for i := 0; i < 11; i++ {
		pc.doCheck()
		lc.doCheck()
	}
}
