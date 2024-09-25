/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

// Jpg convert options
type JpgConvertOptions struct {
	// Start conversion from FromPage page
	FromPage int32 `json:"FromPage"`
	// Number of pages to convert
	PagesCount int32 `json:"PagesCount"`
	// Convert specific pages. The list contains the page indexes of the pages to be converted
	Pages []int32 `json:"Pages,omitempty"`
	// Watermark specific options
	WatermarkOptions *WatermarkOptions `json:"WatermarkOptions,omitempty"`
	// Desired image width after conversion
	Width int32 `json:"Width"`
	// Desired image height after conversion
	Height int32 `json:"Height"`
	// Desired image horizontal resolution after conversion. The default resolution is the resolution of the input file or 96dpi
	HorizontalResolution int32 `json:"HorizontalResolution"`
	// Desired image vertical resolution after conversion. The default resolution is the resolution of the input file or 96dpi
	VerticalResolution int32 `json:"VerticalResolution"`
	// Convert to grayscale image
	Grayscale bool `json:"Grayscale"`
	// Image rotation angle
	RotateAngle int32 `json:"RotateAngle"`
	// If true, the input firstly is converted to PDF and after that to desired format
	UsePdf bool `json:"UsePdf"`
	// Adjust image brightness
	Brightness int32 `json:"Brightness"`
	// Adjust image contrast
	Contrast int32 `json:"Contrast"`
	// Adjust image gamma
	Gamma    float64      `json:"Gamma"`
	FlipMode FlipModeEnum `json:"FlipMode"`
	// Desired image quality when converting to Jpeg. The value must be between 0 and 100. The default value is 100.
	Quality int32 `json:"Quality"`
}
