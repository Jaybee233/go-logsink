// Copyright © 2017 Sascha Andres <sascha.andres@outlook.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "20170124"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print out the currect version",
	Long: `Print out the version.

Version tag has format YYYYMMDD, where YYYY is the year, MM the month and DD the
day of the release.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("go-logsink version %s\n", version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
