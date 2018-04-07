# GetAddressInfoResponseUtxos

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **float32** | Index of the UTXO at this address | [optional] [default to null]
**Txid** | **string** | Txid of this UTXO | [optional] [default to null]
**Blockheight** | **float32** | Blockheight of the UTXO | [optional] [default to null]
**Blocktime** | **float32** | Blocktime of the UTXO | [optional] [default to null]
**ScriptPubKey** | [***interface{}**](interface{}.md) | Object representing the scruptPubKey of the UTXO | [optional] [default to null]
**Used** | **bool** | Whether the UTXO has been used | [optional] [default to null]
**Value** | **float32** | Value of the UTXO in NEBL satoshi | [optional] [default to null]
**Tokens** | [**[]GetAddressInfoResponseTokens**](getAddressInfoResponse_tokens.md) | Array of NTP1 tokens in this UTXO. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


