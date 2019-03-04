# GetTransactionInfoResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hex** | **string** | Transaction in raw hex | [optional] 
**Txid** | **string** | TXID of transaction | [optional] 
**Version** | **float32** | Transaction version | [optional] 
**Locktime** | **float32** | Transaction locktime | [optional] 
**Vin** | [**[]GetTransactionInfoResponseVin**](getTransactionInfoResponse_vin.md) | Array of transaction inputs | [optional] 
**Vout** | [**[]GetTransactionInfoResponseVout**](getTransactionInfoResponse_vout.md) | Array of transaction outputs | [optional] 
**Blocktime** | **float32** | Block time of this transaction | [optional] 
**Blockheight** | **float32** | Block height of this transaction | [optional] 
**Totalsent** | **float32** | Total NEBL sent in this transaction in satoshis | [optional] 
**Fee** | **float32** | Total NEBL used as fee for this transcation in satoshis | [optional] 
**Blockhash** | **string** | Hash of the block this transaction is in | [optional] 
**Time** | **float32** | Transaction time | [optional] 
**Confirmations** | **float32** | Number of transaction confirmations | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


