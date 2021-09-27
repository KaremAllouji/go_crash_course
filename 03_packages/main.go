package main

import (
	"fmt"
	"math"

	"github.com/KaremAllouji/go_crash_course/03_packages/strutil"
)

// global var are also allowed

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(strutil.Reverse("Hello"))
}

// go run <name.ext>
// go build <name.ext>
// to create bin :
// go mod init <name.ext>
// go install
