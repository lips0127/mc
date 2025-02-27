// Copyright (c) 2015-2022 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import "github.com/minio/cli"

var adminIDPLdapRemoveCmd = cli.Command{
	Name:         "remove",
	Usage:        "remove LDAP IDP server configuration",
	Action:       mainAdminIDPLDAPRemove,
	Before:       setGlobalsFromContext,
	Flags:        globalFlags,
	OnUsageError: onUsageError,
	CustomHelpTemplate: `NAME:
  {{.HelpName}} - {{.Usage}}

USAGE:
  {{.HelpName}} TARGET [CFG_NAME]

FLAGS:
  {{range .VisibleFlags}}{{.}}
  {{end}}
EXAMPLES:
  1. Remove the default LDAP IDP configuration (CFG_NAME is omitted).
     {{.Prompt}} {{.HelpName}} play/
`,
}

func mainAdminIDPLDAPRemove(ctx *cli.Context) error {
	if len(ctx.Args()) < 1 || len(ctx.Args()) > 2 {
		showCommandHelpAndExit(ctx, 1)
	}

	args := ctx.Args()

	var cfgName string
	if len(args) == 2 {
		cfgName = args.Get(1)
	}
	return adminIDPRemove(ctx, false, cfgName)
}
