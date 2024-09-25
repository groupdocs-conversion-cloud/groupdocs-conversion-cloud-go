/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

import (
	"time"
)

// Contains a document metadata
type DocumentMetadata struct {
	// Document file type
	FileType string `json:"FileType,omitempty"`
	// Gets pages count if applicable to the current document format
	PageCount int32 `json:"PageCount"`
	// Document bytes size
	Size int64 `json:"Size"`
	// Returns detected width if applicable to the current document format
	Width int32 `json:"Width"`
	// Returns detected height if applicable to the current document format
	Height int32 `json:"Height"`
	// Returns detected horizontal resolution if applicable to the current document format
	HorizontalResolution int32 `json:"HorizontalResolution"`
	// Returns detected vertical resolution if applicable to the current document format
	VerticalResolution int32 `json:"VerticalResolution"`
	// Returns detected bits per pixel if applicable to the current document format
	BitsPerPixel int32 `json:"BitsPerPixel"`
	// Returns document title width if applicable to the current document format
	Title string `json:"Title,omitempty"`
	// Returns detected document author if applicable to the current document format
	Author string `json:"Author,omitempty"`
	// Returns detected document creation date if it's applicable to the current document format
	CreatedDate time.Time `json:"CreatedDate"`
	// Returns detected document modification date if applicable to the current document format
	ModifiedDate time.Time `json:"ModifiedDate"`
	// Returns list of layer names if applicable to the current document format
	Layers []string `json:"Layers,omitempty"`
	// Is document password protected
	IsPasswordProtected bool `json:"IsPasswordProtected"`
}
