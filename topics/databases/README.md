# Databases

Package [database/sql](https://golang.org/pkg/database/sql/) is the entry point
for working with databases.

The standard library package is an abstract, the adapters for database engines are separated. Here is a list of available drivers:

* [wiki/SQLDrivers](https://github.com/golang/go/wiki/SQLDrivers)

>  The sql.DB abstraction is designed to keep you from worrying about how to
>  manage concurrent access to the underlying datastore.

## A complete example

We first create a user table with the following schema:

```sql
    CREATE TABLE `userinfo` (
        `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
        `username` VARCHAR(64) NULL,
        `departname` VARCHAR(64) NULL,
        `created` DATE NULL
    );
```

Insertion and query example.

```go
    package main

    import (
        "database/sql"
        "fmt"
        "time"
        _ "github.com/mattn/go-sqlite3"
    )

    func main() {
        db, err := sql.Open("sqlite3", "./foo.db")
        checkErr(err)

        // insert
        stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
        checkErr(err)

        res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
        checkErr(err)

        id, err := res.LastInsertId()
        checkErr(err)

        fmt.Println(id)
        // update
        stmt, err = db.Prepare("update userinfo set username=? where uid=?")
        checkErr(err)

        res, err = stmt.Exec("astaxieupdate", id)
        checkErr(err)

        affect, err := res.RowsAffected()
        checkErr(err)

        fmt.Println(affect)

        // query
        rows, err := db.Query("SELECT * FROM userinfo")
        checkErr(err)
        var uid int
        var username string
        var department string
        var created time.Time

        for rows.Next() {
            err = rows.Scan(&uid, &username, &department, &created)
            checkErr(err)
            fmt.Println(uid)
            fmt.Println(username)
            fmt.Println(department)
            fmt.Println(created)
        }

        rows.Close() //good habit to close

        // delete
        stmt, err = db.Prepare("delete from userinfo where uid=?")
        checkErr(err)

        res, err = stmt.Exec(id)
        checkErr(err)

        affect, err = res.RowsAffected()
        checkErr(err)

        fmt.Println(affect)

        db.Close()

    }

    func checkErr(err error) {
        if err != nil {
            log.Fatal(err)
        }
    }
```