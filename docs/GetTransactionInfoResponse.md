# GetTransactionInfoResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hex** | **string** | Transaction in raw hex | [optional] [default to null]
**Txid** | **string** | TXID of transaction | [optional] [default to null]
**Version** | **float32** | Transaction version | [optional] [default to null]
**Locktime** | **float32** | Transaction locktime | [optional] [default to null]
**Vin** | [**[]GetTransactionInfoResponseVin**](getTransactionInfoResponse_vin.md) | Array of transaction inputs | [optional] [default to null]
**Vout** | [**[]GetTransactionInfoResponseVout**](getTransactionInfoResponse_vout.md) | Array of transaction outputs | [optional] [default to null]
**Blocktime** | **float32** | Block time of this transaction | [optional] [default to null]
**Blockheight** | **float32** | Block height of this transaction | [optional] [default to null]
**Totalsent** | **float32** | Total NEBL sent in this transaction in satoshis | [optional] [default to null]
**Fee** | **float32** | Total NEBL used as fee for this transcation in satoshis | [optional] [default to null]
**Blockhash** | **string** | Hash of the block this transaction is in | [optional] [default to null]
**Time** | **float32** | Transaction time | [optional] [default to null]
**Confirmations** | **float32** | Number of transaction confirmations | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


