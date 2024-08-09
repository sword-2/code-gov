// main.go. Reads in file liaison.json, outputs a partial html file (*.l14)

package main

import (
	"fmt"
	"liaisons-go/liaisons"
)


func main() {
	fmt.Println("go's main function started v2\n")
	pathL14 := "htmlTable.go.l14"
	liaisons.RemoveOldL14(&pathL14) //if an old .l14 is present, remove it
	pathLiaisons := "liaisons.json"
	liaisons.ReadFile(&pathLiaisons)

	var sortCode int // variable declaration
	sortCode = 2 //1 sort on agency name, 2 sort on legal reference.
	liaisons.SortStruct(&sortCode)
	liaisons.PrintStruct()
	liaisons.WriteL14(&pathL14)
}
