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

package agent

import (
	// Standard
	"testing"
	"time"

	// Merlin Main
	"github.com/testtoto1337/merzhin/pkg/messages"

	// Merlin
	merlinHTTP "github.com/testtoto1337/merzhin-agent/clients/http"
	testserver "github.com/testtoto1337/merzhin-agent/test/server"
)

var agentConfig = Config{
	Skew:     "100",
	Sleep:    "10s",
	MaxRetry: "7",
	KillDate: "0",
}

var clientConfig = merlinHTTP.Config{
	Protocol:    "h2",
	URL:         []string{"https://127.0.0.1:8080"},
	PSK:         "test",
	Padding:     "0",
	AuthPackage: "opaque",
}

// TestNewAgent ensure the agent.New function handles input for every agent.Config setting without error
func TestNewAgent(t *testing.T) {
	if _, err := New(agentConfig); err != nil {
		t.Error(err)
	}
}

// TestNewAgentClient ensures that the agent.clients.http.New() function handles input for every configuration setting without error
func TestNewAgentClient(t *testing.T) {
	a, err := New(agentConfig)
	if err != nil {
		t.Error(err)
	}

	// Setup Client Config
	config := clientConfig
	config.AID = a.ID
	config.UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 Mobile/15E148 Safari/604.1"
	config.JA3 = "771,49192-49191-49172-49171-159-158-57-51-157-156-61-60-53-47-49196-49195-49188-49187-49162-49161-106-64-56-50,0-10-11-13-23-65281,23-24,0"
	config.Host = "fake.cloudfront.net"
	config.Proxy = "http://127.0.0.1:8081"

	if _, err = merlinHTTP.New(config); err != nil {
		t.Error(err)
	}
}

// TestNewHTTPClient ensure the client.New function returns a http client without error
func TestNewHTTPClient(t *testing.T) {
	a, err := New(agentConfig)

	if err != nil {
		t.Error(err)
	}

	// Client config
	config := clientConfig
	config.AID = a.ID
	config.Protocol = "http"

	// Get the client
	if _, err = merlinHTTP.New(config); err != nil {
		t.Error(err)
	}
}

// TestNewHTTPClient ensure the client.New function returns a https client without error
func TestNewHTTPSClient(t *testing.T) {
	a, err := New(agentConfig)

	if err != nil {
		t.Error(err)
	}

	// Client config
	config := clientConfig
	config.AID = a.ID
	config.Protocol = "https"

	// Get the client
	if _, err = merlinHTTP.New(config); err != nil {
		t.Error(err)
	}
}

// TestNewH2CClient ensure the client.New function returns a http/2 clear-text, h2c, client without error
func TestNewH2CClient(t *testing.T) {
	a, err := New(agentConfig)

	if err != nil {
		t.Error(err)
	}

	// Client config
	config := clientConfig
	config.AID = a.ID
	config.Protocol = "h2c"

	// Get the client
	if _, err = merlinHTTP.New(config); err != nil {
		t.Error(err)
	}
}

// TestNewH2Client ensure the client.New function returns a http/2 client without error
func TestNewH2Client(t *testing.T) {
	a, err := New(agentConfig)

	if err != nil {
		t.Error(err)
	}

	// Client config
	config := clientConfig
	config.AID = a.ID
	config.Protocol = "h2"

	// Get the client
	if _, err = merlinHTTP.New(config); err != nil {
		t.Error(err)
	}
}

// TestNewHTTP3Client ensure the client.New function returns a http/3 client without error
func TestNewHTTP3Client(t *testing.T) {
	a, err := New(agentConfig)

	if err != nil {
		t.Error(err)
	}

	// Client config
	config := clientConfig
	config.AID = a.ID
	config.Protocol = "http3"

	// Get the client
	if _, err = merlinHTTP.New(config); err != nil {
		t.Error(err)
	}
}

