/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 */

package models

// Options for loading personal storage documents.
type PersonalStorageLoadOptions struct {
	// The format of input file, (\"docx\", for example). This field must be filled with correct input file format when using ConvertDirect method, which accept input file as binary stream, and, because of that, API can correctly handle LoadOptions. In regular conversion, the input file format taken from the input file name and this field ignored.
	Format string `json:"Format,omitempty"`
	// Folder which to be processed Default is Inbox
	Folder string `json:"Folder,omitempty"`
	// Controls how many levels in depth to perform conversion
	Depth int32 `json:"Depth,omitempty"`
}
