/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 */

package models

// Represents information about supported conversion for SourceFormat
type SupportedFormat struct {
	// Gets or sets source format
	SourceFormat string `json:"SourceFormat,omitempty"`
	// Gets or sets target formats
	TargetFormats []string `json:"TargetFormats,omitempty"`
}
