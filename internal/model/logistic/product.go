package logistic

import "fmt"

type Product struct {
	Id    uint64
	Title string
}

func (p *Product) String() string {
	return fmt.Sprintf("[ID: %d] %s \n", p.Id, p.Title)
}
