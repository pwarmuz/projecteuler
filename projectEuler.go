/*
   @Project:            projecteuler
   @Author:             Phil
   @Date:               2017-05-06 02:20:29
   +Last Modified by:   Phil
   +Last Modified time: 2017-05-07 20:35:45
*/
package main

import "fmt"

func main() {
	t, p := problem0001(1000)
	fmt.Println("Math theory:", t, "Programmatic:", p)

	fmt.Println("Total:", problem0002())
	problem0003()
}
