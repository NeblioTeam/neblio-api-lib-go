/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

type GetTransactionInfoResponse struct {
	// Transaction in raw hex
	Hex string `json:"hex,omitempty"`
	// TXID of transaction
	Txid string `json:"txid,omitempty"`
	// Transaction version
	Version float32 `json:"version,omitempty"`
	// Transaction locktime
	Locktime float32 `json:"locktime,omitempty"`
	// Array of transaction inputs
	Vin []GetTransactionInfoResponseVin `json:"vin,omitempty"`
	// Array of transaction outputs
	Vout []GetTransactionInfoResponseVout `json:"vout,omitempty"`
	// Block time of this transaction
	Blocktime float32 `json:"blocktime,omitempty"`
	// Block height of this transaction
	Blockheight float32 `json:"blockheight,omitempty"`
	// Total NEBL sent in this transaction in satoshis
	Totalsent float32 `json:"totalsent,omitempty"`
	// Total NEBL used as fee for this transcation in satoshis
	Fee float32 `json:"fee,omitempty"`
	// Hash of the block this transaction is in
	Blockhash string `json:"blockhash,omitempty"`
	// Transaction time
	Time float32 `json:"time,omitempty"`
	// Number of transaction confirmations
	Confirmations float32 `json:"confirmations,omitempty"`
}
