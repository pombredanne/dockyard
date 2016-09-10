/*
Copyright 2015 The ContainerOps Authors All rights reserved.

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

package cmd

import (
	"github.com/urfave/cli"
)

var CmdMonitor = cli.Command{
	Name:        "monitor",
	Usage:       "monitor utils for all service",
	Description: "Monitor service health, database status, object storage service and so on.",
	Action:      runMonitor,
	Flags:       []cli.Flag{},
}

func runMonitor(c *cli.Context) error {
	startServerService()
	return nil
}
