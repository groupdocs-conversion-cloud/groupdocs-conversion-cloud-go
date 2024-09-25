/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

// Metered license consumption information
type ConsumptionResult struct {
	// Amount of used credits
	Credit float32 `json:"Credit"`
	// Amount of MBs processed
	Quantity float32 `json:"Quantity"`
}
