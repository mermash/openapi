# \AuthCheckerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthCheckerCheck**](AuthCheckerApi.md#AuthCheckerCheck) | **Get** /v1/session/check/{ID} | 
[**AuthCheckerCreate**](AuthCheckerApi.md#AuthCheckerCreate) | **Post** /v1/session/create | 
[**AuthCheckerDelete**](AuthCheckerApi.md#AuthCheckerDelete) | **Post** /v1/session/delete | 



## AuthCheckerCheck

> SessionSession AuthCheckerCheck(ctx, iD).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    iD := "iD_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthCheckerApi.AuthCheckerCheck(context.Background(), iD).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthCheckerApi.AuthCheckerCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthCheckerCheck`: SessionSession
    fmt.Fprintf(os.Stdout, "Response from `AuthCheckerApi.AuthCheckerCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iD** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthCheckerCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SessionSession**](SessionSession.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthCheckerCreate

> SessionSessionID AuthCheckerCreate(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewSessionSession() // SessionSession | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthCheckerApi.AuthCheckerCreate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthCheckerApi.AuthCheckerCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthCheckerCreate`: SessionSessionID
    fmt.Fprintf(os.Stdout, "Response from `AuthCheckerApi.AuthCheckerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthCheckerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SessionSession**](SessionSession.md) |  | 

### Return type

[**SessionSessionID**](SessionSessionID.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthCheckerDelete

> SessionNothing AuthCheckerDelete(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewSessionSessionID() // SessionSessionID | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthCheckerApi.AuthCheckerDelete(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthCheckerApi.AuthCheckerDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthCheckerDelete`: SessionNothing
    fmt.Fprintf(os.Stdout, "Response from `AuthCheckerApi.AuthCheckerDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthCheckerDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SessionSessionID**](SessionSessionID.md) |  | 

### Return type

[**SessionNothing**](SessionNothing.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

