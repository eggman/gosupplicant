package main

import (
    "fmt"
    wifi "github.com/mdlayher/wifi"
)

func main() {

    var c, _ = wifi.New()
    defer c.Close()

    ifis, _ := c.Interfaces()
    for _, ifi := range ifis {
        fmt.Printf("%s\n", ifi.Name, )
    }

    return
}


