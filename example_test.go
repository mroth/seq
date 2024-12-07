package seq_test

import (
	"fmt"

	"github.com/mroth/seq"
)

func ExampleWindow() {
	type Person struct {
		Name string
		Age  int
	}

	type People []Person

	people := People{
		{"Gopher", 13},
		{"Alice", 20},
		{"Bob", 5},
		{"Vera", 24},
		{"Zac", 15},
	}

	// Sliding window over people in []Person 2 elements at a time.
	for c := range seq.Window(people, 2) {
		fmt.Println(c)
	}

	// Output:
	// [{Gopher 13} {Alice 20}]
	// [{Alice 20} {Bob 5}]
	// [{Bob 5} {Vera 24}]
	// [{Vera 24} {Zac 15}]
}
