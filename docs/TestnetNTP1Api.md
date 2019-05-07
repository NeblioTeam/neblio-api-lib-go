# \TestnetNTP1Api

All URIs are relative to *https://ntp1node.nebl.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestnetBroadcastTx**](TestnetNTP1Api.md#TestnetBroadcastTx) | **Post** /testnet/ntp1/broadcast | Broadcasts a signed raw transaction to the network
[**TestnetBurnToken**](TestnetNTP1Api.md#TestnetBurnToken) | **Post** /testnet/ntp1/burntoken | Builds a transaction that burns an NTP1 Token
[**TestnetGetAddressInfo**](TestnetNTP1Api.md#TestnetGetAddressInfo) | **Get** /testnet/ntp1/addressinfo/{address} | Information On a Neblio Address
[**TestnetGetTokenHolders**](TestnetNTP1Api.md#TestnetGetTokenHolders) | **Get** /testnet/ntp1/stakeholders/{tokenid} | Get Addresses Holding a Token
[**TestnetGetTokenId**](TestnetNTP1Api.md#TestnetGetTokenId) | **Get** /testnet/ntp1/tokenid/{tokensymbol} | Returns the tokenId representing a token
[**TestnetGetTokenMetadata**](TestnetNTP1Api.md#TestnetGetTokenMetadata) | **Get** /testnet/ntp1/tokenmetadata/{tokenid} | Get Metadata of Token
[**TestnetGetTokenMetadataOfUtxo**](TestnetNTP1Api.md#TestnetGetTokenMetadataOfUtxo) | **Get** /testnet/ntp1/tokenmetadata/{tokenid}/{utxo} | Get UTXO Metadata of Token
[**TestnetGetTransactionInfo**](TestnetNTP1Api.md#TestnetGetTransactionInfo) | **Get** /testnet/ntp1/transactioninfo/{txid} | Information On an NTP1 Transaction
[**TestnetIssueToken**](TestnetNTP1Api.md#TestnetIssueToken) | **Post** /testnet/ntp1/issue | Builds a transaction that issues a new NTP1 Token
[**TestnetSendToken**](TestnetNTP1Api.md#TestnetSendToken) | **Post** /testnet/ntp1/sendtoken | Builds a transaction that sends an NTP1 Token


# **TestnetBroadcastTx**
> BroadcastTxResponse TestnetBroadcastTx(ctx, broadcastTxRequest)
Broadcasts a signed raw transaction to the network

Broadcasts a signed raw transaction to the network. If successful returns the txid of the broadcast trasnaction. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **broadcastTxRequest** | [**BroadcastTxRequest**](BroadcastTxRequest.md)| Object representing a transaction to broadcast | 

### Return type

[**BroadcastTxResponse**](broadcastTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetBurnToken**
> BurnTokenResponse TestnetBurnToken(ctx, burnTokenRequest)
Builds a transaction that burns an NTP1 Token

Builds an unsigned raw transaction that burns an NTP1 token on the Neblio blockchain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **burnTokenRequest** | [**BurnTokenRequest**](BurnTokenRequest.md)| Object representing the token to be burned | 

### Return type

[**BurnTokenResponse**](burnTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetAddressInfo**
> GetAddressInfoResponse TestnetGetAddressInfo(ctx, address)
Information On a Neblio Address

Returns both NEBL and NTP1 token UTXOs held at the given address. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Neblio Address to get information on. | 

### Return type

[**GetAddressInfoResponse**](getAddressInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetTokenHolders**
> GetTokenHoldersResponse TestnetGetTokenHolders(ctx, tokenid)
Get Addresses Holding a Token

Returns the the the addresses holding a token and how many tokens are held 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenid** | **string**| TokenId to request metadata for | 

### Return type

[**GetTokenHoldersResponse**](getTokenHoldersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetTokenId**
> GetTokenIdResponse TestnetGetTokenId(ctx, tokensymbol)
Returns the tokenId representing a token

Translates a token symbol to a tokenId if a token exists with that symbol on the network 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokensymbol** | **string**| Token symbol | 

### Return type

[**GetTokenIdResponse**](getTokenIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetTokenMetadata**
> GetTokenMetadataResponse TestnetGetTokenMetadata(ctx, tokenid)
Get Metadata of Token

Returns the metadata associated with a token. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenid** | **string**| TokenId to request metadata for | 

### Return type

[**GetTokenMetadataResponse**](getTokenMetadataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetTokenMetadataOfUtxo**
> GetTokenMetadataResponse TestnetGetTokenMetadataOfUtxo(ctx, tokenid, utxo)
Get UTXO Metadata of Token

Returns the metadata associated with a token for that specific utxo instead of the issuance transaction. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenid** | **string**| TokenId to request metadata for | 
  **utxo** | **string**| Specific UTXO to request metadata for | 

### Return type

[**GetTokenMetadataResponse**](getTokenMetadataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetGetTransactionInfo**
> GetTransactionInfoResponse TestnetGetTransactionInfo(ctx, txid)
Information On an NTP1 Transaction

Returns detailed information regarding an NTP1 transaction. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Neblio txid to get information on. | 

### Return type

[**GetTransactionInfoResponse**](getTransactionInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetIssueToken**
> IssueTokenResponse TestnetIssueToken(ctx, issueTokenRequest)
Builds a transaction that issues a new NTP1 Token

Builds an unsigned raw transaction that issues a new NTP1 token on the Neblio blockchain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueTokenRequest** | [**IssueTokenRequest**](IssueTokenRequest.md)| Object representing the token to be created | 

### Return type

[**IssueTokenResponse**](issueTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestnetSendToken**
> SendTokenResponse TestnetSendToken(ctx, sendTokenRequest)
Builds a transaction that sends an NTP1 Token

Builds an unsigned raw transaction that sends an NTP1 token on the Neblio blockchain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sendTokenRequest** | [**SendTokenRequest**](SendTokenRequest.md)| Object representing the token to be sent | 

### Return type

[**SendTokenResponse**](sendTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

