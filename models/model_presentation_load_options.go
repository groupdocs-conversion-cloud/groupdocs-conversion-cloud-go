/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 */

package models

// Presentation document load options
type PresentationLoadOptions struct {
	// The format of input file, (\"docx\", for example). This field must be filled with correct input file format when using ConvertDirect method, which accept input file as binary stream, and, because of that, API can correctly handle LoadOptions. In regular conversion, the input file format taken from the input file name and this field ignored.
	Format string `json:"Format,omitempty"`
	// Default font for rendering the presentation. The following font will be used if a presentation font is missing.
	DefaultFont string `json:"DefaultFont,omitempty"`
	// Substitute specific fonts when converting Slides document.
	FontSubstitutes map[string]string `json:"FontSubstitutes,omitempty"`
	// Set password to unprotect protected document
	Password string `json:"Password,omitempty"`
	// Represents the way comments are printed with the slide. Default is None.
	CommentsPosition CommentsPositionEnum `json:"CommentsPosition,omitempty"`
	// Represents the way notes are printed with the slide. Default is None.
	NotesPosition NotesPositionEnum `json:"NotesPosition,omitempty"`
	// Show hidden slides
	ShowHiddenSlides bool `json:"ShowHiddenSlides"`
	// Determines whether the document structure should be preserved when converting to PDF (default is false).
	PreserveDocumentStructure bool `json:"PreserveDocumentStructure,omitempty"`
	// Value indicating whether custom document properties should be cleared.
	ClearCustomDocumentProperties bool `json:"ClearCustomDocumentProperties,omitempty"`
	///Value indicating whether built in document properties should be cleared.
	ClearBuiltInDocumentProperties bool `json:"ClearBuiltInDocumentProperties,omitempty"`
	// Implements GroupDocs.Conversion.Contracts.IDocumentsContainerLoadOptions.Depth     Default: 1
	Depth int `json:"Depth,omitempty"`
	// Implements GroupDocs.Conversion.Contracts.IDocumentsContainerLoadOptions.ConvertOwned     Default is false
	ConvertOwned bool `json:"ConvertOwned,omitempty"`
	// Implements GroupDocs.Conversion.Contracts.IDocumentsContainerLoadOptions.ConvertOwner     Default is true
	ConvertOwner bool `json:"ConvertOwner,omitempty"`
}
