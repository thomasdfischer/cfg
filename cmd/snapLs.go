// Copyright © 2018 Thomas Fischer
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
	"github.com/spf13/cobra"

	"github.com/Confbase/cfg/snap"
)

var snapLineMode bool

// snapLsCmd represents the snapLs command
var snapLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List snapshots",
	Long: `Lists snapshots.`,
	Run: func(cmd *cobra.Command, args []string) {
		snap.Ls(snapLineMode)
	},
}

func init() {
	snapCmd.AddCommand(snapLsCmd)
	snapLsCmd.Flags().BoolVarP(&snapLineMode, "lines", "l", false, "print one snap per line")
}
