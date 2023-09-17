package main

import (
	"fmt"
	"os"
)

func usage() {
    fmt.Println("./pass-vault [ARGS]")
    fmt.Println("ARGS: ")
    fmt.Println("-f [filename]")
    fmt.Println("-a [add, remove, list]")
    fmt.Println("----------------------")
    fmt.Println("if add: ");
    fmt.Println("    -u [username]")
    fmt.Println("    -p [password]")
    fmt.Println("endif add")
    fmt.Println("----------------------")
    fmt.Println("if remove: ")
    fmt.Println("    -u [username to remove]")
    fmt.Println("----------------------")
}

func main() {
    fmt.Println("pass_vault")
    if len(os.Args) < 5 {
        usage()
    }

}
