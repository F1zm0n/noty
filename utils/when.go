package utils

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"sync"
	"time"
)

//	func SetAlarm(info *AlarmInfo, diff *time.Duration) error {
//		fmt.Println("alarm start")
//		go func(info *AlarmInfo, diff time.Duration) {
//			fmt.Println("Alarm is started", info.Time, info.Name)
//			for {
//				select {
//				case <-time.After(diff):
//					err := beeep.Alert(info.Name, info.Message, "assets/information.png")
//					if err != nil {
//						fmt.Println("error sending alert message")
//						return
//					}
//					err = DeleteDataFromFile(info.Name)
//					if err != nil {
//						fmt.Println("error deleting data from file")
//						return
//					}
//				case name := <-info.Stop:
//					if name == info.Name {
//						fmt.Printf("An alarm %s stopped", info.Name)
//						return
//					}
//				}
//			}
//
//		}(info, *diff)
//		return nil
//	}
func SetAlarm(info *AlarmInfo, diff *time.Duration, wg *sync.WaitGroup) error {
	fmt.Println("Alarm start")
	wg.Add(1)

	go func(info *AlarmInfo, diff time.Duration, wg *sync.WaitGroup) {
		fmt.Println("Alarm is started", info.Time, info.Name)
		defer wg.Done()

		select {
		case <-time.After(diff):
			err := beeep.Alert(info.Name, info.Message, "assets/information.png")
			if err != nil {
				fmt.Println("Error sending alert message:", err)
				return
			}
			err = DeleteDataFromFile(info.Name)
			if err != nil {
				fmt.Println("Error deleting data from file:", err)
				return
			}
			return
		case name := <-info.Stop:
			if name == info.Name {
				fmt.Printf("An alarm %s stopped\n", info.Name)
				return
			}
		}
	}(info, *diff, wg)
	fmt.Println("finished")
	return nil
}

func TimeParser(info *AlarmInfo, now time.Time) (AlarmData, *time.Duration, error) {
	w := when.New(nil)

	w.Add(en.All...)

	w.Add(common.All...)
	t, err := w.Parse(info.Time, now)
	if err != nil {
		return AlarmData{}, nil, err
	}
	if t == nil {
		return AlarmData{}, nil, UnableToParseTimeErr
	}

	if now.After(t.Time) {
		return AlarmData{}, nil, PastTimeErr
	}
	fmt.Println(t.Time)
	layout := "2006-01-02 15:04:05.9999 -0700 -07"

	parsedTime, err := time.Parse(layout, t.Time.String())
	if err != nil {
		return AlarmData{}, nil, err
	}
	Data := &AlarmData{
		Time:    parsedTime,
		Name:    info.Name,
		Message: info.Message,
	}
	diff := t.Time.Sub(now)
	return *Data, &diff, nil
}
