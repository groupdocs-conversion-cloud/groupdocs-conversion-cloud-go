/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 */

package models

// Spreadsheet сonvert options class
type SpreadsheetConvertOptions struct {
	// Start conversion from FromPage page
	FromPage int32 `json:"FromPage,omitempty"`
	// Number of pages to convert
	PagesCount int32 `json:"PagesCount,omitempty"`
	// Convert specific pages. The list contains the page indexes of the pages to be converted
	Pages []int32 `json:"Pages,omitempty"`
	// Watermark specific options
	WatermarkOptions *WatermarkOptions `json:"WatermarkOptions,omitempty"`
	// Set this property if you want to protect the converted document with a password
	Password string `json:"Password,omitempty"`
	// Specifies the zoom level in percentage. Default is 100.
	Zoom int32 `json:"Zoom,omitempty"`
	// If true, the input firstly is converted to PDF and after that to desired format
	UsePdf bool `json:"UsePdf,omitempty"`
}
