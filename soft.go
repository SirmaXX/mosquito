package main

import (
	"os/exec"
	"fmt"
L "./lib"
) 

func main() {


cmd ,err:= exec.Command("/bin/sh", test.sh)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("output is %s\n", cmd)


}
	
