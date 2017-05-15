package main

import (
	"fmt"
	"github.com/JokiNS/First-time/06_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
