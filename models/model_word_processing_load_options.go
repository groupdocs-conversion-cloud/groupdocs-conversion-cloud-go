/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

// WordProcessing document load options
type WordProcessingLoadOptions struct {
	// The format of input file, (\"docx\", for example). This field must be filled with correct input file format when using ConvertDirect method, which accept input file as binary stream, and, because of that, API can correctly handle LoadOptions. In regular conversion, the input file format taken from the input file name and this field ignored.
	Format string `json:"Format,omitempty"`
	// Default font for Words document. The following font will be used if a font is missing.
	DefaultFont string `json:"DefaultFont,omitempty"`
	// If AutoFontSubstitution is disabled, GroupDocs.Conversion uses the DefaultFont for the substitution of missing fonts. If AutoFontSubstitution is enabled, GroupDocs.Conversion evaluates all the related fields in FontInfo (Panose, Sig etc) for the missing font and finds the closest match among the available font sources. Note that font substitution mechanism will override the DefaultFont in cases when FontInfo for the missing font is available in the document. The default value is True.
	AutoFontSubstitution bool `json:"AutoFontSubstitution"`
	// Substitute specific fonts when converting Words document.
	FontSubstitutes map[string]string `json:"FontSubstitutes,omitempty"`
	// Set password to unprotect protected document
	Password string `json:"Password,omitempty"`
	// Hide markup and track changes for Word documents
	HideWordTrackedChanges bool `json:"HideWordTrackedChanges"`
	// Hide comments
	HideComments bool `json:"HideComments"`
	// Specifies the default level in the document outline at which to display Word bookmarks. Default is 0. Valid range is 0 to 9.
	BookmarksOutlineLevel int32 `json:"BookmarksOutlineLevel,omitempty"`
	// Specifies how many levels of headings (paragraphs formatted with the Heading styles) to include in the document outline. Default is 0. Valid range is 0 to 9.
	HeadingsOutlineLevels int32 `json:"HeadingsOutlineLevels,omitempty"`
	// Specifies how many levels in the document outline to show expanded when the file is viewed. Default is 0. Valid range is 0 to 9. Note that this options will not work when saving to XPS.
	ExpandedOutlineLevels int32 `json:"ExpandedOutlineLevels,omitempty"`
}
