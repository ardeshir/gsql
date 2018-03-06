package main

import (
    "fmt"
    "runtime"
    "time"
    "database/sql"
    "os"
    "log"

   u "github.com/ardeshir/version"
    // pg is the library that allows us to connect
    // to postgres with database/sql
    _ "github.com/lib/pq"
)


var (
 debug bool = false
 version string = "0.0.1"
 )

func printEnv() {

    fmt.Print("We're using ", runtime.Compiler, " ")
    fmt.Println("On a", runtime.GOARCH, "machine")
    fmt.Println("With Go version", runtime.Version())
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
//+++++++++  main
// Get the postgres connection URL
// from an envirment variable
 pgURL := os.Getenv("PGURL")
 if pgURL == "" {
   log.Fatal("PGURL is empty")
 }

// Open a database valie. Specifiy the postgres driver
// for databases/sql
  db, err := sql.Open("postgres", pgURL)
  u.ErrNil(err, "Unable to open postgres connection")

  defer db.Close()

// sql.Open() does not establish any connection to the database
// It just prepares the database connection value
// for later use. To make sure the database is available and 
// accessible, we will use db.Ping()
  if err := db.Ping(); err != nil {
	log.Fatal(err)
 }



//++++++++++  footer 
  if debugTrue() {

       printEnv()

       checkGC()

    u.V(version)
  }

}

// Function to check env variable DEFAULT_DEBUG bool is set
func debugTrue() bool {

     if os.Getenv("DEFAULT_DEBUG") != "" {
        return true
     }
     return false 
}
