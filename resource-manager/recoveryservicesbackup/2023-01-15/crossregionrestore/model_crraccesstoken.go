package crossregionrestore

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrrAccessToken interface {
}

func unmarshalCrrAccessTokenImplementation(input []byte) (CrrAccessToken, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CrrAccessToken into map[string]interface: %+v", err)
	}

	value, ok := temp["objectType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "WorkloadCrrAccessToken") {
		var out WorkloadCrrAccessToken
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkloadCrrAccessToken: %+v", err)
		}
		return out, nil
	}

	type RawCrrAccessTokenImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawCrrAccessTokenImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
