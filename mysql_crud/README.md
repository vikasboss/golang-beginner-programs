# MySQL CRUD

This is an example project that demonstrates the basics of working with the MySQL database in Golang. Go provides an abstract data layer using the `database/sql` standard library. To work with a specific database, you need to load the corresponding driver, and then working with most databases is relatively similar.

## Prerequisites

Before you proceed, ensure you have the following prerequisites in place:

1. Local MySQL instance - Make sure you have a running MySQL server accessible from your Golang application.
2. Valid Golang installation - Install Golang on your system and set up the environment variables correctly.

## Implementation

In this project, we will cover basic CRUD (Create, Read, Update, Delete) operations using MySQL with Golang.

### 1. Setting up the Database Connection

To start, you need to connect to your MySQL database. Import the necessary packages and create a connection to the database using the `database/sql` package and the MySQL driver.

```go
// Import necessary packages
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

// Create a global database connection
var db *sql.DB

// Connect to the MySQL database
func ConnectDB() {
    // Replace "user", "password", "host", and "database_name" with your MySQL credentials
    connectionString := "user:password@tcp(host)/database_name"

    // Open a connection to the database
    var err error
    db, err = sql.Open("mysql", connectionString)
    if err != nil {
        panic(err)
    }

    // Check if the connection is successful
    if err = db.Ping(); err != nil {
        panic(err)
    }

    fmt.Println("Connected to the MySQL database!")
}
