package api

import (
	"github.com/dcos/dcos-diagnostics/config"
	"github.com/dcos/dcos-diagnostics/dcos"
)

// httpResponse a structure of http response from a remote host.
type httpResponse struct {
	Status int
	Units  []dcos.Unit
	Node   dcos.Node
}

// UnitsHealthResponseJSONStruct json response /system/health/v1
type UnitsHealthResponseJSONStruct struct {
	Array       []HealthResponseValues `json:"units"`
	Hostname    string                 `json:"hostname"`
	IPAddress   string                 `json:"ip"`
	DcosVersion string                 `json:"dcos_version"`
	Role        string                 `json:"node_role"`
	MesosID     string                 `json:"mesos_id"`
	TdtVersion  string                 `json:"dcos_diagnostics_version"`
}

// HealthResponseValues is a health values json response.
type HealthResponseValues struct {
	UnitID     string      `json:"id"`
	UnitHealth dcos.Health `json:"health"`
	UnitOutput string      `json:"output"`
	UnitTitle  string      `json:"description"`
	Help       string      `json:"help"`
	PrettyName string      `json:"name"`
}

// UnitsResponseJSONStruct contains health overview, collected from all hosts
type UnitsResponseJSONStruct struct {
	Array []UnitResponseFieldsStruct `json:"units"`
}

// UnitResponseFieldsStruct contains systemd unit health report.
type UnitResponseFieldsStruct struct {
	UnitID     string      `json:"id"`
	PrettyName string      `json:"name"`
	UnitHealth dcos.Health `json:"health"`
	UnitTitle  string      `json:"description"`
}

// NodesResponseJSONStruct contains an array of responses from nodes.
type NodesResponseJSONStruct struct {
	Array []*NodeResponseFieldsStruct `json:"nodes"`
}

// NodeResponseFieldsStruct contains a response from a node.
type NodeResponseFieldsStruct struct {
	HostIP     string      `json:"host_ip"`
	NodeHealth dcos.Health `json:"health"`
	NodeRole   string      `json:"role"`
}

// NodeResponseFieldsWithErrorStruct contains node response with errors.
type NodeResponseFieldsWithErrorStruct struct {
	HostIP     string      `json:"host_ip"`
	NodeHealth dcos.Health `json:"health"`
	NodeRole   string      `json:"role"`
	UnitOutput string      `json:"output"`
	Help       string      `json:"help"`
}

// Agent response json format
type agentsResponse struct {
	Agents []struct {
		Hostname   string `json:"hostname"`
		Attributes struct {
			PublicIP string `json:"public_ip"`
		} `json:"attributes"`
	} `json:"slaves"`
}

type exhibitorNodeResponse struct {
	Code        int
	Description string
	Hostname    string
	IsLeader    bool
}

// Dt is a struct of dependencies used in dcos-diagnostics code. There are 2 implementations, the one runs on a real system and
// the one used for testing.
type Dt struct {
	Cfg               *config.Config
	DtDCOSTools       dcos.Tools
	DtDiagnosticsJob  *DiagnosticsJob
	RunPullerChan     chan bool
	RunPullerDoneChan chan bool
	SystemdUnits      *SystemdUnits
	MR                *MonitoringResponse
}

type bundle struct {
	File string `json:"file_name"`
	Size int64  `json:"file_size"`
}

// UnitPropertiesResponse is a structure to unmarshal dbus.GetunitProperties response
type UnitPropertiesResponse struct {
	ID             string `json:"Id"`
	LoadState      string
	ActiveState    string
	SubState       string
	Description    string
	ExecMainStatus int

	InactiveExitTimestampMonotonic  uint64
	ActiveEnterTimestampMonotonic   uint64
	ActiveExitTimestampMonotonic    uint64
	InactiveEnterTimestampMonotonic uint64
}
