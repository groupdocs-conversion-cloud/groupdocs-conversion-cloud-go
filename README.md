# GroupDocs.Conversion Cloud Go SDK

Go package for communicating with the GroupDocs.Conversion Cloud API

## Installation

The package is available at [github.com](https://github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go). You can install it with:

```shell
go get github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go
```

To add dependency to your app copy following into your go.mod and run `go mod tidy`:

```go
require (
 github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go v1.0.0
)
```

## Getting Started

Please follow the [installation](#installation) procedure and then run the following code:

```go
package main

import (
 "context"
 "fmt"

 "github.com/antihax/optional"
 conversion "github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go"
 "golang.org/x/oauth2/clientcredentials"
)

func main() {

 // Create an OAuth2 config using client_credentials grant type
 conf := &clientcredentials.Config{
  ClientID:     "961234b6-aa74-40c2-a0c0-9d6e2ff92eab",      // Your client_id
  ClientSecret: "a1b6c8aa82e0cc3fc642b4bc0e1d38b2",          // Your client_secret
  TokenURL:     "https://api.groupdocs.cloud/connect/token", // The token URL
  Scopes:       []string{},                                  // Optional: specify any required scopes
 }

 // Request the token
 token, err := conf.Token(context.Background())
 if err != nil {
  fmt.Printf("Unable to retrieve token: %v", err)
 }

 // Output the access token
 fmt.Printf("Access Token: %s\n", token.AccessToken)

 cfg := conversion.NewConfiguration()
 client := conversion.NewAPIClient(cfg)
 auth := context.WithValue(context.Background(), conversion.ContextAccessToken, token.AccessToken)
 opts := &conversion.InfoApiGetSupportedConversionTypesOpts{
  Format: optional.NewString("pdf"),
 }

 formats, _, _ := client.InfoApi.GetSupportedConversionTypes(auth, opts)

 for _, format := range formats {
  fmt.Printf("Source Format: %s\n", format.SourceFormat)
  fmt.Printf("Target Formats: %v\n", format.TargetFormats)
  fmt.Println()
 }
}
```

## Licensing

GroupDocs.Conversion Cloud Go SDK licensed under [MIT License](LICENSE).

## Resources

+ [**Website**](https://www.groupdocs.cloud)
+ [**Product Home**](https://products.groupdocs.cloud/conversion)
+ [**Documentation**](https://docs.groupdocs.cloud/display/conversioncloud/Home)
+ [**Free Support Forum**](https://forum.groupdocs.cloud/c/conversion)
+ [**Blog**](https://blog.groupdocs.cloud/category/conversion)

## Contact Us

Your feedback is very important to us. Please feel free to contact us using our [Support Forums](https://forum.groupdocs.cloud/c/conversion).
