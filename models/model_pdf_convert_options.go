/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

// Options for to PDF conversion
type PdfConvertOptions struct {
	// Start conversion from FromPage page
	FromPage int32 `json:"FromPage"`
	// Number of pages to convert
	PagesCount int32 `json:"PagesCount"`
	// Convert specific pages. The list contains the page indexes of the pages to be converted
	Pages []int32 `json:"Pages,omitempty"`
	// Watermark specific options
	WatermarkOptions *WatermarkOptions `json:"WatermarkOptions,omitempty"`
	// Desired page width in pixels after conversion
	Width int32 `json:"Width"`
	// Desired page height in pixels after conversion
	Height int32 `json:"Height"`
	// Desired page DPI after conversion. The default resolution is: 96dpi
	Dpi int32 `json:"Dpi"`
	// Set this property if you want to protect the converted document with a password
	Password string `json:"Password,omitempty"`
	// Desired page top margin in pixels after conversion
	MarginTop int32 `json:"MarginTop"`
	// Desired page bottom margin in pixels after conversion
	MarginBottom int32 `json:"MarginBottom"`
	// Desired page left margin in pixels after conversion
	MarginLeft int32 `json:"MarginLeft"`
	// Desired page right margin in pixels after conversion
	MarginRight int32         `json:"MarginRight"`
	PdfFormat   PdfFormatEnum `json:"PdfFormat"`
	// Remove Pdf-A Compliance
	RemovePdfaCompliance bool `json:"RemovePdfaCompliance"`
	// Specifies the zoom level in percentage. Default is 100.
	Zoom int32 `json:"Zoom"`
	// Linearize PDF Document for the Web
	Linearize bool `json:"Linearize"`
	// Link duplicate streams
	LinkDuplicateStreams bool `json:"LinkDuplicateStreams"`
	// Remove unused objects
	RemoveUnusedObjects bool `json:"RemoveUnusedObjects"`
	// Remove unused streams
	RemoveUnusedStreams bool `json:"RemoveUnusedStreams"`
	// If CompressImages set to true, all images in the document are recompressed. The compression is defined by the ImageQuality property.
	CompressImages bool `json:"CompressImages"`
	// Value in percent where 100% is unchanged quality and image size. To decrease the image size, use ImageQuality less than 100
	ImageQuality int32 `json:"ImageQuality"`
	// Make fonts not embedded if set to true
	UnembedFonts bool `json:"UnembedFonts"`
	// Convert a PDF from RGB colorspace to Grayscale
	Grayscale bool `json:"Grayscale"`
	// Specify whether position of the document's window will be centered on the screen. Default: false.
	CenterWindow bool          `json:"CenterWindow"`
	Direction    DirectionEnum `json:"Direction"`
	// Specifying whether document's window title bar should display document title. Default: false.
	DisplayDocTitle bool `json:"DisplayDocTitle"`
	// Specify whether document window must be resized to fit the first displayed page. Default: false.
	FitWindow bool `json:"FitWindow"`
	// Specify whether menu bar should be hidden when document is active. Default: false.
	HideMenuBar bool `json:"HideMenuBar"`
	// Specifying whether toolbar should be hidden when document is active. Default: false.
	HideToolBar bool `json:"HideToolBar"`
	// Specify whether user interface elements should be hidden when document is active. Default: false.
	HideWindowUI          bool                      `json:"HideWindowUI"`
	NonFullScreenPageMode NonFullScreenPageModeEnum `json:"NonFullScreenPageMode"`
	PageLayout            PageLayoutEnum            `json:"PageLayout"`
	PageMode              PageModeEnum              `json:"PageMode"`
	Rotate                RotateEnum                `json:"Rotate"`
	PageSize              PageSizeEnum              `json:"PageSize"`
	PageOrientation       PageOrientationEnum       `json:"PageOrientation"`
}
