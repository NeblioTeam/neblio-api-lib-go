# SendTokenRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fee** | **float32** | Fee in satoshi to include in the issuance transaction min 10000 (0.0001 NEBL) | [default to null]
**From** | **[]string** | Array of addresses to send the token from | [optional] [default to null]
**Sendutxo** | **[]string** | Array of UTXOs to send the token from | [optional] [default to null]
**To** | [**[]SendTokenRequestTo**](sendTokenRequest_to.md) |  | [default to null]
**Flags** | [***IssueTokenRequestFlags**](issueTokenRequest_flags.md) |  | [optional] [default to null]
**Metadata** | [***IssueTokenRequestMetadata**](issueTokenRequest_metadata.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


