package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
    Username string;
    Password string;
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

func ListUsers(users []User) {
    fmt.Println("All users: ")
    for _, user := range users {
        fmt.Printf("Username: %s, Password: %s\n", user.Username, user.Password)
    }
}

func AddUser(users *[]User, user User) {
    *users = append(*users, user)
}

func RemoveUser(users *[]User, user User) {
    for idx, usr := range *users {
        if usr.Username == user.Username {
            firstHalf := (*users)[:idx]
            secondHalf := (*users)[idx+1:]
            *users = append(firstHalf, secondHalf...)
        } 
    }
}

func Quit(outFileName string, users []User) {
    data, _ := json.Marshal(users) 
    file, _ := os.OpenFile(outFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
    defer file.Close()
    file.WriteString(string(data))

}

func main() {
    fmt.Println("pass_vault")
    if len(os.Args) < 5 {
        usage()
        os.Exit(1)
    }

    outFile := os.Args[2]
    action := os.Args[4] 
    

    var users []User
    filedata, _ := os.ReadFile(outFile)
    err := json.Unmarshal(filedata, &users)
    
    if err != nil {
        log.Fatalln(err)
    }
    if action == "list" {
        ListUsers(users)
    }  
    if action == "add" {
        user := User {os.Args[6], os.Args[8]}
        AddUser(&users, user)
        Quit(outFile, users)
    }
    if action == "remove" {
        userToRemove := User {os.Args[5], ""}
        RemoveUser(&users, userToRemove)
        fmt.Println(users) 
        Quit(outFile, users)
    }
}
