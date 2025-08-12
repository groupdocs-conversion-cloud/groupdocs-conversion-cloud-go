/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 */

package models

// Pdf document load options
type PdfLoadOptions struct {
	// The format of input file, (\"docx\", for example). This field must be filled with correct input file format when using ConvertDirect method, which accept input file as binary stream, and, because of that, API can correctly handle LoadOptions. In regular conversion, the input file format taken from the input file name and this field ignored.
	Format string `json:"Format,omitempty"`
	// Remove embedded files
	RemoveEmbeddedFiles bool `json:"RemoveEmbeddedFiles,omitempty"`
	// Set password to unprotect protected document
	Password string `json:"Password,omitempty"`
	// Hide annotations in Pdf documents
	HidePdfAnnotations bool `json:"HidePdfAnnotations,omitempty"`
	// Flatten all the fields of the PDF form
	FlattenAllFields bool `json:"FlattenAllFields,omitempty"`
	// Clear built-in document properties
	ClearBuiltInDocumentProperties bool `json:"ClearBuiltInDocumentProperties"`
	// Clear custom document properties
	ClearCustomDocumentProperties bool `json:"ClearCustomDocumentProperties"`
	// Enable or disable generation of page numbering in converted document. Default:     false
	PageNumbering bool `json:"PageNumbering"`
	// Default font for Pdf document. The following font will be used if a font is missing.
	DefaultFont string `json:"DefaultFont,omitempty"`
	// Remove javascript
	RemoveJavascript bool `json:"RemoveJavascript"`
	// Substitute specific fonts when converting Words document.
	FontSubstitutes map[string]string `json:"FontSubstitutes,omitempty"`
}
