// Copyright © 2017 Sascha Andres <sascha.andres@outlook.com>
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
	"github.com/sascha-andres/go-logsink/v2/internal/console"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listenCmd represents the listen command
var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Start a server instance of go-logsink",
	Long: `This command is used to create a go-logsink server.
Call it to have clients forward log messages here.

If debug mode is enable you call open /debug/statsviz/ in your browser`,
	Run: func(cmd *cobra.Command, args []string) {
		handleLock(console.Listen)
	},
}

func init() {
	RootCmd.AddCommand(listenCmd)
	listenCmd.Flags().StringP("bind", "b", ":50051", "Provide bind definition")
	viper.BindPFlag("listen.bind", listenCmd.Flags().Lookup("bind"))
}
