# GroupDocs.Conversion Cloud Go SDK

This repository contains GroupDocs.Conversion Cloud SDK for Go source code. This SDK has been developed to help you get started with using our document conversion REST API, allowing to seamlessly convert your documents to any format you need. With this single API, you can convert back and forth between over 50 types of documents and images, including all Microsoft Office and OpenDocument file formats, PDF documents, HTML, CAD, raster images and many more.

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

### Create an account
Creating an account is very simple. Go to Dashboard to create a free account.
We’re using Single Sign On across our websites, therefore, if you already have an account with our services, you can use it to also acccess the [Dashboard](https://dashboard.groupdocs.cloud).

### Create an API client app
Before you can make any requests to GroupDocs Cloud API you need to get a Client Id and a Client Secret. This will will be used to invoke GroupDocs Cloud API. You can get it by creating a new [Application](https://dashboard.groupdocs.cloud/applicationsV).

## Convert document

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
    ClientID:     "XXXX-XXXX-XXXX-XXXX",      // Your client_id
    ClientSecret: "XXXXXXXXXXXXXXXX",         // Your client_secret
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
    ctx := context.WithValue(context.Background(), conversion.ContextAccessToken, token.AccessToken)


    path := "./myFile.docx"
    localFile, err := os.Open(path)
    if err != nil {
        fmt.Printf("Failed to open file %s: %v\n", path, err)
        return
    }
    defer localFile.Close()

    result, _, err := client.ConvertApi.ConvertDocumentDirect(ctx, "pdf", localFile, nil)

    if err != nil {
        fmt.Printf("ConvertToPdfDirect error: %v\n", err)
        return
    }

    fi, _ := result.Stat()

    fmt.Printf("Document converted successfully: %v\n", fi.Size())
  
}
```

## Convert document using cloud storage

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
    ClientID:     "XXXX-XXXX-XXXX-XXXX",      // Your client_id
    ClientSecret: "XXXXXXXXXXXXXXXX",         // Your client_secret
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
    ctx := context.WithValue(context.Background(), conversion.ContextAccessToken, token.AccessToken)

    // Open the local file
    path := "./myFile.docx"
    localFile, err := os.Open(path)
    if err != nil {
        fmt.Printf("Failed to open file %s: %v\n", path, err)
        return
    }
    defer localFile.Close()

    // Upload the file to cloud storage with the relative path
    _, _, err = client.FileApi.UploadFile(Ctx, "myFile.docx", localFile, nil)
    if err != nil {
        fmt.Printf("Failed to upload file %s: %v\n", path, err)
    } else {
        fmt.Printf("File %s uploaded successfully.\n", path)
    }

    // Perform conversion
	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "myFile.docx",
		OutputPath: "converted",
		ConvertOptions: &models.PdfConvertOptions{
			CenterWindow:         true,
			CompressImages:       false,
			DisplayDocTitle:      true,
			Dpi:                  1024,
			FitWindow:            false,
			FromPage:             1,
			Grayscale:            false,
			ImageQuality:         100,
			Linearize:            false,
			MarginTop:            5,
			MarginLeft:           5,
			UnembedFonts:         true,
			RemoveUnusedStreams:  true,
			RemoveUnusedObjects:  true,
			RemovePdfaCompliance: false,
		},
	}

	result, _, err := client.ConvertApi.ConvertDocument(ctx, settings)

	if err != nil {
		fmt.Printf("ConvertToPdf error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)

    // Download file
    var dlPath = "converted/myFile.pdf"
    dlFileResp, _, err := client.FileApi.DownloadFile(ctx, dlPath, nil)
    if err != nil {
      fmt.Print("Failed to download file %v\n", err)
    }

    // Get file info
    fileInfo, errInfo := dlFileResp.Stat()
    if errInfo != nil {
      fmt.Print(errInfo)
    }

    // Get the size of the file
    fileSize := fileInfo.Size()

    if fileSize == 0 {
      fmt.Print("File size is zero")
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
