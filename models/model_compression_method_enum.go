/*
 * GroupDocs.Conversion Cloud API Reference
 *
 * Please visit [Dashboard](https://dashboard.groupdocs.cloud/#/applications) to obtain your Client Id and Client Secret keys.
 *
 */

package models

type CompressionMethodEnum string

const (
	CompressionMethodEnumRaw                  CompressionMethodEnum = "Raw"
	CompressionMethodEnumRle                  CompressionMethodEnum = "Rle"
	CompressionMethodEnumZipWithoutPrediction CompressionMethodEnum = "ZipWithoutPrediction"
	CompressionMethodEnumZipWithPrediction    CompressionMethodEnum = "ZipWithPrediction"
)
