/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

// Ebook convert options
type EBookConvertOptions struct {
	// Start conversion from FromPage page
	FromPage int32 `json:"FromPage"`
	// Number of pages to convert
	PagesCount int32 `json:"PagesCount"`
	// Convert specific pages. The list contains the page indexes of the pages to be converted
	Pages []int32 `json:"Pages,omitempty"`
	// Watermark specific options
	WatermarkOptions *WatermarkOptions   `json:"WatermarkOptions,omitempty"`
	PageSize         PageSizeEnum        `json:"PageSize"`
	PageOrientation  PageOrientationEnum `json:"PageOrientation"`
}
