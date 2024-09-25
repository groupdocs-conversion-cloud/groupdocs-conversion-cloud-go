/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

// Contains single converted item. Result is provided as url to a storage
type StoredConvertedResult struct {
	// Name of converted item
	Name string `json:"Name,omitempty"`
	// Size of converted item
	Size int64 `json:"Size"`
	// Path of resource file in storage
	Path string `json:"Path,omitempty"`
	// Uri in the storage of the converted item
	Url string `json:"Url,omitempty"`
}
