package kata

import (
	"fmt"
	"strings"
)

func main() {
	var name string

	name = "Sam Harris"
	fmt.Println(AbbrevName(name))

}

func AbbrevName(name string) string {

	arr := strings.Split(name, " ")
	a1 := arr[0] //Sam
	a2 := arr[1] //Harris

	a3 := string(a1[0])
	a4 := string(a2[0])

	a5 := a3 + "." + a4

	return a5
}
