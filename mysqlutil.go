package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.evanpurkhiser.com/netgear"
	"log"
	"os"
	"strconv"
	// s "strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InsertDeviceConnectionLog(change *netgear.ChangedDevice) {

	log.Println(time.Now())

	dbInstance := os.Getenv("MYSQL_DB_INSTANCE")
	dbPort := 3306
	database := "IoT"
	dbUser := os.Getenv("MYSQL_DB_USER")
	dbPassword := os.Getenv("MYSQL_DB_PASSWORD")

	// fmt.Printf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbInstance, strconv.Itoa(dbPort), database)

	fmt.Println("At logging - changed device: ", change.Device.IP)

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbInstance, strconv.Itoa(dbPort), database))
	checkErr(err)

	defer db.Close()

	// stmt, err := db.Prepare("INSERT INTO DeviceConnectionLog(IPAddress, MACAddress, DeviceName, DeviceSignal, ConnectionType, LinkRate, DeviceState, DeviceEventType, ChangedTimeStamp) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)")
	// stmt, err := db.Prepare("INSERT DeviceConnectionLog set IPAddress=?, MACAddress=?, DeviceName=?, DeviceSignal=?, ConnectionType=?, LinkRate=?, DeviceState=?, DeviceEventType=?, ChangedTimeStamp=?")
	// stmt, err := db.Prepare("INSERT DeviceConnectionLog SET IPAddress=?,MACAddress=?,DeviceName=?")
	// checkErr(err)
	// defer stmt.Close()

	// stmt.Exec(change.Device.IP.String, change.Device.MAC.String, change.Device.Name, change.Device.Signal, change.Device.Type, change.Device.LinkRate, change.Change, time.Now())
	var deviceState int
	if change.Change == "added" {
		deviceState = 1
	} else {
		deviceState = -1
	}

	_, err = db.Exec("INSERT INTO DeviceConnectionLog(IPAddress, MACAddress, DeviceName, DeviceSignal, ConnectionType, LinkRate, DeviceState, DeviceEventType, ChangedTimeStamp) VALUES(?,?,?,?,?,?,?,?,?)", change.Device.IP.String(), change.Device.MAC.String(), change.Device.Name, change.Device.Signal, change.Device.Type, change.Device.LinkRate, deviceState, change.Change, time.Now())
	// _, err = stmt.Exec()
	checkErr(err)

}
