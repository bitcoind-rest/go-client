# \MemoryPoolApi

All URIs are relative to *http://localhost:3000/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestGetutxos**](MemoryPoolApi.md#RestGetutxos) | **Get** /getutxos/checkmempool/{txHash}-{txOutput}.{format} | Returns Unspent Transaction (TX) Outputs
[**RestHeaders**](MemoryPoolApi.md#RestHeaders) | **Get** /headers/{headerCount}/{blockHash}.{format} | Returns headers.
[**RestMempoolContents**](MemoryPoolApi.md#RestMempoolContents) | **Get** /mempool/contents.json | Returns transactions in the TX mempool.
[**RestMempoolInfo**](MemoryPoolApi.md#RestMempoolInfo) | **Get** /mempool/info.json | Returns various information about the TX mempool.


# **RestGetutxos**
> InlineResponse200 RestGetutxos(ctx, txHash, txOutput, format)
Returns Unspent Transaction (TX) Outputs

Only supports JSON as output format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txHash** | **string**| The transaction hash | 
  **txOutput** | **string**| The transaction output | 
  **format** | **string**| The expected format | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestHeaders**
> string RestHeaders(ctx, headerCount, blockHash, format)
Returns headers.

Only supports BIN and HEX as output format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **headerCount** | **int32**| The header count | 
  **blockHash** | **string**| The block hash | 
  **format** | **string**| The expected format | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestMempoolContents**
> InlineResponseDefault RestMempoolContents(ctx, )
Returns transactions in the TX mempool.

Only supports JSON as output format.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponseDefault**](inline_response_default.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestMempoolInfo**
> MemoryPool RestMempoolInfo(ctx, )
Returns various information about the TX mempool.

Only supports JSON as output format.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MemoryPool**](MemoryPool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

