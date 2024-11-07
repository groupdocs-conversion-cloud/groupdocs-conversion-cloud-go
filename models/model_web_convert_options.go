/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 */

package models

// Options for to Html conversion
type WebConvertOptions struct {
	// Start conversion from FromPage page
	FromPage int32 `json:"FromPage,omitempty"`
	// Number of pages to convert
	PagesCount int32 `json:"PagesCount,omitempty"`
	// Convert specific pages. The list contains the page indexes of the pages to be converted
	Pages []int32 `json:"Pages,omitempty"`
	// Watermark specific options
	WatermarkOptions *WatermarkOptions `json:"WatermarkOptions,omitempty"`
	// If true, the input firstly is converted to PDF and after that to desired format
	UsePdf bool `json:"UsePdf,omitempty"`
	// If true fixed layout will be used e.g. absolutely positioned html elements Default:  true
	FixedLayout bool `json:"FixedLayout,omitempty"`
	// Show page borders when converting to fixed layout. Default is True
	FixedLayoutShowBorders bool `json:"FixedLayoutShowBorders,omitempty"`
	// Specifies the zoom level in percentage. Default is 100.
	Zoom int32 `json:"Zoom,omitempty"`
}
