/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/F1zm0n/noty/utils"
	"github.com/spf13/cobra"
	"sync"
	"time"
)

var (
	times   string
	name    string
	message string
)

// whenCmd represents the when command
var whenCmd = &cobra.Command{
	Use:   "when",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup

		DataMap := make(map[string]utils.AlarmData)

		now := time.Now()
		if name == "" || message == "" || times == "" {
			fmt.Println("Pass all three flags with value to create an alarm (name,message,time)")
		}
		info := &utils.AlarmInfo{
			Name:    name,
			Message: message,
			Time:    times,
			Stop:    make(chan string),
		}
		data, diff, err := utils.TimeParser(info, now)
		if err != nil {
			fmt.Println("couldn't parse given time, please retry, ", err)
			return
		}
		fmt.Println(diff)
		DataMap[data.Name] = data
		err = utils.WriteFileWithJSON(DataMap)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = utils.SetAlarm(info, diff, &wg)

		if err != nil {
			fmt.Println(err)
			return
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(whenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// whenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// whenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	whenCmd.PersistentFlags().StringVarP(&times, "time", "t", "", "Specify time for an alarm")
	whenCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Specify name for an alarm")
	whenCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "Specify message for an alarm")

}
