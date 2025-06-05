/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 */

package models

// WordProcessingLoadOptions defines options for loading WordProcessing documents.
type WordProcessingLoadOptions struct {
	// The format of input file, (\"docx\", for example). This field must be filled with correct input file format when using ConvertDirect method, which accept input file as binary stream, and, because of that, API can correctly handle LoadOptions. In regular conversion, the input file format taken from the input file name and this field ignored.
	Format string `json:"Format,omitempty"`
	// Default font for Words document. The following font will be used if a font is missing.
	DefaultFont string `json:"DefaultFont,omitempty"`
	// Substitute specific fonts when converting Words document.
	FontSubstitutes map[string]string `json:"FontSubstitutes,omitempty"`
	// Set password to unprotect protected document
	Password string `json:"Password,omitempty"`
	// Hide markup and track changes for Word documents
	HideWordTrackedChanges bool `json:"HideWordTrackedChanges,omitempty"`
	// Specifies the default level in the document outline at which to display Word bookmarks. Default is 0. Valid range is 0 to 9.
	BookmarksOutlineLevel int32 `json:"BookmarksOutlineLevel,omitempty"`
	// Specifies how many levels of headings (paragraphs formatted with the Heading styles) to include in the document outline. Default is 0. Valid range is 0 to 9.
	HeadingsOutlineLevels int32 `json:"HeadingsOutlineLevels,omitempty"`
	// Specifies how many levels in the document outline to show expanded when the file is viewed. Default is 0. Valid range is 0 to 9. Note that this options will not work when saving to XPS.
	ExpandedOutlineLevels int32 `json:"ExpandedOutlineLevels,omitempty"`
	// Clear custom document properties. Default is false.
	ClearCustomDocumentProperties bool `json:"ClearCustomDocumentProperties,omitempty"`
	// Clear built-in document properties. Default is false.
	ClearBuiltInDocumentProperties bool `json:"ClearBuiltInDocumentProperties,omitempty"`
	// Option to control how many levels in depth to perform conversion. Default: 1.
	Depth int32 `json:"Depth,omitempty"`
	// Option to control whether the owned documents in the documents container must be converted
	ConvertOwned bool `json:"ConvertOwned,omitempty"`
	// Option to control whether the documents container itself must be converted. If true, the documents container will be the first converted document. Default is true.
	ConvertOwner bool `json:"ConvertOwner,omitempty"`
	// Gets or sets value determining whether automatic hyphenation is turned on for the document. Default value is false.
	AutoHyphenation bool `json:"AutoHyphenation,omitempty"`
	// Gets or sets value determining whether words written in all capital letters are hyphenated. Default value is true.
	HyphenateCaps bool `json:"HyphenateCaps,omitempty"`
	// Enable or disable generation of page numbering in converted document. Default: false
	PageNumbering bool `json:"PageNumbering,omitempty"`
	// Determines whether the document structure should be preserved when converting to PDF (default is false).
	PreserveDocumentStructure bool `json:"PreserveDocumentStructure,omitempty"`
	// If true all external resource will not be loading. Default is true.
	SkipExternalResources bool `json:"SkipExternalResources,omitempty"`
	// Specifies whether to use a text shaper for better kerning display. Default is false.
	UseTextShaper bool `json:"UseTextShaper,omitempty"`
	// Specifies whether to preserve Microsoft Word form fields as form fields in PDF or convert them to text. Default is false.
	PreserveFormFields bool `json:"PreserveFormFields,omitempty"`
	// Specifies how comments should be displayed in the output document. Default is Balloon.
	CommentDisplayMode CommentDisplayModeEnum `json:"CommentDisplayMode,omitempty"`
	// Keep original value of date field. Default: false
	KeepDateFieldOriginalValue bool `json:"KeepDateFieldOriginalValue,omitempty"`
	// Update fields after loading. Default: false
	UpdateFields bool `json:"UpdateFields,omitempty"`
	// Update page layout after loading. Default: false
	UpdatePageLayout bool `json:"UpdatePageLayout,omitempty"`
	// If EmbedTrueTypeFonts is true, embed true type fonts in the output document. Default: true
	EmbedTrueTypeFonts bool `json:"EmbedTrueTypeFonts,omitempty"`
	// Automatically substitutes missing fonts based on FontInfo in the document. Default: false.
	FontInfoSubstitutionEnabled bool `json:"FontInfoSubstitutionEnabled,omitempty"`
	// Automatically substitutes missing fonts based on FontConfig in the system. Default: false.
	FontConfigSubstitutionEnabled bool `json:"FontConfigSubstitutionEnabled,omitempty"`
	// Automatically substitutes missing fonts based on the font name. Default: false.
	FontNameSubstitutionEnabled bool `json:"FontNameSubstitutionEnabled,omitempty"`
	// Show full commenter name in comments. Default is false.
	ShowFullCommenterName bool `json:"ShowFullCommenterName,omitempty"`
}
