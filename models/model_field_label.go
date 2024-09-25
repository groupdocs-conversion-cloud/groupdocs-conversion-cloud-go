/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

// Represents field label
type FieldLabel struct {
	Field FieldEnum `json:"Field"`
	// The label e.g. \"Sender\"
	Label string `json:"Label,omitempty"`
}
