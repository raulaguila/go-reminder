package main

import (
	"fmt"
	"os/exec"
	"time"
)

type LinuxNotificator struct {
	Title       string
	Description string
}

func (l *LinuxNotificator) Notify() error {
	return exec.Command("notify-send", l.Title, l.Description).Run()
}

func (l *LinuxNotificator) Sound() error {
	return exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/alarm-clock-elapsed.oga").Run()
}

func (l *LinuxNotificator) Reminder(every time.Duration, from time.Time) {

	for {
		now := time.Now()

		if (now.Unix()-from.Unix())%int64(every.Seconds()) == 0 {
			go l.Notify()
			go l.Sound()

			sleeptime := now.Add(every).Add(-500 * time.Millisecond)
			sleep := time.Since(sleeptime)
			time.Sleep(sleep.Abs())
		}
	}
}

func main() {

	title := "Alert title!"
	description := "Alert description!"
	every := 1 * time.Second
	from := time.Now()

	notificator := LinuxNotificator{
		Title:       title,
		Description: description,
	}

	go notificator.Reminder(every, from)

	var input string
	fmt.Scanln(&input)
}
