package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"ps-manager/models"
	"strings"
)

var newPath = ""

func main() {
	http.HandleFunc("/getProcesses", getProcessesHandler())
	http.HandleFunc("/postPid", postPidProcessHandler())
	http.HandleFunc("/postCommand", postCommandProcessHandler())

	//openbrowser("index.html")
	log.Fatal(http.ListenAndServe(":8080", nil))
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

		//out, err := exec.Command("bash", "-c", "ps -eo pcpu,pid,user,command --no-headers | sort -t. -nk1,2 -k4,4 -r").Output()
		out, err := exec.Command("bash", "-c", "ps -eo pcpu,pid,user,command --no-headers --sort -pid").Output()
		if err != nil {
			go log.Println("could not run os command!")
			return
		}
		file, err := os.Create("tmp")
		if err != nil {
			fmt.Println(err)
		} else {
			file.WriteString(string(out))
		}
		file.Close()

		file, err = os.Open("tmp")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		processes := []models.Process{}

		for scanner.Scan() {
			tmp := scanner.Text()
			s := strings.Fields(tmp)

			process := new(models.Process)
			process.CPU = s[0]
			process.Pid = s[1]
			process.User = s[2]
			process.Command = s[3]

			processes = append(processes, *process)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
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

func openbrowser(url string) {
	cmd := exec.Command("bash", "-c", "google-chrome-stable --disable-web-security --user-data-dir=\"/home/desktop/chrome-disabled-web-security\" /home/desktop/git/ps-manager/cmd/ps-manager/index.html") //Linux example, its tested
	cmd.Run()

	// var err error

	// switch runtime.GOOS {
	// case "linux":
	// 	err = exec.Command("xdg-open", url).Start()
	// case "windows":
	// 	err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	// case "darwin":
	// 	err = exec.Command("open", url).Start()
	// default:
	// 	err = fmt.Errorf("unsupported platform")
	// }
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("browser open")
}
