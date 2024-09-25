/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

// Class for disc space information.
type DiscUsage struct {
	// Application used disc space.
	UsedSize int64 `json:"UsedSize"`
	// Total disc space.
	TotalSize int64 `json:"TotalSize"`
}
