package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start pomodoro timer",
	Long: `Start pomodoro timer with custom timer.

  Example: pomo start -t 90
  starts the pomodoro timer to countdown from 90 minutes.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },

	Run: startTimer,
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().IntP("timer", "t", 120, "Duration of timer")
}

func startTimer(cmd *cobra.Command, args []string) {
	timer, _ := cmd.Flags().GetInt("timer")
	timerSeconds := timer * 60

	fmt.Println("Begin!")

	for i := range timerSeconds {
		remaining := timerSeconds - i
		minutes := remaining / 60
		seconds := remaining % 60

		fmt.Printf("\r%02d:%02d", minutes, seconds)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Pomodoro completed!")
}
