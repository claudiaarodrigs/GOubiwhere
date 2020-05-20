package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	variables := [3]string{"var1", "var2", "var4"}
	var n int = 3
	var data int
	var sum int
	var j int = 0

	//DATAbase file
	database, _ := sql.Open("sqlite3", "./DATAbase.db")
	statement2, _ := database.Prepare("CREATE TABLE IF NOT EXISTS ExtDevice (ID INTEGER PRIMARY KEY, var1 INTEGER, var2 INTEGER, var3 INTEGER, var4 INTEGER)")
	statement2.Exec()

	biarray := make([][]int, n) // Make the outer slice and give it size n
	for i := 0; i < n; i++ {

		rows, _ := database.Query("SELECT " + variables[i] + " FROM ExtDevice ORDER BY ID DESC LIMIT " + strconv.Itoa(n) + "")
		biarray[i] = make([]int, len(variables)) // Make one inner slice per iteration and give it size

		for rows.Next() {
			rows.Scan(&data)
			biarray[i][j] = data
			j++
		}
		j = 0
	}

	//Print
	for i := 0; i < len(variables); i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", biarray[j][i])
		}
		fmt.Println("")
	}

	//average
	for i := 0; i < len(variables); i++ {
		sum = 0
		for j := 0; j < n; j++ {

			sum += biarray[i][j]
		}
		fmt.Printf("%d ", sum/n)
		fmt.Println("")
	}

}
