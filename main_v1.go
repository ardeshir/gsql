package main

import (
    "fmt"
    "runtime"
    "time"
    _ "github.com/go-sql-driver/mysql"
)

var debug bool = true 

func printEnv() {
    
    fmt.Print("We're using ", runtime.Compiler, " ")
    fmt.Println("on a", runtime.GOARCH, "machine")
    fmt.Println("with Go version", runtime.Version())
    fmt.Println("Number of Goroutines: ", runtime.NumGoroutine())
    fmt.Println(" ")
}

func printStats(mem runtime.MemStats) {
    
    
    runtime.ReadMemStats(&mem)
    
    fmt.Println("mem.Alloc: ", mem.Alloc)
    fmt.Println("mem.TotalAlloc: ", mem.TotalAlloc)
    fmt.Println("mem.HeapAlloc: ", mem.HeapAlloc)
    fmt.Println("mem.NumGC: ", mem.NumGC)
    fmt.Println("------")
}

func checkGC() {
    
    var mem runtime.MemStats
    printStats(mem)

    for i := 0; i< 10; i++ {
        s := make([]byte, 100000000)
        
        if s == nil {
           fmt.Println("Operation failed")
        }
     }


     printStats(mem)

      for i := 0; i< 10; i++ {
          
         s := make([]byte, 100000000)
         if s == nil {
            fmt.Println("Operation failed")
         }
         time.Sleep(5 * time.Second)
    }
   printStats(mem)
    
}

func main() {

    if debug == true {
        
       fmt.Println("Using the MySQL Go Driver, and testing GC!")
       fmt.Println(" ")
       
       printEnv()   
       
       checkGC()
    }
    
}

