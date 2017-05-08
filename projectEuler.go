/*
   @Project:            projecteuler
   @Author:             Phil
   @Date:               2017-05-06 02:20:29
   +Last Modified by:   Phil
   +Last Modified time: 2017-05-08 01:17:54
*/
package projecteuler

import "fmt"

func main() {
	t, p := Problem0001(1000)
	fmt.Println("Math theory:", t, "Programmatic:", p)

	fmt.Println("Total:", Problem0002())
	Problem0003()
}
