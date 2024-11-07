/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 */

package models

// Options for to PDF conversion
type PdfConvertOptions struct {
	// Start conversion from FromPage page
	FromPage int32 `json:"FromPage,omitempty"`
	// Number of pages to convert
	PagesCount int32 `json:"PagesCount,omitempty"`
	// Convert specific pages. The list contains the page indexes of the pages to be converted
	Pages []int32 `json:"Pages,omitempty"`
	// Watermark specific options
	WatermarkOptions *WatermarkOptions `json:"WatermarkOptions,omitempty"`
	// Desired page width in pixels after conversion
	Width int32 `json:"Width,omitempty"`
	// Desired page height in pixels after conversion
	Height int32 `json:"Height,omitempty"`
	// Desired page DPI after conversion. The default resolution is: 96dpi
	Dpi int32 `json:"Dpi,omitempty"`
	// Set this property if you want to protect the converted document with a password
	Password string `json:"Password,omitempty"`
	// Desired page top margin in pixels after conversion
	MarginTop int32 `json:"MarginTop,omitempty"`
	// Desired page bottom margin in pixels after conversion
	MarginBottom int32 `json:"MarginBottom,omitempty"`
	// Desired page left margin in pixels after conversion
	MarginLeft int32 `json:"MarginLeft,omitempty"`
	// Desired page right margin in pixels after conversion
	MarginRight int32         `json:"MarginRight,omitempty"`
	PdfFormat   PdfFormatEnum `json:"PdfFormat,omitempty"`
	// Remove Pdf-A Compliance
	RemovePdfaCompliance bool `json:"RemovePdfaCompliance,omitempty"`
	// Specifies the zoom level in percentage. Default is 100.
	Zoom int32 `json:"Zoom,omitempty"`
	// Linearize PDF Document for the Web
	Linearize bool `json:"Linearize,omitempty"`
	// Link duplicate streams
	LinkDuplicateStreams bool `json:"LinkDuplicateStreams,omitempty"`
	// Remove unused objects
	RemoveUnusedObjects bool `json:"RemoveUnusedObjects,omitempty"`
	// Remove unused streams
	RemoveUnusedStreams bool `json:"RemoveUnusedStreams,omitempty"`
	// If CompressImages set to true, all images in the document are recompressed. The compression is defined by the ImageQuality property.
	CompressImages bool `json:"CompressImages,omitempty"`
	// Value in percent where 100% is unchanged quality and image size. To decrease the image size, use ImageQuality less than 100
	ImageQuality int32 `json:"ImageQuality,omitempty"`
	// Make fonts not embedded if set to true
	UnembedFonts bool `json:"UnembedFonts,omitempty"`
	// Convert a PDF from RGB colorspace to Grayscale
	Grayscale bool `json:"Grayscale,omitempty"`
	// Specify whether position of the document's window will be centered on the screen. Default: false.
	CenterWindow bool          `json:"CenterWindow,omitempty"`
	Direction    DirectionEnum `json:"Direction,omitempty"`
	// Specifying whether document's window title bar should display document title. Default: false.
	DisplayDocTitle bool `json:"DisplayDocTitle,omitempty"`
	// Specify whether document window must be resized to fit the first displayed page. Default: false.
	FitWindow bool `json:"FitWindow,omitempty"`
	// Specify whether menu bar should be hidden when document is active. Default: false.
	HideMenuBar bool `json:"HideMenuBar,omitempty"`
	// Specifying whether toolbar should be hidden when document is active. Default: false.
	HideToolBar bool `json:"HideToolBar,omitempty"`
	// Specify whether user interface elements should be hidden when document is active. Default: false.
	HideWindowUI          bool                      `json:"HideWindowUI,omitempty"`
	NonFullScreenPageMode NonFullScreenPageModeEnum `json:"NonFullScreenPageMode,omitempty"`
	PageLayout            PageLayoutEnum            `json:"PageLayout,omitempty"`
	PageMode              PageModeEnum              `json:"PageMode,omitempty"`
	Rotate                RotateEnum                `json:"Rotate,omitempty"`
	PageSize              PageSizeEnum              `json:"PageSize,omitempty"`
	PageOrientation       PageOrientationEnum       `json:"PageOrientation,omitempty"`
}
