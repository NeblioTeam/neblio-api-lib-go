# GetTokenMetadataResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | ID of the token | [optional] 
**SomeUtxo** | **string** | Example UTXO containing this token. | [optional] 
**Divisibility** | **float32** | Decimal places the token is divisible to | [optional] 
**LockStatus** | **bool** | Whether issuance of more tokens is locked | [optional] 
**AggregationPolicy** | **string** | Whether the tokens are aggregatable | [optional] 
**TotalSupply** | **float32** | Total number of tokens in supply | [optional] 
**NumOfHolders** | **float32** | Total number of addresses this token is held at | [optional] 
**NumOfTransfers** | **float32** | Total number of transactions of this token | [optional] 
**NumOfIssuance** | **float32** | Total number of times this token has been issued | [optional] 
**NumOfBurns** | **float32** | Number of times tokens have been burned | [optional] 
**FirstBlock** | **float32** | Block number token was issued in | [optional] 
**IssuanceTxid** | **string** | TXID the token was issued with | [optional] 
**IssueAddress** | **string** | Address that issued the tokens | [optional] 
**MetadataOfIssuence** | [**GetTokenMetadataResponseMetadataOfIssuence**](getTokenMetadataResponse_metadataOfIssuence.md) |  | [optional] 
**MetadataOfUtxo** | [**GetTokenMetadataResponseMetadataOfUtxo**](getTokenMetadataResponse_metadataOfUtxo.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


