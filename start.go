package main

import (
	"fmt"
	"github.com/jefersonm/tech-linter/validation/lints"
)

func main() {
	fmt.Printf("Starting Tech Linter\n")
	lints.HasSpecificDependency("JUnit")
}