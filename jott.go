/*
    jott - a terminal note taking program
    Trevor Summerfield
    http://trevorsummerfield.com

*/

package main

import "fmt"
import "os"

func syntax() {
    fmt.Println("This is the syntax error message!")
}

func main() {

    args := os.Args[1:]

    if len(args) == 0 {
        syntax()
        return
    }



    if args[0] == "-n" || args[0] == "new"{
        fmt.Println("Making a new jott!")
    } else {
        syntax()
    }
}
