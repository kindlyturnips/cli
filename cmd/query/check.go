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
package query

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/openfga/cli/lib/cmd-utils"
	"github.com/openfga/go-sdk/client"
	"github.com/spf13/cobra"
)

func check(
	fgaClient client.SdkClient,
	user string,
	relation string,
	object string,
	contextualTuples []client.ClientTupleKey,
) (string, error) {
	body := &client.ClientCheckRequest{
		User:             user,
		Relation:         relation,
		Object:           object,
		ContextualTuples: &contextualTuples,
	}
	options := &client.ClientCheckOptions{}

	response, err := fgaClient.Check(context.Background()).Body(*body).Options(*options).Execute()
	if err != nil {
		return "", fmt.Errorf("failed to check due to %w", err)
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		return "", fmt.Errorf("failed to check due to %w", err)
	}

	return string(responseJSON), nil
}

// checkCmd represents the check command.
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check",
	Long:  "Check if a user has a particular relation with an object. E.g. \"check user:anne can_view document:roadmap\"",
	Args:  cobra.ExactArgs(3), //nolint:gomnd
	RunE: func(cmd *cobra.Command, args []string) error {
		clientConfig := cmdutils.GetClientConfig(cmd)
		fgaClient, err := clientConfig.GetFgaClient()
		if err != nil {
			return fmt.Errorf("failed to initialize FGA Client due to %w", err)
		}

		contextualTuples, err := cmdutils.ParseContextualTuples(cmd)
		if err != nil {
			return fmt.Errorf("error parsing contextual tuples for check: %w", err)
		}

		output, err := check(fgaClient, args[0], args[1], args[2], contextualTuples)
		if err != nil {
			return fmt.Errorf("error calling check: %w", err)
		}
		fmt.Print(output)

		return nil
	},
}

func init() {}