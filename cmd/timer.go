package cmd

import (
	"Boat/internal/timer"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var calculateTime string
var duration string

// timerCmd represents the timer command
var timerCmd = &cobra.Command{
	Use:   "timer",
	Short: "时间格式处理",
	Long:  `时间格式处理`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

// currentTimeCmd current time
var currentTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  `获取当前时间`,
	Run: func(cmd *cobra.Command, args []string) {
		now := timer.GetCurrentTime()
		log.Printf("输出结果: %s, %d", now.Format("2006-01-02 15:04:05"), now.Unix())
	},
}

// calculateTimeCmd calculate time
var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetCurrentTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")

			if space == 0 {
				layout = "2006-01-02"
			}

			if space == 1 {
				layout = "2006-01-02 15:04"
			}

			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer error: %v", err)
		}

		log.Printf("输出结果: %s, %d", t.Format(layout), t.Unix())
	},
}

func init() {
	timerCmd.AddCommand(currentTimeCmd)
	timerCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间，有效单位为时间戳或已格式化后的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效时间单位为"ns", "us" (or "µ s"), "ms", "s", "m", "h"`)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
