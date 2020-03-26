// sw audio devices project main.go
package main

import (
	"bufio"
	"fmt"

	//"io/ioutil"
	"log"
	"os"
	"os/exec"

	//"strings"

	//"strconv"
	"time"
)

func gethomedir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func getnames() {
	var devices string = gethomedir() + "\\audiodevices.txt"
	deviceList, err := os.Open(devices)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(deviceList)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	deviceList.Close()

	arrLen := len(txtlines)

	for i := 1; i <= arrLen-1; i++ {
		fmt.Printf(txtlines[i])
	}

}

func firstTimeRun() {

	var psScript string = gethomedir() + "\\getdevices.ps1"
	getList := exec.Command("powershell.exe", "-noexit", "-file", psScript)
	getList.Start()

	fmt.Println("Just a moment I'm getting the list of your devices for you....")
	time.Sleep(5 * time.Second)
	getnames()

}

func switcher() {
	headset := exec.Command("nircmd.exe", "setdefaultsounddevice", "Commonly Used Headset")
	boxes := exec.Command("nircmd.exe", "setdefaultsounddevice", "Commonly Used Boxes")

	if _, err := os.Stat("1.check"); err == nil {
		boxes.Start()
		os.Rename("1.check", "0.check")
	} else if _, err := os.Stat("0.check"); err == nil {
		headset.Start()
		os.Rename("0.check", "1.check")
	} else {
		os.Create("0.check")
	}
}

func main() {
	firstTimeRun()
}
