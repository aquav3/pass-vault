package main

import (
	"fmt"
	"os"
)

type User struct {
    username string;
    password string;
}

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
        os.Exit(1)
    }

    action := os.Args[4] 

    if action == "list" {
        fmt.Println("well priting all the pass_vault")
    }  
    if action == "add" {
        user := User {os.Args[6], os.Args[8]}
        fmt.Println(user)    
    }
    if action == "remove" {
        user_to_remove := User {"", os.Args[6]}
        fmt.Println(user_to_remove)
    }

    

}
