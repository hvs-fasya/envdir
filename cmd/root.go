package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/hvs-fasya/envdir/internal/env"
)

var RootCmd = &cobra.Command{
	Use:   "go-envdir dir child",
	Short: `go-envdir - runs another program with environment modified according to files in a specified directory`,
	Long: "go-envdir runs another program with environment modified according to files in a specified directory. \n" +
		"dir is a single argument. child consists of one or more arguments. \n" +
		"go-envdir sets various environment variables as specified by files in the directory named  dir. \n" +
		"It then runs child.",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := env.SetEnvs(args[0])
		if err != nil {
			log.Fatal("set env vars error: " + err.Error())
		}
		//todo: args
		var childArgs = []string{}
		if len(args) > 2 {
			childArgs = args[2:]
		}
		command := exec.Command(args[1], childArgs...)
		var out bytes.Buffer
		command.Stdout = &out
		err = command.Run()
		if err != nil {
			log.Fatal("execute child program error: " + err.Error())
		}
		fmt.Printf("%s", out.String())
	},
}
