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
	// Hide comments
	HideComments bool `json:"HideComments"`
	// Show hidden slides
	ShowHiddenSlides bool `json:"ShowHiddenSlides"`
}
