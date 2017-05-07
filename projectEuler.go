/*
   @Project:            projecteuler
   @Author:             Phil
   @Date:               2017-05-06 02:20:29
   +Last Modified by:   Phil
   +Last Modified time: 2017-05-06 21:06:36
*/
package main

import "fmt"

func main() {
	t, p := problem1()
	fmt.Println("Math theory:", t, "Programmatic:", p)

	fmt.Println("Total:", problem2())
}
