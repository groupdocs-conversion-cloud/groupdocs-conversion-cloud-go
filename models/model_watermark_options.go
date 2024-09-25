/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

// Options for settings watermark to the converted document
type WatermarkOptions struct {
	// Watermark text
	Text string `json:"Text,omitempty"`
	// Watermark font name if text watermark is applied
	FontName string `json:"FontName,omitempty"`
	// Watermark font name if text watermark is applied
	FontSize int32 `json:"FontSize"`
	// Watermark font bold style if text watermark is applied
	Bold bool `json:"Bold"`
	// Watermark font italic style if text watermark is applied
	Italic bool `json:"Italic"`
	// Watermark font color if text watermark is applied
	Color string `json:"Color,omitempty"`
	// Watermark width
	Width int32 `json:"Width"`
	// Watermark height
	Height int32 `json:"Height"`
	// Watermark top position
	Top int32 `json:"Top"`
	// Watermark left position
	Left int32 `json:"Left"`
	// Watermark rotation angle
	RotationAngle int32 `json:"RotationAngle"`
	// Watermark transparency. Value between 0 and 1. Value 0 is fully visible, value 1 is invisible.
	Transparency float64 `json:"Transparency"`
	// Indicates that the watermark is stamped as background. If the value is true, the watermark is layed at the bottom. By default is false and the watermark is layed on top.
	Background bool `json:"Background"`
	// Image watermark
	Image string `json:"Image,omitempty"`
	// Auto scale the watermark. If the value is true the font size and the position is automatically calculated to fit the page size.
	AutoAlign bool `json:"AutoAlign"`
}
