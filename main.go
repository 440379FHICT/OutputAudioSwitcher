// sw audio devices project main.go
package main

import (
	"bufio"
	//"bytes"
	"encoding/json"
	"fmt"

	//"io"

	//"io/ioutil"
	"log"
	"os"
	"os/exec"

	"strings"

	//"strconv"
	"time"
)

type Config struct {
	Devices struct {
		AllDev  []string `json:"all"`
		UsedDev []string `json:"used"`
	}
	Ftr bool `json:"firsttimerun"`
}

func LoadConf() (Config, error) {
	var conf Config
	confFile, err := os.Open(gethomedir() + "\\config.json")
	defer confFile.Close()
	if err != nil {
		return conf, err
	}
	jsonParse := json.NewDecoder(confFile)
	jsonParse.Decode(&conf)
	return conf, err
}

func gethomedir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func getnames() (list []string) {
	var deviceList = make([]string, 2, 100)
	psScript := gethomedir() + "\\getdevices.ps1"
	getList := exec.Command("powershell.exe", "-noexit", "-file", psScript)
	pipe, err := getList.StdoutPipe()
	if err != nil {
		panic(err)
	}

	getList.Start()
	index := 0
	num := 0
	scanner := bufio.NewScanner(pipe)

	for scanner.Scan() {

		index++
		if index < 6 {
			continue
		} else {
			tekst := scanner.Text()
			if tekst == "" {
				continue
			} else {
				tekst = strings.Trim(tekst, " \t")
				if tekst == "PS"+" "+gethomedir()+">" {
					continue
				} else {
					if index >= 9 {
						deviceList = append(deviceList, tekst)
					} else {
						tekst = strings.Trim(tekst, "\t")
						deviceList[num] = tekst
						num++
					}

				}
			}
		}
	}
	fmt.Println(deviceList)
	getList.Wait()
	return deviceList
}

func firstTimeRun() (list []string) {

	var psScript string = gethomedir() + "\\getdevices.ps1"
	getList := exec.Command("powershell.exe", "-noexit", "-file", psScript)
	getList.Start()

	fmt.Println("Just a moment I'm getting the list of your devices for you....")
	time.Sleep(5 * time.Second)
	list = getnames()
	return list

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
	config, _ := LoadConf()
	config.Devices.AllDev = firstTimeRun()
	fmt.Println(config.Devices.AllDev)

}
