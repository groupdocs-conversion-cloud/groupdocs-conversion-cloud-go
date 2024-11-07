/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 */

package models

// Defines conversion request
type ConvertSettings struct {
	// StorageName which contains the file
	StorageName string `json:"StorageName,omitempty"`
	// Gets or sets absolute path to a file in the storage
	FilePath string `json:"FilePath"`
	// Gets or sets requested conversion format
	Format string `json:"Format"`
	// Gets or sets format specific load options for source file
	LoadOptions interface{} `json:"LoadOptions,omitempty"`
	// Gets or sets format specific convert options for output file
	ConvertOptions interface{} `json:"ConvertOptions,omitempty"`
	// Gets or sets converted file save path
	OutputPath string `json:"OutputPath,omitempty"`
	// The path to directory containing custom fonts in storage
	FontsPath string `json:"FontsPath,omitempty"`
}
