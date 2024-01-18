// Измените программу echo так, чтобы она выводила также os.Args[0],
// имя выполняемой программы.
// Echo2 выводит аргументы командной строки.

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
	}
	fmt.Println(s)
}
