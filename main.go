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

  // Query the database
  rows, err := db.Query(`
       SELECT 
	     sepal_length as sLength,
	     sepal_width  as sWidth,
	     petal_length as pLength,
	     petal_width  as pWidth,
             species   
       FROM iris
       WHERE species = $1`, "setosa")
   u.ErrNil(err, "Unable to select from iris")
   defer rows.Close()

   // Iterate over the rows, sending the sults to 
   // standard out
   for rows.Next() {
      var (
             sLength float64
             sWidth  float64
             pLength float64
             pWidth  float64
             species string
        )
        if err := rows.Scan(&sLength, &sWidth, &pLength, &pWidth, &species) ; err != nil {
		log.Fatal(err)
        }
        fmt.Printf("%.2f, %.2f, %.2f, %.2f, %s\n", sLength, sWidth, pLength, pWidth, species)
    }
     // check for errors after we are done iterating over rows.
     if err := rows.Err(); err != nil {
       log.Fatal(err)
     }

     // Update some values
     res, err := db.Exec("UPDATE iris SET species = 'setosa' WHERE species = 'Iris-setosa'")
     u.ErrNil(err, "Unable to update iris")

     // See how many rows are affected
     rowCount, err := res.RowsAffected()
     u.ErrNil(err, "Error checking rows were affected")
     // Output the number of rows affected
     fmt.Printf("Affected rows:  %d\n", rowCount)

// sql.Open() does not establish any connection to the database
// It just prepares the database connection value
// for later use. To make sure the database is available and 
// accessible, we will use db.Ping()
//  if err := db.Ping(); err != nil {
//	panic(err)
//  }
//  fmt.Println("Successfully Connected to ", pgURL)

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
