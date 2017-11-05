package zoo

import (
	"fmt"
	"os"
	"strconv"
	"animals"
)

func usage()  {
	fmt.Println("usage: zoo <animal> [count]")
}

func main() {
	args := os.Args
	if len(args) < 3 {
		usage()
	}
	switch args[1] {
	case "cat":
		count, err1 := strconv.Atoi(args[2])
		if err1 != nil {
			fmt.Println("error:", err1)
		}
		sound := animals.CatCall(count)
		fmt.Println(sound)
	}

}
