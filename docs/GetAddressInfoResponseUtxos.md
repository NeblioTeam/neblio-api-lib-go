# GetAddressInfoResponseUtxos

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **float32** | Index of the UTXO at this address | [optional] 
**Txid** | **string** | Txid of this UTXO | [optional] 
**Blockheight** | **float32** | Blockheight of the UTXO | [optional] 
**Blocktime** | **float32** | Blocktime of the UTXO | [optional] 
**ScriptPubKey** | [**map[string]interface{}**](.md) | Object representing the scruptPubKey of the UTXO | [optional] 
**Used** | **bool** | Whether the UTXO has been used | [optional] 
**Value** | **float32** | Value of the UTXO in NEBL satoshi | [optional] 
**Tokens** | [**[]GetAddressInfoResponseTokens**](getAddressInfoResponse_tokens.md) | Array of NTP1 tokens in this UTXO. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


