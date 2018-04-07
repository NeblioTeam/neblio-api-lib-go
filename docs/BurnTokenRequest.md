# BurnTokenRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fee** | **float32** | Fee in satoshi to include in the issuance transaction min 10000 (0.0001 NEBL) | [default to null]
**From** | **[]string** | Array of addresses to send the token from | [optional] [default to null]
**Transfer** | [**[]SendTokenRequestTo**](sendTokenRequest_to.md) |  | [optional] [default to null]
**Burn** | [**[]BurnTokenRequestBurn**](burnTokenRequest_burn.md) | Array of objects representing tokens to be burned | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


