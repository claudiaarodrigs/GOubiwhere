//****** DATAbase.go ******//
// V1 - development version
// Created by: Cláudia Rodrigues claudiagr.rodrigues@gmail.com
// With: Visual Studio Code
// Started at: 17-05-2020
// Finished at: tbd
//
// This file creates a data base of periodic reading values of CPU combined percentage and RAM free memory, in kB,
// with a second period lecture. Also, it simulates an external device, generating periodically (with period of 1s)
// random data samples
// DATABASE organization: Table OScollector for CPU and RAM values;
//						  Table ExtDevice is for samples of the external device

package main

import (
	"database/sql"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type sample struct {
	var1 int
	var2 int
	var3 int
	var4 int
}

func main() {

	var samp sample

	//DATAbase file
	database, _ := sql.Open("sqlite3", "./DATAbase.db")

	// DATA BASE table for CPU AND RAM values
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS OScollector (ID INTEGER PRIMARY KEY, FreeMemory INTEGER, CPUcombined INTEGER)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO OScollector (FreeMemory, CPUcombined) VALUES (?,?)")

	//EXTERNAL device simulator
	statement2, _ := database.Prepare("CREATE TABLE IF NOT EXISTS ExtDevice (ID INTEGER PRIMARY KEY, var1 INTEGER, var2 INTEGER, var3 INTEGER, var4 INTEGER)")
	statement2.Exec()
	statement2, _ = database.Prepare("INSERT INTO ExtDevice (var1, var2, var3, var4) VALUES (?,?,?,?)")

	for {
		//colecting OS info
		m, _ := mem.VirtualMemory()
		percent, _ := cpu.Percent(time.Second, false)
		statement.Exec(m.Available/1024, percent[0])

		//colecting sample and save it to DATAbase
		samp = randomGenerator()
		statement2.Exec(samp.var1, samp.var2, samp.var3, samp.var4)

		time.Sleep(time.Second)
	}
}

// Random generator of 1 sample : Simulates the data from the external device
// return: structure of type sample
func randomGenerator() sample {

	var out sample

	rand.Seed(time.Now().UnixNano())

	out.var1 = rand.Intn(1000)
	out.var2 = rand.Intn(1000)
	out.var3 = rand.Intn(1000)
	out.var4 = rand.Intn(1000)

	return out
}
