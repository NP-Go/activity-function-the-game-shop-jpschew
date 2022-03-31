package main

import "fmt"

//declare List
type list []game

//declare method for List Print
func (l list) printContent() {
	for _, item := range l {
		fmt.Println(item)
	}
}
