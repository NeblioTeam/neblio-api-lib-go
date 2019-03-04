# SendTokenRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fee** | **float32** | Fee in satoshi to include in the issuance transaction min 10000 (0.0001 NEBL) | 
**From** | **[]string** | Array of addresses to send the token from | [optional] 
**Sendutxo** | **[]string** | Array of UTXOs to send the token from | [optional] 
**To** | [**[]SendTokenRequestTo**](sendTokenRequest_to.md) |  | 
**Flags** | [**IssueTokenRequestFlags**](issueTokenRequest_flags.md) |  | [optional] 
**Metadata** | [**IssueTokenRequestMetadata**](issueTokenRequest_metadata.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


