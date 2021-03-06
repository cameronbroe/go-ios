package diagnostics

import (
	"bytes"
	plist "howett.net/plist"
)

type diagnosticsRequest struct {
	Request string
}

type diagnosticsStatus struct {
	Status string
}

func diagnosticsStatusfromBytes(plistBytes []byte) diagnosticsStatus {
	decoder := plist.NewDecoder(bytes.NewReader(plistBytes))
	var data diagnosticsStatus
	_ = decoder.Decode(&data)
	return data
}

func diagnosticsRequestfromBytes(plistBytes []byte) diagnosticsRequest {
	decoder := plist.NewDecoder(bytes.NewReader(plistBytes))
	var data diagnosticsRequest
	_ = decoder.Decode(&data)
	return data
}

func diagnosticsfromBytes(plistBytes []byte) allDiagnosticsResponse {
	decoder := plist.NewDecoder(bytes.NewReader(plistBytes))
	var data allDiagnosticsResponse
	_ = decoder.Decode(&data)
	return data
}

type allDiagnosticsResponse struct {
	Diagnostics Diagnostics
	Status      string
}
type Diagnostics struct {
	GasGauge GasGauge
	HDMI     HDMI
	NAND     NAND
	WiFi     WiFi
}
type WiFi struct {
	Active string
	Status string
}
type NAND struct {
	Status string
}
type HDMI struct {
	Connection string
	Status     string
}
type GasGauge struct {
	CycleCount         uint64
	DesignCapacity     uint64
	FullChargeCapacity uint64
	Status             string
}
