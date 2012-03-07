// Partisci is a command line utility to communicate with partiscid.
package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/briandorsey/partisci/clients/go/client"
	"github.com/briandorsey/partisci/version"
	"strconv"
	"strings"
)

// add a flag Usage string documenting commands
var server *string = flag.String("server", "localhost", "partiscid address (IP or DNS name)")
var port *int = flag.Int("port", 7777, "partiscid port")

var commands map[string]func([]string) int

const (
	cmdUsage string = `
  partisci update APP VERSION [[HOST] INSTANCE]
`
)

func init() {
	commands = make(map[string]func([]string) int)
	commands["UPDATE"] = cmdUpdate
	flag.Usage = func() {
		app := os.Args[0]
		fmt.Fprintf(os.Stderr, "Usage of %s:", app)
		fmt.Fprintf(os.Stderr, cmdUsage)
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
}

func cmdUpdate(args []string) int {
	v := new(version.Version)
	if len(args) < 2 {
		flag.Usage()
		fmt.Fprintf(os.Stderr, "UPDATE requires APP and VERSION\n")
		return 1
	}
	v.App = args[0]
	v.Ver = args[1]
	if len(args) > 2 {
		v.Host = args[2]
	}
	if len(args) > 3 {
		i, err := strconv.ParseUint(args[3], 10, 16)
		if err != nil {
			fmt.Println("Error parsing Intstance value:", args[3])
			return 1
		}
		v.Instance = uint16(i)
	}
	err := client.SendUDP(*server, *port, *v)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		return
	}
	cmd := flag.Args()[0]
	cmd = strings.ToUpper(cmd)
	if f, ok := commands[cmd]; ok {
		args := flag.Args()[1:]
		fmt.Printf("command: %s\n", cmd)
		f(args)
	} else {
		fmt.Println("command not recognized: ", cmd)
		return
	}

}
