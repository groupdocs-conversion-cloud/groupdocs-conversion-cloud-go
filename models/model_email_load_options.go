/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 */

package models

// Options for loading Email documents
type EmailLoadOptions struct {
	// The format of input file, (\"docx\", for example). This field must be filled with correct input file format when using ConvertDirect method, which accept input file as binary stream, and, because of that, API can correctly handle LoadOptions. In regular conversion, the input file format taken from the input file name and this field ignored.
	Format string `json:"Format,omitempty"`
	// Option to display or hide the email header. Default: true
	DisplayHeader bool `json:"DisplayHeader,omitempty"`
	// Option to display or hide \"from\" email address. Default: true
	DisplayFromEmailAddress bool `json:"DisplayFromEmailAddress"`
	// Option to display or hide \"to\" email address. Default: true
	DisplayToEmailAddress bool `json:"DisplayToEmailAddress"`
	// Option to display or hide \"Cc\" email address. Default: false
	DisplayCcEmailAddress bool `json:"DisplayCcEmailAddress"`
	// Option to display or hide \"Bcc\" email address. Default: false
	DisplayBccEmailAddress bool `json:"DisplayBccEmailAddress"`
	// Gets or sets the Coordinated Universal Time (UTC) offset for the message dates. This property defines the time zone difference, between the localtime and UTC.
	TimeZoneOffset string `json:"TimeZoneOffset,omitempty"`
	// Option to convert attachments in source email or not. Default: false.
	ConvertAttachments bool `json:"ConvertAttachments"`
	// The mapping between email message field and field text representation
	FieldLabels []FieldLabel `json:"FieldLabels,omitempty"`
	// Defines whether need to keep original date header string in mail message when saving or not (Default value is true)
	PreserveOriginalDate bool `json:"PreserveOriginalDate,omitempty"`
}
