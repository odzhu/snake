package snake

import (
	"fmt"
	"net/http"
	"time"
)

//Whereami must be type for check result outputs
type Whereami struct {
	Name      string
	Timestamp time.Time
	Result    string
}

//Run runs the check
func (c *Whereami) Run() {
	c.Name = "Whereami"
	c.Timestamp = time.Now()
}

//Resulthttp handles http
func (c *Whereami) Resulthttp(w http.ResponseWriter, r *http.Request) {
	c.Run()
	fmt.Fprintf(w, "whereami %s:", c)
}

//NewWhereami is construcor
func NewWhereami() (c *Whereami) {
	c = new(Whereami)
	return c
}
