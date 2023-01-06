package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ssleert/cdropp/pkg/dropper"
	"github.com/ssleert/ginip"
	"github.com/ssleert/memory"
)

const confFile = "/etc/cdropp/conf.ini"

var (
	// minimal ram amount to trigger in mb
	minTrigger = 256

	// timer between checks in ms
	timerBetweenChecks = time.Duration(3000)

	// drop strength (1, 2, 3)
	dropStrength = 3

	// if true logging is on
	debug = false
)

func printLog(s string, a ...any) {
	if debug {
		log.Printf(s, a...)
	}
}

func exitMsg(err error) {
	if debug {
		log.Fatalln("err:", err)
	}
	fmt.Println("err:", err)
	os.Exit(1)
}

// get values from ini config
func conf() error {
	confPath := os.Getenv("CDROPP_CONF_PATH")
	if confPath == "" {
		_, err := os.Stat(confFile)
		if err == nil {
			confPath = confFile
		}
	}

	ini, err := ginip.Load(confPath)
	if err != nil {
		return err
	}
	confValue, _ := ini.GetValueInt("", "trigger")
	if confValue != 0 {
		minTrigger = confValue
	}
	confValue, _ = ini.GetValueInt("", "timer")
	if confValue != 0 {
		timerBetweenChecks = time.Duration(confValue)
	}
	confValue, _ = ini.GetValueInt("", "strength")
	if confValue != 0 {
		dropStrength = confValue
	}

	confValue, _ = ini.GetValueInt("", "debug")
	if confValue == 0 {
		debug = false
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		exitMsg(errors.New(" no args defined"))
	}

	switch os.Args[1] {
	case "--daemon", "-d":
		if dropper.NoPermissions() {
			exitMsg(errors.New("open /proc/sys/vm/drop_caches: permission denied. try to start cdropp as root"))
		}

		err := conf()
		if err != nil {
			exitMsg(err)
		}

		tiker := time.NewTicker(timerBetweenChecks * time.Millisecond)
		var panicState int
		for range tiker.C {
			free, err := memory.GetFreeRam()
			if err != nil {
				if panicState >= 10 {
					exitMsg(err)
				}
				panicState++
				continue
			}

			if free < minTrigger {
				err = dropper.Drop(dropStrength)
				if err != nil {
					if panicState >= 10 {
						exitMsg(err)
					}
					panicState++
					continue
				}
				printLog("cache dropped | panic: %v", panicState)
			}
			printLog("cache checked | panic: %d", panicState)

		}

	case "--check", "-c":
		ram, err := memory.GetRam()
		if err != nil {
			exitMsg(err)
		}
		fmt.Println("ram info:")
		fmt.Printf(" total     = %d mb\n", ram.Total)
		fmt.Printf(" free      = %d mb\n", ram.Free)
		fmt.Printf(" availible = %d mb\n", ram.Available)

	case "--version", "-v":
		fmt.Println("cdropp - 0.0.1")

	case "--help", "-h":
		fmt.Println("cdropp - simple deamon for dropping caches in ram")
		fmt.Println(" -d --daemon  | start main loop")
		fmt.Println(" -c --check   | check current ram usage")
		fmt.Println(" -v --version | print program version")
		fmt.Println(" -h --help    | print help message")

	default:

	}
}