// TestPSK ensure that the agent can't successfully communicate with the server using the wrong PSK
func TestPSK(t *testing.T) {
	a, err := New(agentConfig)

	if err != nil {
		t.Error(err)
	}
	// Get the client
	config := clientConfig
	config.AID = a.ID

	a.Client, err = merlinHTTP.New(config)
	if err != nil {
		t.Error(err)
	}

	a.WaitTime = 5000 * time.Millisecond
	err = a.Client.Set("secret", "wrongPassword")
	if err != nil {
		t.Error(err)
	}

	//signalling chans for start/end of test
	setup := make(chan struct{})
	ended := make(chan struct{})

	go testserver.TestServer{}.Start("8080", ended, setup, t)
	//wait until set up
	<-setup

	m := messages.Base{
		Version: 1.0,
		ID:      a.ID,
		Type:    messages.CHECKIN,
	}

	_, errSend := a.Client.Send(m)
	if errSend == nil {
		t.Error("Agent successfully sent an encrypted message using the wrong key")
		return
	}

	close(ended)
}

// TestOPAQUE verifies that agent is able to successfully complete the OPAQUE protocol Registration and Authentication steps
func TestOPAQUE(t *testing.T) {
	a, err := New(agentConfig)
	if err != nil {
		t.Error(err)
	}

	// Get the client
	config := clientConfig
	config.AID = a.ID
	config.URL = []string{"https://127.0.0.1:8082"}
	if a.Client, err = merlinHTTP.New(clientConfig); err != nil {
		t.Error(err)
	}

	// Setup and start test server
	setup := make(chan struct{}) // Channel to determine when the server setup has completed
	ended := make(chan struct{}) // Channel to determine when the server has quit
	go testserver.TestServer{}.Start("8082", ended, setup, t)
	<-setup //wait until set up

	// Perform client authentication which consists of both OPAQUE Registration and Authentication
	_, err = a.Client.Auth("opaque", true)
	if err != nil {
		t.Error(err)
	}

	close(ended)
}

// TestAgentInitialCheckin verifies the Agent's initialCheckin() function returns without error
func TestAgentInitialCheckIn(t *testing.T) {
	a, err := New(agentConfig)
	if err != nil {
		t.Error(err)
		return
	}
	a.WaitTime = 5000 * time.Millisecond

	// Get the client
	config := clientConfig
	config.AID = a.ID
	config.URL = []string{"https://127.0.0.1:8083/merlin"}
	a.Client, err = merlinHTTP.New(config)
	if err != nil {
		t.Error(err)
	}

	//signalling chans for start/end of test
	setup := make(chan struct{})
	ended := make(chan struct{})

	go testserver.TestServer{}.Start("8083", ended, setup, t)
	//wait until set up
	<-setup

	_, err = a.Client.Initial(a.getAgentInfoMessage())
	if err != nil {
		t.Errorf("error with initial checkin:\r\n%s", err)
	}
	close(ended)
}

// TestBadAuthentication verifies unsuccessful authentication using the wrong PSK
func TestBadAuthentication(t *testing.T) {
	a, err := New(agentConfig)

	if err != nil {
		t.Error(err)
		return
	}
	a.WaitTime = 5000 * time.Millisecond

	// Get the client
	config := clientConfig
	config.AID = a.ID
	config.URL = []string{"https://127.0.0.1:8085"}
	config.PSK = "neverGonnaGiveYouUp"
	a.Client, err = merlinHTTP.New(config)
	if err != nil {
		t.Error(err)
	}

	//signalling chans for start/end of test
	setup := make(chan struct{})
	ended := make(chan struct{})

	// AppVeyor uses port 8084 for something else
	go testserver.TestServer{}.Start("8085", ended, setup, t)
	//wait until set up
	<-setup

	_, err = a.Client.Initial(a.getAgentInfoMessage())
	if err == nil {
		t.Error("the agent successfully authenticated with the wrong PSK")
	}
	close(ended)
}

// Bad content-type header
// TODO test every function of the message handler
