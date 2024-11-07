/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 */

package models

// Html document load options
type WebLoadOptions struct {
	// The format of input file, (\"docx\", for example). This field must be filled with correct input file format when using ConvertDirect method, which accept input file as binary stream, and, because of that, API can correctly handle LoadOptions. In regular conversion, the input file format taken from the input file name and this field ignored.
	Format string `json:"Format,omitempty"`
	// Enable or disable generation of page numbering in converted document. Default: false
	PageNumbering bool `json:"PageNumbering,omitempty"`
	// Use pdf for the conversion. Default: false
	UsePdf bool `json:"UsePdf,omitempty"`
	// Controls how HTML content is rendered. Default: AbsolutePositioning
	RenderingMode HtmlRenderingModeEnum `json:"RenderingMode,omitempty"`
}
