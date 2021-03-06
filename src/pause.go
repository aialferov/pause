package main

import (
    "fmt"
    "os"
    "os/signal"
    "time"
    "syscall"
    "strconv"
)

var version string
var git_sha string

const Usage = "Usage: pause [version] [--number N]"

const VersionFormat = "Version %s (git-%s)\n"
const TsFormat = "2006-01-02T15:04:05.999Z"

func main() {
    if len(os.Args) > 1 && os.Args[1] == "version" {
        fmt.Printf(VersionFormat, version, git_sha)
        os.Exit(0)
    }

    n := 1
    var err error
    if len(os.Args) > 2 && os.Args[1] == "--number" {
        n, err = strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Println(Usage)
            os.Exit(2)
        }
    }

    c := make(chan os.Signal, n)
    signal.Notify(c, os.Interrupt,
                     syscall.SIGTERM, syscall.SIGINT, syscall.SIGCHLD)

    for i := n; i > 0; i-- {
        fmt.Printf("I%s: Waiting for %d signal(s)...\n",
                   time.Now().UTC().Format(TsFormat), i)
        <-c
    }

    fmt.Printf("I%s: Done.\n", time.Now().UTC().Format(TsFormat))
}
