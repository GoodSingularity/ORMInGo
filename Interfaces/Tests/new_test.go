package Interfaces_test
import (
	. "github.com/stretchr/testify/assert"
	"testing"
	"Interfaces"
)
	
func TestFailure(t *testing.T) {
	type Person struct {
		Name string
		Surname string		
	}
	p := Person{"Kamil", ""}
	method := Interfaces.New(p)
	Equal(t, nil, method)
}

func TestSuccess(t *testing.T) {
	type Person struct {
		Name string
		Surname string		
	}
	p := Person{"Kamil", "M"}
	method := Interfaces.New(p)
	Equal(t, p, method)
}
