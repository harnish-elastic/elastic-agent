// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build !darwin && !windows
// +build !darwin,!windows

package paths

const (
	// BinaryName is the name of the installed binary.
	BinaryName = "elastic-agent"

	// InstallPath is the installation path using for install command.
	InstallPath = "/opt/Elastic/Agent"

	// ControlSocketPath is the control socket path used when installed.
	ControlSocketPath = "unix:///run/elastic-agent.sock"

	// ShipperSocketPipePattern is the socket path used when installed for a shipper pipe.
	ShipperSocketPipePattern = "unix:///run/elastic-agent-%s-pipe.sock"

	// ServiceName is the service name when installed.
	ServiceName = "elastic-agent"

	// ShellWrapperPath is the path to the installed shell wrapper.
	ShellWrapperPath = "/usr/bin/elastic-agent"

	// ShellWrapper is the wrapper that is installed.
	ShellWrapper = `#!/bin/sh
exec /opt/Elastic/Agent/elastic-agent $@
`
)

// ArePathsEqual determines whether paths are equal taking case sensitivity of os into account.
func ArePathsEqual(expected, actual string) bool {
	return expected == actual
}
