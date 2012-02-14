// Command line tool to communicate with partiscid
package main

import (
    "fmt"
    "flag"
    "strings"
)

var server *string = flag.String("server", "localhost", "partiscid address (IP or DNS name)")
var port *int = flag.Int("port", 7777, "partiscid port")

var commands map[string]func([]string) int

func init() {
    commands = make(map[string]func([]string) int)
    commands["UPDATE"] = cmdUpdate
}

func cmdUpdate(args []string) int {
    fmt.Println("cmdUpdate", args)
    return 0
}

func main() {
    flag.Parse()
    fmt.Println("server:", *server)
    fmt.Println("port:", *port)
    if flag.NArg() > 0 {
        cmd := flag.Args()[0]
        cmd = strings.ToUpper(cmd)
        fmt.Println(cmd)
        if f, ok := commands[cmd]; ok {
            f(flag.Args())
        } else {
            fmt.Println("command not recognized: ", cmd)
        }
    }

}


