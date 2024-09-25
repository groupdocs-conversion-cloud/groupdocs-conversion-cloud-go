/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

// Object exists
type ObjectExist struct {
	// Indicates that the file or folder exists.
	Exists bool `json:"Exists"`
	// True if it is a folder, false if it is a file.
	IsFolder bool `json:"IsFolder"`
}
