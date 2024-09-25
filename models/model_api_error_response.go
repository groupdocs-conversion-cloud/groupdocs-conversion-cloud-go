/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 * API version: 24.8
 */

package models

type ApiErrorResponse struct {
	RequestId string    `json:"RequestId,omitempty"`
	Error_    *ApiError `json:"Error,omitempty"`
}
