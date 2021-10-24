package main

import (
	"fmt"
	"github.com/h-u-m-a-n/les2lib"
)

func main() {
	// 1.0.0
	fmt.Printf("%c\n", les2lib.ChangeLetterCase('a'))
	fmt.Printf("%c\n", les2lib.ChangeLetterCase('B'))
	fmt.Printf("%c\n", les2lib.ChangeLetterCase('-'))

	// 1.0.1
	r1, r2, ok := les2lib.FindRoots(3, 2,1)
	showFindRootsResult(r1,r2,ok)
	r1,r2,ok = les2lib.FindRoots(3,2,-1)
	showFindRootsResult(r1,r2,ok)
	r1,r2,ok = les2lib.FindRoots(9,-12,4)
	showFindRootsResult(r1,r2,ok)
	
	// 1.0.2
	fmt.Println(les2lib.GetUUID())
}

func showFindRootsResult(r1, r2 float64, ok error) {
	if ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Println(r1, r2)
	}
}