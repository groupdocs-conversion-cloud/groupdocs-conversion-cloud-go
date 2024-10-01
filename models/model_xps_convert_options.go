/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

// Xps convert options
type XpsConvertOptions struct {
	// Desired page width in pixels after conversion
	Width int32 `json:"Width,omitempty"`
	// Desired page height in pixels after conversion
	Height int32 `json:"Height,omitempty"`
	// Desired page DPI after conversion. The default resolution is: 96dpi
	Dpi int32 `json:"Dpi,omitempty"`
	// Desired page top margin in pixels after conversion
	MarginTop int32 `json:"MarginTop,omitempty"`
	// Desired page bottom margin in pixels after conversion
	MarginBottom int32 `json:"MarginBottom,omitempty"`
	// Desired page left margin in pixels after conversion
	MarginLeft int32 `json:"MarginLeft,omitempty"`
	// Desired page right margin in pixels after conversion
	MarginRight int32 `json:"MarginRight,omitempty"`
	// Watermark specific options
	WatermarkOptions *WatermarkOptions `json:"WatermarkOptions,omitempty"`
	// Set this property if you want to protect the converted document with a password
	Password string `json:"Password,omitempty"`
	// If true, the input firstly is converted to PDF and after that to desired format
	UsePdf bool `json:"UsePdf,omitempty"`
}
