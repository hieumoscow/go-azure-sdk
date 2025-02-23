package networkmanagereffectivesecurityadminrules

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EffectiveBaseSecurityAdminRule interface {
}

func unmarshalEffectiveBaseSecurityAdminRuleImplementation(input []byte) (EffectiveBaseSecurityAdminRule, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EffectiveBaseSecurityAdminRule into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Default") {
		var out EffectiveDefaultSecurityAdminRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EffectiveDefaultSecurityAdminRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Custom") {
		var out EffectiveSecurityAdminRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EffectiveSecurityAdminRule: %+v", err)
		}
		return out, nil
	}

	type RawEffectiveBaseSecurityAdminRuleImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawEffectiveBaseSecurityAdminRuleImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
