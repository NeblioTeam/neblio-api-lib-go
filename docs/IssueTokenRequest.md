# IssueTokenRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueAddress** | **string** | Address issuing the token | [default to null]
**Amount** | **float32** | Number of tokens to issue | [default to null]
**Divisibility** | **float32** | Number of decimal places the token should be divisble by (0-7) | [default to null]
**Fee** | **float32** | Fee in satoshi to include in the issuance transaction min 1000000000 (10 NEBL) | [default to null]
**Reissuable** | **bool** | whether the token should be reissuable | [default to null]
**Metadata** | [***IssueTokenRequestMetadata**](issueTokenRequest_metadata.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


