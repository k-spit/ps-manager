package utils

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func Openbrowser(url, browser string) {
	if browser == "chrome" {
		//cmd := exec.Command("bash", "-c", "google-chrome-stable --incognito --disable-web-security --user-data-dir=\"/home/desktop/chrome-disabled-web-security\" "+url) //Linux example, its tested
		cmd := exec.Command("bash", "-c", "google-chrome-stable --incognito "+url) //Linux example, its tested
		cmd.Run()
	} else {
		var err error

		switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", url).Start()
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		case "darwin":
			err = exec.Command("open", url).Start()
		default:
			err = fmt.Errorf("unsupported platform")
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("browser open")
}
