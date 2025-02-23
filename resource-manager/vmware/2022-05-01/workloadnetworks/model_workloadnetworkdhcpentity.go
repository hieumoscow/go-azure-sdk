package workloadnetworks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworkDhcpEntity interface {
}

func unmarshalWorkloadNetworkDhcpEntityImplementation(input []byte) (WorkloadNetworkDhcpEntity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkloadNetworkDhcpEntity into map[string]interface: %+v", err)
	}

	value, ok := temp["dhcpType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "RELAY") {
		var out WorkloadNetworkDhcpRelay
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkloadNetworkDhcpRelay: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SERVER") {
		var out WorkloadNetworkDhcpServer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkloadNetworkDhcpServer: %+v", err)
		}
		return out, nil
	}

	type RawWorkloadNetworkDhcpEntityImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawWorkloadNetworkDhcpEntityImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
