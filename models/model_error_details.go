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

// The error details
type ErrorDetails struct {
	// The request id
	RequestId string `json:"RequestId,omitempty"`
	// Date
	Date time.Time `json:"Date"`
}
