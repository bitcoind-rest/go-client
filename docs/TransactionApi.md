# \TransactionApi

All URIs are relative to *http://localhost:3000/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestTx**](TransactionApi.md#RestTx) | **Get** /tx/{txHash} | Get transaction by hash.


# **RestTx**
> Transaction RestTx(ctx, txHash)
Get transaction by hash.

Given a transaction hash: returns a transaction in binary, hex-encoded binary, or JSON formats. For full TX query capability, one must enable the transaction index via txindex=1 command line / configuration option.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txHash** | **string**| The transaction hash. | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream, text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

