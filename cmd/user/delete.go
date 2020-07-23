// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package cmduser

import (
	"os"
	"sync"
	"time"

	"github.com/elastic/cloud-sdk-go/pkg/api/userapi"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	sdkcmdutil "github.com/elastic/cloud-sdk-go/pkg/util/cmdutil"
	"github.com/spf13/cobra"

	cmdutil "github.com/elastic/ecctl/cmd/util"
	"github.com/elastic/ecctl/pkg/ecctl"
)

var deleteCmd = &cobra.Command{
	Use:     "delete <user name> <user name>...",
	Short:   cmdutil.AdminReqDescription("Deletes one or more platform users"),
	PreRunE: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			var msg = "Multiple users will be deleted, do you want to continue? [y/n]: "
			if !sdkcmdutil.ConfirmAction(msg, os.Stdin, ecctl.Get().Config.OutputDevice) {
				return nil
			}
		}

		var merr = multierror.NewPrefixed("delete user")
		var wg sync.WaitGroup
		for i := range args {
			wg.Add(1)
			go func(index int) {
				merr = merr.Append(
					userapi.Delete(userapi.DeleteParams{
						API:      ecctl.Get().API,
						UserName: args[index],
					}),
				)
				wg.Done()
			}(i)

			// Only delete a user per second as anything below causes a 500
			<-time.After(time.Second)
		}

		wg.Wait()
		return merr.ErrorOrNil()
	},
}
