/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 */

package models

// Spreadsheet document load options
type SpreadsheetLoadOptions struct {
	// The format of input file, (\"docx\", for example). This field must be filled with correct input file format when using ConvertDirect method, which accept input file as binary stream, and, because of that, API can correctly handle LoadOptions. In regular conversion, the input file format taken from the input file name and this field ignored.
	Format string `json:"Format,omitempty"`
	// Default font for Cells document. The following font will be used if a font is missing.
	DefaultFont string `json:"DefaultFont,omitempty"`
	// Substitute specific fonts when converting Cells document.
	FontSubstitutes map[string]string `json:"FontSubstitutes,omitempty"`
	// Show grid lines when converting Excel files
	ShowGridLines bool `json:"ShowGridLines"`
	// Show hidden sheets when converting Excel files
	ShowHiddenSheets bool `json:"ShowHiddenSheets"`
	// If OnePagePerSheet is true the content of the sheet will be converted to one page in the PDF document. Default value is true.
	OnePagePerSheet bool `json:"OnePagePerSheet"`
	// Convert specific range when converting to other than cells format. Example: \"D1:F8\"
	ConvertRange string `json:"ConvertRange,omitempty"`
	// Skips empty rows and columns when converting. Default is True.
	SkipEmptyRowsAndColumns bool `json:"SkipEmptyRowsAndColumns"`
	// Set password to unprotect protected document
	Password string `json:"Password,omitempty"`
	// Represents the way comments are printed with the sheet. Default is PrintNoComments.
	PrintComments PrintCommentsEnum `json:"PrintComments,omitempty"`
	// Clear custom document properties. Default is false.
	ClearCustomDocumentProperties bool `json:"ClearCustomDocumentProperties"`
	// Clear built-in document properties. Default is false.
	ClearBuiltInDocumentProperties bool `json:"ClearBuiltInDocumentProperties"`
	// Split a worksheet into pages by rows. Default is 0, no pagination.
	RowsPerPage int32 `json:"RowsPerPage"`
	// Split a worksheet into pages by columns. Default is 0, no pagination.
	ColumnsPerPage int32 `json:"ColumnsPerPage"`
	// Autofits all rows when converting
	AutoFitRows bool `json:"AutoFitRows"`
	// If AllColumnsInOnePagePerSheet is true, all column content of one sheet will output to only one page in result. The width of paper size of pagesetup will be invalid, and the other settings of pagesetup will still take effect.
	AllColumnsInOnePagePerSheet bool `json:"AllColumnsInOnePagePerSheet"`
	// System culture info at the time file is loaded
	CultureInfo string `json:"CultureInfo,omitempty"`
	// Whether check restriction of excel file when user modify cells related objects. For example, excel does not allow inputting string value longer than 32K. When you input a value longer than 32K, if this property is true, you will get an Exception. If this property is false, we will accept your input string value as the cell's value so that later you can output the complete string value for other file formats such as CSV. However, if you have set such kind of value that is invalid for excel file format, you should not save the workbook as excel file format later. Otherwise there may be unexpected error for the generated excel file.
	CheckExcelRestriction bool `json:"CheckExcelRestriction"`
	// If True and converting to Pdf the conversion is optimized for better file size than print quality.
	OptimizePdfSize bool `json:"OptimizePdfSize"`
	// List of sheet indexes to convert. The indexes must be zero-based
	SheetIndexes []int32 `json:"SheetIndexes,omitempty"`
	// List of sheet names to convert
	Sheets []string `json:"Sheets,omitempty"`
	// Reset font folders before loading document
	ResetFontFolders bool `json:"ResetFontFolders"`
}
