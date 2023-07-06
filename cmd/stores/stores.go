/*
Copyright © 2023 OpenFGA

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package stores

import (
	"github.com/spf13/cobra"
)

// StoresCmd represents the store command.
var StoresCmd = &cobra.Command{
	Use:   "stores",
	Short: "Create, Get, Delete and List OpenFGA Stores",
}

func init() {
	StoresCmd.AddCommand(createCmd)
	StoresCmd.AddCommand(listCmd)
	StoresCmd.AddCommand(getCmd)
	StoresCmd.AddCommand(deleteCmd)
}