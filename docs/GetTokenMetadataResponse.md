# GetTokenMetadataResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | ID of the token | [optional] [default to null]
**Divisibility** | **float32** | Decimal places the token is divisible to | [optional] [default to null]
**LockStatus** | **bool** | Whether issuance of more tokens is locked | [optional] [default to null]
**AggregationPolicy** | **string** | Whether the tokens are aggregatable | [optional] [default to null]
**TotalSupply** | **float32** | Total number of tokens in supply | [optional] [default to null]
**NumOfHolders** | **float32** | Total number of addresses this token is held at | [optional] [default to null]
**NumOfTransfers** | **float32** | Total number of transactions of this token | [optional] [default to null]
**NumofIssuance** | **float32** | Total number of times this token has been issued | [optional] [default to null]
**NumOfBurns** | **float32** | Number of times tokens have been burned | [optional] [default to null]
**FirstBlock** | **float32** | Block number token was issued in | [optional] [default to null]
**IssuanceTxid** | **string** | TXID the token was issued with | [optional] [default to null]
**IssueAddress** | **string** | Address that issued the tokens | [optional] [default to null]
**MetadataOfIssuance** | [***GetTokenMetadataResponseMetadataOfIssuance**](getTokenMetadataResponse_metadataOfIssuance.md) |  | [optional] [default to null]
**MetadataOfUtxo** | [***GetTokenMetadataResponseMetadataOfIssuance**](getTokenMetadataResponse_metadataOfIssuance.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


