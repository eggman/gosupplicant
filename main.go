package main

import (
    nl80211 "github.com/mdlayher/wifi"
)

func main() {
    var c , err = nl80211.New()

    if err != nil {
        // print error
        return 
    }
    c.Close()
    return
}


