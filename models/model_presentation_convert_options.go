/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

// Options for to presentation conversion
type PresentationConvertOptions struct {
	// Start conversion from FromPage page
	FromPage int32 `json:"FromPage"`
	// Number of pages to convert
	PagesCount int32 `json:"PagesCount"`
	// Convert specific pages. The list contains the page indexes of the pages to be converted
	Pages []int32 `json:"Pages,omitempty"`
	// Watermark specific options
	WatermarkOptions *WatermarkOptions `json:"WatermarkOptions,omitempty"`
	// Set this property if you want to protect the converted document with a password
	Password string `json:"Password,omitempty"`
	// Specifies the zoom level in percentage. Default is 100. Default zoom is supported till Microsoft Powerpoint 2010. Starting from Microsoft Powerpoint 2013 default zoom is no longer set to document, instead it appears to use the zoom factor of the last document that was opened.
	Zoom int32 `json:"Zoom"`
}
