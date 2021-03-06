/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"os"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/nikhilfernandes/mqtt-consumer/cmd"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_WRONLY | os.O_CREATE, 0755)
    if err != nil {
    	fmt.Println("error in opening log file")
        log.Fatal(err)
    }
    defer file.Close()
    log.SetOutput(file)
    log.SetFormatter(&log.JSONFormatter{})
    log.SetLevel(log.InfoLevel)
	cmd.Execute()
}

