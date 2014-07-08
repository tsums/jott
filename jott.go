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
    "bufio"
    "time"
    "strings"
    "github.com/HouzuoGuo/tiedot/db"
)

// constants and vars and stuff

var dbHome string
var jDB db.DB


func msToTime(ms string) (time.Time, error) {
    msInt, err := strconv.ParseInt(ms, 10, 64)
    if err != nil {
        return time.Time{}, err
    }

    return time.Unix(0, msInt*int64(time.Millisecond)), nil
}

// this function gets called whenever we can't parse user syntax
func syntax() {
    fmt.Println("jott - quick notes for your terminal - v0.2 \"itworks\"")
}

func params() {
    fmt.Print("n | new\t\tadd a new jott to this user's store\n" + "ls | list\tlist jotts in this user's store\n" + "purge\t\tremove all jotts from this user's store.\n")
}
// creates new jotts and adds them to the store
func new(text []string) {
    if (len(text) <= 0) {
        fmt.Println("jott:\t enter new jott below, [CTRL+D] to save.")
        scanner := bufio.NewScanner(os.Stdin)
        var line string
        for scanner.Scan() {
            line = scanner.Text()
            text = append(text, line)
        }
    }

    jDB , err := db.OpenDB(dbHome)
    if err != nil {
        panic(err)
    }

    fText := strings.Join(text, " ")
    if len(fText) == 0 {
        fmt.Println("jott:\tNothing entered, quitting.")
        os.Exit(4)
    }

    jotts := jDB.Use("jotts")

    jotts.Insert(map[string]interface{}{"timestamp": int(time.Now().Unix()),"text": fText})
    fmt.Println("jott:\t jott stored.")

}

func list(num int) {
    jDB , err := db.OpenDB(dbHome)
    if err != nil {
        panic(err)
    }
    jotts := jDB.Use("jotts")
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
        fmt.Print(readBack["timestamp"])
        fmt.Printf("\t")
		fmt.Println(readBack["text"])
	}
    if len(queryResult) == 0 {
        fmt.Println("jott:\tno jotts found.")
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

    dbHome = os.Getenv("HOME") + "/.jott"
    // check to see if there is a db here, make one if not
    if _, err := os.Stat(dbHome); err != nil {
        fmt.Println("init: Database not Found.\ninit: Making .jott file at: " + dbHome)
        makeDB()
        os.Exit(0)
    }

    args := os.Args[1:]

    // if no args, throw syntax message
    if len(args) == 0 {
        syntax()
        fmt.Println("run 'jott help' for usage information.")
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
    } else if args[0] == "help" || args[0] == "h" {
        syntax()
        params()
        os.Exit(0)
    } else if (args[0] == "ls" || args[0] == "list") {
        if len(args) == 2 {
            num, _ := strconv.Atoi(args[1])
            list(num)
        } else {
            list(5)
        }
    } else if args[0] == "n" || args[0] == "new" {
        if (len(args) > 1) {
            new(args[1:len(args)])
        } else {
            new([]string{})
        }
    } else {
        fmt.Println("jott:\tcommand not recognized: run 'jott help' for syntax.")
    }
}
