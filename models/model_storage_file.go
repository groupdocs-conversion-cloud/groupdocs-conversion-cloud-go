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

// File or folder information
type StorageFile struct {
	// File or folder name.
	Name string `json:"Name,omitempty"`
	// True if it is a folder.
	IsFolder bool `json:"IsFolder"`
	// File or folder last modified DateTime.
	ModifiedDate time.Time `json:"ModifiedDate,omitempty"`
	// File or folder size.
	Size int64 `json:"Size"`
	// File or folder path.
	Path string `json:"Path,omitempty"`
}
