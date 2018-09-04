package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	ApplicationName  = "search-service"
	ApplicationTitle = "A search service"
	Version          = "0.1"
)

var (
	GitHash   string // should be uninitialized
	BuildTime string // should be uninitialized
)

var RootCmd = &cobra.Command{
	Use:   ApplicationName,
	Short: ApplicationTitle,
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ApplicationInfo())
	},
}

func ApplicationInfo() string {
	return ApplicationTitle + " " + Version + " [" + GitHash + "] " + BuildTime
}
