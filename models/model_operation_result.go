/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

import (
	"time"
)

type MethodEnum string

const (
	MethodEnumConvert        MethodEnum = "Convert"
	MethodEnumConvertAndSave MethodEnum = "ConvertAndSave"
)

type StatusEnum string

const (
	StatusEnumCreated  StatusEnum = "Created"
	StatusEnumStarted  StatusEnum = "Started"
	StatusEnumFailed   StatusEnum = "Failed"
	StatusEnumCanceled StatusEnum = "Canceled"
	StatusEnumFinished StatusEnum = "Finished"
)

// Operation status result
type OperationResult struct {
	Id       string                  `json:"Id"`
	Method   MethodEnum              `json:"Method"`
	Status   StatusEnum              `json:"Status"`
	Created  time.Time               `json:"Created,omitempty"`
	Started  time.Time               `json:"Started,omitempty"`
	Failed   time.Time               `json:"Failed,omitempty"`
	Canceled time.Time               `json:"Canceled,omitempty"`
	Finished time.Time               `json:"Finished,omitempty"`
	Result   []StoredConvertedResult `json:"Result,omitempty"`
	Error_   string                  `json:"Error,omitempty"`
}
