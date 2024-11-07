/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 */

package models

// Txt document load options
type TxtLoadOptions struct {
	// The format of input file, (\"docx\", for example). This field must be filled with correct input file format when using ConvertDirect method, which accept input file as binary stream, and, because of that, API can correctly handle LoadOptions. In regular conversion, the input file format taken from the input file name and this field ignored.
	Format string `json:"Format,omitempty"`
	// Allows to specify how numbered list items are recognized when plain text document is converted. The default value is true.
	DetectNumberingWithWhitespaces bool                      `json:"DetectNumberingWithWhitespaces,omitempty"`
	TrailingSpacesOptions          TrailingSpacesOptionsEnum `json:"TrailingSpacesOptions,omitempty"`
	LeadingSpacesOptions           LeadingSpacesOptionsEnum  `json:"LeadingSpacesOptions,omitempty"`
	// Gets or sets the encoding that will be used when loading Txt document. Can be null. Default is null.
	Encoding string `json:"Encoding,omitempty"`
}
