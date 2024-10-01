/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

// Tiff convert options
type TiffConvertOptions struct {
	// Start conversion from FromPage page
	FromPage int32 `json:"FromPage,omitempty"`
	// Number of pages to convert
	PagesCount int32 `json:"PagesCount,omitempty"`
	// Convert specific pages. The list contains the page indexes of the pages to be converted
	Pages []int32 `json:"Pages,omitempty"`
	// Watermark specific options
	WatermarkOptions *WatermarkOptions `json:"WatermarkOptions,omitempty"`
	// Desired image width after conversion
	Width int32 `json:"Width,omitempty"`
	// Desired image height after conversion
	Height int32 `json:"Height,omitempty"`
	// Desired image horizontal resolution after conversion. The default resolution is the resolution of the input file or 96dpi
	HorizontalResolution int32 `json:"HorizontalResolution,omitempty"`
	// Desired image vertical resolution after conversion. The default resolution is the resolution of the input file or 96dpi
	VerticalResolution int32 `json:"VerticalResolution,omitempty"`
	// Convert to grayscale image
	Grayscale bool `json:"Grayscale,omitempty"`
	// Image rotation angle
	RotateAngle int32 `json:"RotateAngle,omitempty"`
	// If true, the input firstly is converted to PDF and after that to desired format
	UsePdf bool `json:"UsePdf,omitempty"`
	// Adjust image brightness
	Brightness int32 `json:"Brightness,omitempty"`
	// Adjust image contrast
	Contrast int32 `json:"Contrast,omitempty"`
	// Adjust image gamma
	Gamma       float64         `json:"Gamma,omitempty"`
	FlipMode    FlipModeEnum    `json:"FlipMode,omitempty"`
	Compression CompressionEnum `json:"Compression,omitempty"`
}
