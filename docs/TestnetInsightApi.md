# \TestnetInsightApi

All URIs are relative to *https://ntp1node.nebl.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestnetGetAddress**](TestnetInsightApi.md#TestnetGetAddress) | **Get** /testnet/ins/addr/{address} | Returns address object
[**TestnetGetAddressBalance**](TestnetInsightApi.md#TestnetGetAddressBalance) | **Get** /testnet/ins/addr/{address}/balance | Returns address balance in sats
[**TestnetGetAddressTotalReceived**](TestnetInsightApi.md#TestnetGetAddressTotalReceived) | **Get** /testnet/ins/addr/{address}/totalReceived | Returns total received by address in sats
[**TestnetGetAddressTotalSent**](TestnetInsightApi.md#TestnetGetAddressTotalSent) | **Get** /testnet/ins/addr/{address}/totalSent | Returns total sent by address in sats
[**TestnetGetAddressUnconfirmedBalance**](TestnetInsightApi.md#TestnetGetAddressUnconfirmedBalance) | **Get** /testnet/ins/addr/{address}/unconfirmedBalance | Returns address unconfirmed balance in sats
[**TestnetGetAddressUtxos**](TestnetInsightApi.md#TestnetGetAddressUtxos) | **Get** /testnet/ins/addr/{address}/utxo | Returns all UTXOs at a given address
[**TestnetGetBlock**](TestnetInsightApi.md#TestnetGetBlock) | **Get** /testnet/ins/block/{blockhash} | Returns information regarding a Neblio block
[**TestnetGetBlockIndex**](TestnetInsightApi.md#TestnetGetBlockIndex) | **Get** /testnet/ins/block-index/{blockindex} | Returns block hash of block
[**TestnetGetRawTx**](TestnetInsightApi.md#TestnetGetRawTx) | **Get** /testnet/ins/rawtx/{txid} | Returns raw transaction hex
[**TestnetGetStatus**](TestnetInsightApi.md#TestnetGetStatus) | **Get** /testnet/ins/status | Utility API for calling several blockchain node functions
[**TestnetGetSync**](TestnetInsightApi.md#TestnetGetSync) | **Get** /testnet/ins/sync | Get node sync status
[**TestnetGetTx**](TestnetInsightApi.md#TestnetGetTx) | **Get** /testnet/ins/tx/{txid} | Returns transaction object
[**TestnetGetTxs**](TestnetInsightApi.md#TestnetGetTxs) | **Get** /testnet/ins/txs | Get transactions by block or address
[**TestnetSendTx**](TestnetInsightApi.md#TestnetSendTx) | **Post** /testnet/ins/tx/send | Broadcasts a signed raw transaction to the network (not NTP1 specific)


# **TestnetGetAddress**
> GetAddressResponse TestnetGetAddress(ctx, address)
Returns address object

Returns NEBL address object containing information on a specific address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address | 

### Return type

[**GetAddressResponse**](getAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetAddressBalance**
> float32 TestnetGetAddressBalance(ctx, address)
Returns address balance in sats

Returns NEBL address balance in satoshis

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address | 

### Return type

**float32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetAddressTotalReceived**
> float32 TestnetGetAddressTotalReceived(ctx, address)
Returns total received by address in sats

Returns total NEBL received by address in satoshis

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address | 

### Return type

**float32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetAddressTotalSent**
> float32 TestnetGetAddressTotalSent(ctx, address)
Returns total sent by address in sats

Returns total NEBL sent by address in satoshis

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address | 

### Return type

**float32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetAddressUnconfirmedBalance**
> float32 TestnetGetAddressUnconfirmedBalance(ctx, address)
Returns address unconfirmed balance in sats

Returns NEBL address unconfirmed balance in satoshis

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address | 

### Return type

**float32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetAddressUtxos**
> []map[string]interface{} TestnetGetAddressUtxos(ctx, address)
Returns all UTXOs at a given address

Returns information on each Unspent Transaction Output contained at an address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetBlock**
> GetBlockResponse TestnetGetBlock(ctx, blockhash)
Returns information regarding a Neblio block

Returns blockchain data for a given block based upon the block hash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockhash** | **string**| Block Hash | 

### Return type

[**GetBlockResponse**](getBlockResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetBlockIndex**
> GetBlockIndexResponse TestnetGetBlockIndex(ctx, blockindex)
Returns block hash of block

Returns the block hash of a block at a given block index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockindex** | **float32**| Block Index | 

### Return type

[**GetBlockIndexResponse**](getBlockIndexResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetRawTx**
> GetRawTxResponse TestnetGetRawTx(ctx, txid)
Returns raw transaction hex

Returns raw transaction hex representing a NEBL transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**GetRawTxResponse**](getRawTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetStatus**
> map[string]interface{} TestnetGetStatus(ctx, optional)
Utility API for calling several blockchain node functions

Utility API for calling several blockchain node functions - getInfo, getDifficulty, getBestBlockHash, getLastBlockHash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TestnetGetStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestnetGetStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Function to call, getInfo, getDifficulty, getBestBlockHash, or getLastBlockHash | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetSync**
> GetSyncResponse TestnetGetSync(ctx, )
Get node sync status

Returns information on the node's sync progress

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetSyncResponse**](getSyncResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetTx**
> GetTxResponse TestnetGetTx(ctx, txid)
Returns transaction object

Returns NEBL transaction object representing a NEBL transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**GetTxResponse**](getTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetTxs**
> GetTxsResponse TestnetGetTxs(ctx, optional)
Get transactions by block or address

Returns all transactions by block or address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TestnetGetTxsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestnetGetTxsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **optional.String**| Address | 
 **block** | **optional.String**| Block Hash | 
 **pageNum** | **optional.Float32**| Page number to display | 

### Return type

[**GetTxsResponse**](getTxsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetSendTx**
> BroadcastTxResponse TestnetSendTx(ctx, sendTxRequest)
Broadcasts a signed raw transaction to the network (not NTP1 specific)

Broadcasts a signed raw transaction to the network. If successful returns the txid of the broadcast trasnaction. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sendTxRequest** | [**SendTxRequest**](SendTxRequest.md)| Object representing a transaction to broadcast | 

### Return type

[**BroadcastTxResponse**](broadcastTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

