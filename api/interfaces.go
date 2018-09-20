package api

import (
	"time"
)

// DCOSHelper DC/OS specific tools interface.
type DCOSHelper interface {
	// open dbus connection
	InitializeUnitControllerConnection() error

	// close dbus connection
	CloseUnitControllerConnection() error

	// function to get Connection.GetUnitProperties(pname)
	// returns a maps of properties https://github.com/coreos/go-systemd/blob/master/dbus/methods.go#L176
	GetUnitProperties(string) (map[string]interface{}, error)

	// A wrapper to /opt/mesosphere/bin/detect_ip script
	// should return empty string if script fails.
	DetectIP() (string, error)

	// get system's hostname
	GetHostname() (string, error)

	// Detect node role: master/agent
	GetNodeRole() (string, error)

	// Get DC/OS systemd units on a system
	GetUnitNames() ([]string, error)

	// Get journal output
	GetJournalOutput(string) (string, error)

	// Get mesos node id, first argument is a function to determine a role.
	GetMesosNodeID() (string, error)

	// Get makes HTTP GET request, return read arrays of bytes
	Get(string, time.Duration) ([]byte, int, error)

	// Post makes HTTP GET request, return read arrays of bytes
	Post(string, time.Duration) ([]byte, int, error)

	// LookupMaster will lookup a masters in DC/OS cluster.
	// Initial lookup will be done by making HTTP GET request to exhibitor.If GET request fails, the next lookup
	// will failover to history service for one minute, it this fails or no nodes found, masters will be looked up
	// in history service for last hour.
	GetMasterNodes() ([]Node, error)
	//
	//// GetAgentsFromMaster will lookup agents in DC/OS cluster.
	GetAgentNodes() ([]Node, error)

	// Get timestamp
	GetTimestamp() time.Time
}
