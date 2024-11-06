package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
    height := 30
    width := 200
    var init bool = true
    for {
        flush()
        grid := initGrid(width,height,init) 
        init = false
        for i:=0;i<height;i++{
            fmt.Printf("\n")
            for j:=0;j<width;j++{
                fmt.Printf("%s",grid.cells[i][j])   
            }
        }
    fmt.Printf("\n")
    time.Sleep(1*time.Second)
    }
}

func flush() {
    cmd := exec.Command("clear")   
    cmd.Stdout = os.Stdout
    cmd.Run()
}



