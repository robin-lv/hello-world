package cobrax

import (
	"github.com/spf13/cobra"
	"strings"
)

const (
	dash       = "-"
	doubleDash = "--"
	assign     = "="
)

func supportGoStdFlag(rootCmd *Command, args []string) []string {
	copyArgs := append([]string(nil), args...)
	parentCmd, _, err := rootCmd.Traverse(args[:1])
	if err != nil { // ignore it to let cobra handle the error.
		return copyArgs
	}

	for idx, arg := range copyArgs[0:] {
		parentCmd, _, err = parentCmd.Traverse([]string{arg})
		if err != nil { // ignore it to let cobra handle the error.
			break
		}
		if !strings.HasPrefix(arg, dash) {
			continue
		}

		flagExpr := strings.TrimPrefix(arg, doubleDash)
		flagExpr = strings.TrimPrefix(flagExpr, dash)
		flagName, flagValue := flagExpr, ""
		assignIndex := strings.Index(flagExpr, assign)
		if assignIndex > 0 {
			flagName = flagExpr[:assignIndex]
			flagValue = flagExpr[assignIndex:]
		}

		if !isBuiltin(flagName) {
			// The method Flag can only match the user custom flags.
			f := parentCmd.Flag(flagName)
			if f == nil {
				continue
			}
			if f.Shorthand == flagName {
				continue
			}
		}

		goStyleFlag := doubleDash + flagName
		if assignIndex > 0 {
			goStyleFlag += flagValue
		}

		copyArgs[idx] = goStyleFlag
	}
	return copyArgs
}

func isBuiltin(name string) bool {
	return name == "version" || name == "help"
}
func getCommandName(cmd *cobra.Command) string {
	if cmd.HasParent() {
		return getCommandName(cmd.Parent()) + "." + cmd.Name()
	}
	return cmd.Name()
}

func getCommandsRecursively(parent *cobra.Command) []*cobra.Command {
	var commands []*cobra.Command
	for _, cmd := range parent.Commands() {
		commands = append(commands, cmd)
		commands = append(commands, getCommandsRecursively(cmd)...)
	}
	return commands
}
