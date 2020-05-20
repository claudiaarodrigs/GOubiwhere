//****** Interface.go ******//
// V1 - development version
// Created by: Cláudia Rodrigues claudiagr.rodrigues@gmail.com
// With: Visual Studio Code
// Started at: 19-05-2020
// Finished at: 20-05-2020
//
// This file creates an app that implements a command line interface between a data base file (.db) and the user.
// To do that, the binary files of both DataCollector.exe and Interface.exe must be in the same directory, and
//DataCollector must be executed first to initialize the .db file.

package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	//Evaluate max and min number of args
	args := os.Args
	if len(args) < 2 || len(args) > 7 {
		fmt.Println("Wrong number of arguments!")
		fmt.Println("See 'Interface.exe -h' for usage")
		os.Exit(1)
	}

	argsWithoutProg := os.Args[1:]

	switch argsWithoutProg[0] {
	//////////////////////////////////////////////////////////////////////////////////
	//case return metrics of all variables
	case "-ma":
		if len(argsWithoutProg) > 2 {
			fmt.Println("Too much arguments!")
			fmt.Println("See 'Interface.exe -h' for usage")
			os.Exit(3)
		}
		// second argument must be the number of samples to be returned
		n, err := strconv.Atoi(argsWithoutProg[1])
		if err != nil || n < 0 {
			fmt.Printf("%q is not a positive number.\n", argsWithoutProg[1])
			fmt.Println("See 'Interface.exe -h' for usage")
			os.Exit(4)
		}
		variables := []string{"", "var1", "var2", "var3", "var4"}
		printCmd(variables, n, false)

	//////////////////////////////////////////////////////////////////////////////////
	//case return n metrics of x variables
	case "-m":
		if len(argsWithoutProg) < 3 {
			fmt.Println("Too few arguments!")
			fmt.Println("See 'Interface.exe -h' for usage")
			os.Exit(5)
		}
		// second argument must be the number of samples to be returned
		n, err := strconv.Atoi(argsWithoutProg[1])
		if err != nil || n < 0 {
			fmt.Printf("%q is not a positive number.\n", argsWithoutProg[1])
			fmt.Println("See 'Interface.exe -h' for usage")
			os.Exit(6)
		}
		// next arguments must be variables
		for i := 2; i < len(argsWithoutProg); i++ {
			if argsWithoutProg[i] != "var1" && argsWithoutProg[i] != "var2" && argsWithoutProg[i] != "var3" && argsWithoutProg[i] != "var4" {
				fmt.Printf("%q is not a valid argument.\n", argsWithoutProg[i])
				fmt.Println("See 'Interface.exe -h' for usage")
				os.Exit(7)
			}
		}
		variables := make([]string, len(argsWithoutProg)-1)
		variables[0] = "id"
		j := 1
		for i := 2; i < len(argsWithoutProg); i++ {
			variables[j] = argsWithoutProg[i]
			j++
		}
		printCmd(variables, n, false)

	//////////////////////////////////////////////////////////////////////////////////
	//case return average of x variables
	case "-a":
		// next arguments must be variables
		for i := 1; i < len(argsWithoutProg); i++ {
			if argsWithoutProg[i] != "var1" && argsWithoutProg[i] != "var2" && argsWithoutProg[i] != "var3" && argsWithoutProg[i] != "var4" {
				fmt.Printf("%q is not a valid argument.\n", argsWithoutProg[i])
				fmt.Println("See 'Interface.exe -h' for usage")
				os.Exit(8)
			}
		}

		variables := make([]string, len(argsWithoutProg))
		variables[0] = "id"
		j := 1
		for i := 1; i < len(argsWithoutProg); i++ {
			variables[j] = argsWithoutProg[i]
			j++
		}
		n := 5
		printCmd(variables, n, true)

	//////////////////////////////////////////////////////////////////////////////////
	//case help
	case "-h":
		fmt.Println("")
		fmt.Println("-------------------------------------------------------------------------------------------------------------")
		fmt.Println("This app interacts with a data base where external device's samples are stored.")
		fmt.Println("Each sample is composed by 4 variables: var1, var2, var3 and var4")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Flags:")
		fmt.Println(" -ma, 'metrics all': returns the last metrics for all the variables for n samples")
		fmt.Println(" -m,  'metrics':     returns the last metrics for n samples of one or more variables")
		fmt.Println(" -a,  'average':     returns the average for all the samples of one or more variables")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("The Flags must be used precedenting the number of samples (n) and, next, the variables you want to evaluate.")
		fmt.Println("")
		fmt.Println("Example (1) [get metrics for all variables of 5 samples]:")
		fmt.Println(">>Interference.exe -ma 5")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Example (2) [get metrics for var1 and var3 of 8 samples]:")
		fmt.Println(">>Interference.exe -m 8 var1 var3")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Example (2) [get average for var1 and var2]:")
		fmt.Println(">>Interference.exe -a var1 var2")
		fmt.Println("-------------------------------------------------------------------------------------------------------------")
		os.Exit(9)

	//////////////////////////////////////////////////////////////////////////////////
	default:
		fmt.Println("Wrong first argument")
		fmt.Println("See 'Interface.exe -h' for usage")
		os.Exit(2)
	}
}

func printCmd(variables []string, n int, average bool) {

	var data int
	var sum int
	var j int = 0
	var count int = 0

	//DATAbase file connection
	database, _ := sql.Open("sqlite3", "./DATAbase.db")
	statement2, _ := database.Prepare("CREATE TABLE IF NOT EXISTS ExtDevice (ID INTEGER PRIMARY KEY, var1 INTEGER, var2 INTEGER, var3 INTEGER, var4 INTEGER)")
	statement2.Exec()

	if !average { //Not an average command
		biarray := make([][]int, len(variables))
		//Get n last rows from database and save on a 2D array
		for i := 0; i < len(variables); i++ {
			rows, err := database.Query("SELECT " + variables[i] + " FROM ExtDevice ORDER BY ID DESC LIMIT " + strconv.Itoa(n) + "")
			//check error
			if err != nil {
				fmt.Println("Failed to run query", err)
				return
			}

			biarray[i] = make([]int, n)

			for rows.Next() {
				rows.Scan(&data)
				biarray[i][j] = data
				j++
			}
			j = 0
		}
		// Print data
		for i := 0; i < len(variables); i++ {
			fmt.Printf(variables[i] + ": ")
			for j := 0; j < n; j++ {
				fmt.Printf("%d ", biarray[i][j])
			}
			fmt.Println("")
		}

	} else { // Average command
		//Get complete column and calculate average
		for i := 1; i < len(variables); i++ {
			rows, err := database.Query("SELECT " + variables[i] + " FROM ExtDevice")
			//check error
			if err != nil {
				fmt.Println("Failed to run query", err)
				return
			}

			for rows.Next() {
				rows.Scan(&data)
				sum += data
				count++
			}
			fmt.Println("Average "+variables[i]+":", sum/count)
			count = 0
			sum = 0
		}

	}
}
