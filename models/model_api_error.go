/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard-qa.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package models

import (
	"time"
)

type ApiError struct {
	Code        string    `json:"Code,omitempty"`
	Message     string    `json:"Message,omitempty"`
	Description string    `json:"Description,omitempty"`
	DateTime    time.Time `json:"DateTime,omitempty"`
	InnerError  *ApiError `json:"InnerError,omitempty"`
}