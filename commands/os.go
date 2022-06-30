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

import "github.com/testtoto1337/merzhin-agent/cli"

// Setup is used to prepare the environment or context for subsequent commands and is specific to each operating system
func Setup() error {
	cli.Message(cli.DEBUG, "entering Setup() function from the commands.os package")
	return nil
}

// TearDown is the opposite of Setup and removes and environment or context applications
func TearDown() error {
	cli.Message(cli.DEBUG, "entering TearDown() function from the commands.os package")
	return nil
}
