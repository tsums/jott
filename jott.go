/*
    jott - a terminal note taking program
    Written in Golang because reasons.
    Uses HouzuoGuo's tiedot NoSQL Database for storage.
    Trevor Summerfield
    http://trevorsummerfield.com
*/

package main

import (
    "encoding/json"
    "fmt"
    "os"
    "strconv"
    "time"
    "github.com/HouzuoGuo/tiedot/db"
)

// constants and vars and stuff

var dbHome string
var jDB db.DB


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

    jDB , err := db.OpenDB(dbHome)
    if err != nil {
        panic(err)
    }

    jotts := jDB.Use("jotts")

    docID, err := jotts.Insert(map[string]interface{}{"timestamp": int32(time.Now().Unix()),"text": "golang.org"})
    if err != nil {
        panic(err)
    }
    fmt.Println(docID)

    var query interface{}
    json.Unmarshal([]byte(`[{"has": ["timestamp"]}]`), &query)

    queryResult := make(map[int]struct{}) // query result (document IDs) goes into map keys

    if err := db.EvalQuery(query, jotts, &queryResult); err != nil {
        panic(err)
    }

    for id := range queryResult {
		readBack, err := jotts.Read(id)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Query returned document %v\n", readBack)
	}

}

func makeDB() {

    jDB , err := db.OpenDB(dbHome)

    if (err != nil) {
        panic(err)
    }
    if err := jDB.Create("jotts"); err != nil {
        panic(err)
    }

    jotts := jDB.Use("jotts")
    if err := jotts.Index([]string{"timestamp"}); err != nil {
        panic(err)
    }
}

func main() {

    dbHome = "/tmp/jott"
    // check to see if there is a db here, make one if not
    if _, err := os.Stat(dbHome); err != nil {
        fmt.Println("init: Database not Found.\ninit: Making .jott file at current wrkdir")
        makeDB()
        os.Exit(0)
    }
    // if there was, open it

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
            os.RemoveAll(dbHome)
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
