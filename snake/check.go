package snake

import (
	"fmt"
)

//Checker must be implemented by any check
type Checker interface {
	Run() error
	Stop() error
	Result() (Check, error)
}

//Check must be type for check result outputs
type Check struct {
	Name      string
	Timestamp int
	Result    string
}

//Printresult prints result
func (c *Check) Printresult() error {
	_, err := fmt.Println("Checkname: ", c.Name, " result: ", c.Result)
	return err
}

//NewCheck is construcor
func NewCheck() (c *Check) {
	c = new(Check)
	return c
}
