package diagnostic

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlwaysLog string

const (
	AlwaysLogAllErrors AlwaysLog = "allErrors"
)

func PossibleValuesForAlwaysLog() []string {
	return []string{
		string(AlwaysLogAllErrors),
	}
}

func parseAlwaysLog(input string) (*AlwaysLog, error) {
	vals := map[string]AlwaysLog{
		"allerrors": AlwaysLogAllErrors,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlwaysLog(input)
	return &out, nil
}

type DataMaskingMode string

const (
	DataMaskingModeHide DataMaskingMode = "Hide"
	DataMaskingModeMask DataMaskingMode = "Mask"
)

func PossibleValuesForDataMaskingMode() []string {
	return []string{
		string(DataMaskingModeHide),
		string(DataMaskingModeMask),
	}
}

func parseDataMaskingMode(input string) (*DataMaskingMode, error) {
	vals := map[string]DataMaskingMode{
		"hide": DataMaskingModeHide,
		"mask": DataMaskingModeMask,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataMaskingMode(input)
	return &out, nil
}

type HTTPCorrelationProtocol string

const (
	HTTPCorrelationProtocolLegacy  HTTPCorrelationProtocol = "Legacy"
	HTTPCorrelationProtocolNone    HTTPCorrelationProtocol = "None"
	HTTPCorrelationProtocolWThreeC HTTPCorrelationProtocol = "W3C"
)

func PossibleValuesForHTTPCorrelationProtocol() []string {
	return []string{
		string(HTTPCorrelationProtocolLegacy),
		string(HTTPCorrelationProtocolNone),
		string(HTTPCorrelationProtocolWThreeC),
	}
}

func parseHTTPCorrelationProtocol(input string) (*HTTPCorrelationProtocol, error) {
	vals := map[string]HTTPCorrelationProtocol{
		"legacy": HTTPCorrelationProtocolLegacy,
		"none":   HTTPCorrelationProtocolNone,
		"w3c":    HTTPCorrelationProtocolWThreeC,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HTTPCorrelationProtocol(input)
	return &out, nil
}

type OperationNameFormat string

const (
	OperationNameFormatName OperationNameFormat = "Name"
	OperationNameFormatUrl  OperationNameFormat = "Url"
)

func PossibleValuesForOperationNameFormat() []string {
	return []string{
		string(OperationNameFormatName),
		string(OperationNameFormatUrl),
	}
}

func parseOperationNameFormat(input string) (*OperationNameFormat, error) {
	vals := map[string]OperationNameFormat{
		"name": OperationNameFormatName,
		"url":  OperationNameFormatUrl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationNameFormat(input)
	return &out, nil
}

type SamplingType string

const (
	SamplingTypeFixed SamplingType = "fixed"
)

func PossibleValuesForSamplingType() []string {
	return []string{
		string(SamplingTypeFixed),
	}
}

func parseSamplingType(input string) (*SamplingType, error) {
	vals := map[string]SamplingType{
		"fixed": SamplingTypeFixed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SamplingType(input)
	return &out, nil
}

type Verbosity string

const (
	VerbosityError       Verbosity = "error"
	VerbosityInformation Verbosity = "information"
	VerbosityVerbose     Verbosity = "verbose"
)

func PossibleValuesForVerbosity() []string {
	return []string{
		string(VerbosityError),
		string(VerbosityInformation),
		string(VerbosityVerbose),
	}
}

func parseVerbosity(input string) (*Verbosity, error) {
	vals := map[string]Verbosity{
		"error":       VerbosityError,
		"information": VerbosityInformation,
		"verbose":     VerbosityVerbose,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Verbosity(input)
	return &out, nil
}
