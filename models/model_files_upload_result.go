/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

// File upload result
type FilesUploadResult struct {
	// List of uploaded file names
	Uploaded []string `json:"Uploaded,omitempty"`
	// List of errors.
	Errors []ModelError `json:"Errors,omitempty"`
}
