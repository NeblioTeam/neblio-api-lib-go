# \InsightApi

All URIs are relative to *https://ntp1node.nebl.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddress**](InsightApi.md#GetAddress) | **Get** /ins/addr/{address} | Returns address object
[**GetAddressBalance**](InsightApi.md#GetAddressBalance) | **Get** /ins/addr/{address}/balance | Returns address balance in sats
[**GetAddressTotalReceived**](InsightApi.md#GetAddressTotalReceived) | **Get** /ins/addr/{address}/totalReceived | Returns total received by address in sats
[**GetAddressTotalSent**](InsightApi.md#GetAddressTotalSent) | **Get** /ins/addr/{address}/totalSent | Returns total sent by address in sats
[**GetAddressUnconfirmedBalance**](InsightApi.md#GetAddressUnconfirmedBalance) | **Get** /ins/addr/{address}/unconfirmedBalance | Returns address unconfirmed balance in sats
[**GetAddressUtxos**](InsightApi.md#GetAddressUtxos) | **Get** /ins/addr/{address}/utxo | Returns all UTXOs at a given address
[**GetBlock**](InsightApi.md#GetBlock) | **Get** /ins/block/{blockhash} | Returns information regarding a Neblio block
[**GetBlockIndex**](InsightApi.md#GetBlockIndex) | **Get** /ins/block-index/{blockindex} | Returns block hash of block
[**GetRawTx**](InsightApi.md#GetRawTx) | **Get** /ins/rawtx/{txid} | Returns raw transaction hex
[**GetStatus**](InsightApi.md#GetStatus) | **Get** /ins/status | Utility API for calling several blockchain node functions
[**GetSync**](InsightApi.md#GetSync) | **Get** /ins/sync | Get node sync status
[**GetTx**](InsightApi.md#GetTx) | **Get** /ins/tx/{txid} | Returns transaction object
[**GetTxs**](InsightApi.md#GetTxs) | **Get** /ins/txs | Get transactions by block or address
[**SendTx**](InsightApi.md#SendTx) | **Post** /ins/tx/send | Broadcasts a signed raw transaction to the network (not NTP1 specific)


# **GetAddress**
> GetAddressResponse GetAddress(ctx, address)
Returns address object

Returns NEBL address object containing information on a specific address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **address** | **string**| Address | 

### Return type

[**GetAddressResponse**](getAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddressBalance**
> GetAddressBalanceResponse GetAddressBalance(ctx, address)
Returns address balance in sats

Returns NEBL address balance in satoshis

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **address** | **string**| Address | 

### Return type

[**GetAddressBalanceResponse**](getAddressBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddressTotalReceived**
> GetAddressTotalReceivedResponse GetAddressTotalReceived(ctx, address)
Returns total received by address in sats

Returns total NEBL received by address in satoshis

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **address** | **string**| Address | 

### Return type

[**GetAddressTotalReceivedResponse**](getAddressTotalReceivedResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddressTotalSent**
> GetAddressTotalSentResponse GetAddressTotalSent(ctx, address)
Returns total sent by address in sats

Returns total NEBL sent by address in satoshis

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **address** | **string**| Address | 

### Return type

[**GetAddressTotalSentResponse**](getAddressTotalSentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddressUnconfirmedBalance**
> GetAddressUnconfirmedBalanceResponse GetAddressUnconfirmedBalance(ctx, address)
Returns address unconfirmed balance in sats

Returns NEBL address unconfirmed balance in satoshis

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **address** | **string**| Address | 

### Return type

[**GetAddressUnconfirmedBalanceResponse**](getAddressUnconfirmedBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddressUtxos**
> GetAddressUtxosResponse GetAddressUtxos(ctx, address)
Returns all UTXOs at a given address

Returns information on each Unspent Transaction Output contained at an address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **address** | **string**| Address | 

### Return type

[**GetAddressUtxosResponse**](getAddressUtxosResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlock**
> GetBlockResponse GetBlock(ctx, blockhash)
Returns information regarding a Neblio block

Returns blockchain data for a given block based upon the block hash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockhash** | **string**| Block Hash | 

### Return type

[**GetBlockResponse**](getBlockResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockIndex**
> GetBlockIndexResponse GetBlockIndex(ctx, blockindex)
Returns block hash of block

Returns the block hash of a block at a given block index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockindex** | **float32**| Block Index | 

### Return type

[**GetBlockIndexResponse**](getBlockIndexResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRawTx**
> GetRawTxResponse GetRawTx(ctx, txid)
Returns raw transaction hex

Returns raw transaction hex representing a NEBL transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**GetRawTxResponse**](getRawTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatus**
> GetStatusResponse GetStatus(ctx, optional)
Utility API for calling several blockchain node functions

tility API for calling several blockchain node functions - getInfo, getDifficulty, getBestBlockHash, getLastBlockHash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string**| Function to call, getInfo, getDifficulty, getBestBlockHash, or getLastBlockHash | 

### Return type

[**GetStatusResponse**](getStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSync**
> GetSyncResponse GetSync(ctx, )
Get node sync status

Returns information on the node's sync progress

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetSyncResponse**](getSyncResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTx**
> GetTxResponse GetTx(ctx, txid)
Returns transaction object

Returns NEBL transaction object representing a NEBL transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**GetTxResponse**](getTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTxs**
> GetTxsResponse GetTxs(ctx, optional)
Get transactions by block or address

Returns all transactions by block or address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string**| Address | 
 **block** | **string**| Block Hash | 
 **page** | **float32**| Page number to display | 

### Return type

[**GetTxsResponse**](getTxsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTx**
> BroadcastTxResponse SendTx(ctx, body)
Broadcasts a signed raw transaction to the network (not NTP1 specific)

Broadcasts a signed raw transaction to the network. If successful returns the txid of the broadcast trasnaction. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**SendTxRequest**](SendTxRequest.md)| Object representing a transaction to broadcast | 

### Return type

[**BroadcastTxResponse**](broadcastTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

