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

package cmdproxy

import (
	"path/filepath"

	"github.com/elastic/cloud-sdk-go/pkg/api/platformapi/proxyapi"
	"github.com/spf13/cobra"

	cmdutil "github.com/elastic/ecctl/cmd/util"
	"github.com/elastic/ecctl/pkg/ecctl"
)

var listProxiesCmd = &cobra.Command{
	Use:     "list",
	Short:   cmdutil.AdminReqDescription("Returns all of the proxies in the platform"),
	PreRunE: cobra.MaximumNArgs(0),
	RunE:    listProxies,
}

func listProxies(cmd *cobra.Command, args []string) error {
	a, err := proxyapi.List(proxyapi.ListParams{
		API:    ecctl.Get().API,
		Region: ecctl.Get().Config.Region,
	})
	if err != nil {
		return err
	}

	return ecctl.Get().Formatter.Format(filepath.Join("proxy", "list"), a)
}

func init() {
	Command.AddCommand(listProxiesCmd)
}
