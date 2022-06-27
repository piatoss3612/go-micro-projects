package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

const (
	markName  = "GOLANG_CLI_REMINDER"
	markValue = "1"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage:%s <hh:mm> <text message>\n", os.Args[0])
		os.Exit(1)
	}

	now := time.Now()

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	t, err := w.Parse(os.Args[1], now)
	if err != nil {
		log.Fatal(err)
	}

	if t == nil {
		log.Fatal("Unable to parse time!")
	}

	if now.After(t.Time) {
		log.Fatal("Set a future time!")
	}

	diff := t.Time.Sub(now)

	if os.Getenv(markName) == markValue {
		time.Sleep(diff)
		err = beeep.Alert("Reminder", strings.Join(os.Args[2:], " "), "assets/information.png")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))
		if err = cmd.Start(); err != nil {
			log.Fatal(err)
		}

		log.Println("Reminder will be displayed after", diff.Round(time.Second))
		os.Exit(0)
	}
}
