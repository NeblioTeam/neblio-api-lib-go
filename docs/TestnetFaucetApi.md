# \TestnetFaucetApi

All URIs are relative to *https://ntp1node.nebl.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestnetGetFaucet**](TestnetFaucetApi.md#TestnetGetFaucet) | **Get** /testnet/faucet | Withdraws testnet NEBL to the specified address


# **TestnetGetFaucet**
> GetFaucetResponse TestnetGetFaucet(ctx, address, optional)
Withdraws testnet NEBL to the specified address

Withdraw testnet NEBL to your Neblio Testnet address. By default amount is 1500000000 or 15 NEBL and has a max of 50 NEBL. Only 2 withdrawals allowed per 24 hour period. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Your Neblio Testnet Address | 
 **optional** | ***TestnetGetFaucetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestnetGetFaucetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amount** | **optional.Float32**| Amount of NEBL to withdrawal in satoshis | 

### Return type

[**GetFaucetResponse**](getFaucetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

