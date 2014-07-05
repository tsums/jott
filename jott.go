/*
    jott - a terminal note taking program
    Written in Golang because reasons.
    Trevor Summerfield
    http://trevorsummerfield.com
*/

package main

import (
    "fmt"
    "os"
    "strconv"
    "github.com/HouzuoGuo/tiedot/db"
)

// constants and vars and stuff

var dbHome string



// this function gets called whenever we can't parse user syntax
func syntax() {
    fmt.Println("This is the syntax error message!")
}
// creates new jotts and adds them to the store
func new(text []string) {
    if (len(text) > 0) {
        fmt.Println("Content was already given at the terminal.")
    } else {
        fmt.Println("Gonna ask for the content now.")
    }
}

func list(num int) {
    fmt.Println("you requested this many jotts: " + strconv.Itoa(num))
}

func makeDB() {

    jDB , err := db.OpenDB(dbHome)

    if (err != nil) {
        panic(err)
    }
    if err := jDB.Create("jotts"); err != nil {
        panic(err)
    }
}

func main() {

    dbHome = ".jott"
    // check to see if there is a db here, make one if not
    if _, err := os.Stat(dbHome); err != nil {
        fmt.Println("Making .jott file at current wrkdir")
        makeDB()
        os.Exit(0)
    // if there was, open it
    } else {
        jDB , err := db.OpenDB(dbHome)
        if err != nil {
            panic(err)
        }
        for _, name := range jDB.AllCols() {
            fmt.Printf("I have a collection called %s\n", name)
        }
    }

    args := os.Args[1:]

    // if no args, throw syntax message
    if len(args) == 0 {
        syntax()
        os.Exit(1)
    }

    if args[0] == "purge" {
        var resp string
        fmt.Printf("Purge all jott information? [y/N]: ")
        fmt.Scanln(&resp)

        if resp == "Y" || resp == "y" {
            os.RemoveAll(".jott")
            fmt.Println("purged jott db")
            os.Exit(0)
        } else {
            fmt.Println("not purging jotts")
            os.Exit(0)
        }
    }

    // if given new flag, check to see if we were given the jott
    if args[0] == "n" || args[0] == "new" {
        if (len(args) > 1) {
            new(args[1:len(args)])
        } else {
            new([]string{})
        }
    // if given list flag, check to see if we were given the num
    } else if (args[0] == "ls" || args[0] == "list") {
        if len(args) == 2 {
            num, _ := strconv.Atoi(args[1])
            list(num)
        } else {
            list(5)
        }
    // we were given args, but they didn't mean anything
    // throw syntax message and quit
    } else {
        syntax()
    }
}
