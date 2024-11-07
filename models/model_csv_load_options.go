/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 */

package models

// Csv document load options
type CsvLoadOptions struct {
	// The format of input file, (\"docx\", for example). This field must be filled with correct input file format when using ConvertDirect method, which accept input file as binary stream, and, because of that, API can correctly handle LoadOptions. In regular conversion, the input file format taken from the input file name and this field ignored.
	Format string `json:"Format,omitempty"`
	// Default font for Cells document. The following font will be used if a font is missing.
	DefaultFont string `json:"DefaultFont,omitempty"`
	// Substitute specific fonts when converting Cells document.
	FontSubstitutes map[string]string `json:"FontSubstitutes,omitempty"`
	// Show grid lines when converting Excel files
	ShowGridLines bool `json:"ShowGridLines,omitempty"`
	// Show hidden sheets when converting Excel files
	ShowHiddenSheets bool `json:"ShowHiddenSheets,omitempty"`
	// If OnePagePerSheet is true the content of the sheet will be converted to one page in the PDF document. Default value is true.
	OnePagePerSheet bool `json:"OnePagePerSheet,omitempty"`
	// Convert specific range when converting to other than cells format. Example: \"D1:F8\"
	ConvertRange string `json:"ConvertRange,omitempty"`
	// Skips empty rows and columns when converting. Default is True.
	SkipEmptyRowsAndColumns bool `json:"SkipEmptyRowsAndColumns,omitempty"`
	// Set password to unprotect protected document
	Password string `json:"Password,omitempty"`
	// Hide comments
	HideComments bool `json:"HideComments,omitempty"`
	// Delimiter of a Csv file
	Separator string `json:"Separator,omitempty"`
	// True means the file contains several encodings
	IsMultiEncoded bool `json:"IsMultiEncoded,omitempty"`
	// Indicates whether text is formula if it starts with \"=\"
	HasFormula bool `json:"HasFormula,omitempty"`
	// Indicates whether the string in the file is converted to numeric. Default is True
	ConvertNumericData bool `json:"ConvertNumericData,omitempty"`
	// Indicates whether the string in the file is converted to date. Default is True
	ConvertDateTimeData bool `json:"ConvertDateTimeData,omitempty"`
	// File encoding
	Encoding string `json:"Encoding,omitempty"`
}
