/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

// Options for to word processing conversion
type WordProcessingConvertOptions struct {
	// Start conversion from FromPage page
	FromPage int32 `json:"FromPage"`
	// Number of pages to convert
	PagesCount int32 `json:"PagesCount"`
	// Convert specific pages. The list contains the page indexes of the pages to be converted
	Pages []int32 `json:"Pages,omitempty"`
	// Watermark specific options
	WatermarkOptions *WatermarkOptions `json:"WatermarkOptions,omitempty"`
	// Desired page width after conversion
	Width int32 `json:"Width"`
	// Desired page height after conversion
	Height int32 `json:"Height"`
	// Desired page DPI after conversion. The default resolution is: 96dpi
	Dpi int32 `json:"Dpi"`
	// Set this property if you want to protect the converted document with a password
	Password string `json:"Password,omitempty"`
	// Specifies the zoom level in percentage. Default is 100. Default zoom is supported till Microsoft Word 2010. Starting from Microsoft Word 2013 default zoom is no longer set to document, instead it appears to use the zoom factor of the last document that was opened.
	Zoom               int32                  `json:"Zoom"`
	PdfRecognitionMode PdfRecognitionModeEnum `json:"PdfRecognitionMode"`
	PageSize           PageSizeEnum           `json:"PageSize"`
	PageOrientation    PageOrientationEnum    `json:"PageOrientation"`
}
