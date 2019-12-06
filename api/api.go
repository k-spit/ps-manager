package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"ps-manager/models"
	"ps-manager/utils"
	"strings"
)

var cpu = 0
var pid = 0
var user = 0
var command = 0
var switcher = "--sort -pcpu"

//NewAPI creates api and starts router
func NewAPI() {
	http.HandleFunc("/getProcesses", getProcessesHandler())
	http.HandleFunc("/postPid", postPidProcessHandler())
	http.HandleFunc("/postCommand", postCommandProcessHandler())

	http.HandleFunc("/cpu", cpuHandler(cpu))
	http.HandleFunc("/pid", pidHandler(cpu))
	http.HandleFunc("/user", userHandler(cpu))
	http.HandleFunc("/command", commandHandler(cpu))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func cpuHandler(cpu int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if utils.Odd(cpu) {
			switcher = "--sort -pcpu"
		}
		if utils.Even(cpu) {
			switcher = "--sort +pcpu"
		}
		cpu++
	})
}

func pidHandler(pid int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if utils.Odd(pid) {
			switcher = "--sort -pid"
		}
		if utils.Even(pid) {
			switcher = "--sort +pid"
		}
		pid++
	})
}

func userHandler(user int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if utils.Odd(user) {
			switcher = "--sort -user"
		}
		if utils.Even(user) {
			switcher = "--sort +user"
		}
		user++
	})
}
func commandHandler(command int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if utils.Odd(command) {
			switcher = "--sort -command"
		}
		if utils.Even(command) {
			switcher = "--sort +command"
		}
		command++
	})
}

func postPidProcessHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
		oldPid := string(body)
		slittedPid := strings.Split(oldPid, "=")

		cmd := exec.Command("bash", "-c", "kill -9 "+slittedPid[1])
		cmd.Run()
	})
}

func postCommandProcessHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
		command := string(body)
		splitEqualChar := strings.Split(command, "=")
		replaceSlashChar := strings.Replace(splitEqualChar[1], "%2F", "/", -1)
		cmd := exec.Command("bash", "-c", replaceSlashChar)
		cmd.Run()
	})
}

func getProcessesHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		out, err := exec.Command("bash", "-c", "ps -eo pcpu,pid,user,command --no-headers "+switcher).Output()
		if err != nil {
			go log.Println("could not run os command!")
			return
		}
		// convert byte[] to string
		pscommand := string(out[:])
		pscommandSlice := strings.Split(pscommand, "\n")
		processes := []models.Process{}

		for _, ps := range pscommandSlice {
			s := strings.Fields(ps)
			if len(s) >= 4 {
				process := new(models.Process)
				process.CPU = s[0]
				process.Pid = s[1]
				process.User = s[2]
				process.Command = s[3]
				processes = append(processes, *process)
			}
		}
		urlsJSON, err := json.Marshal(processes)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Write([]byte(string(urlsJSON)))
	})
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
