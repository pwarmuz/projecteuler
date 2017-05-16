/*
   @Project:            projecteuler
   @Date:               2017-05-06 02:20:29
   +Last Modified time: 2017-05-16 03:25:30
*/
package main

import "fmt"

func main() {
	t, p := Problem0001(1000)
	fmt.Println("Math theory:", t, "Programmatic:", p)

	fmt.Println("Total:", Problem0002())
	Problem0003()
}
