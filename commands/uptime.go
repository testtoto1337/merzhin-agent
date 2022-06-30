// +build !windows

// Merlin is a post-exploitation command and control framework.
// This file is part of Merlin.
// Copyright (C) 2022  Russel Van Tuyl

// Merlin is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// any later version.

// Merlin is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Merlin.  If not, see <http://www.gnu.org/licenses/>.

package commands

import (

	// Merlin
	"github.com/testtoto1337/merzhin-agent/cli"
	"github.com/testtoto1337/merzhin/pkg/jobs"
)

// Uptime retrieves the system's uptime
// Windows only
func Uptime() jobs.Results {
	cli.Message(cli.DEBUG, "entering Uptime()")
	return jobs.Results{
		Stderr: "the Uptime command is not supported by this agent type",
	}
}
