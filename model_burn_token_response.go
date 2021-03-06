/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

type BurnTokenResponse struct {
	// Unsigned, raw transaction hex of the transaction to burn the token
	TxHex string `json:"txHex,omitempty"`
	// Array of indexes transfering NTP1 tokens
	Ntp1OutputIndexes []float32 `json:"ntp1OutputIndexes,omitempty"`
	// Array of indexes of multisig outputs
	MultisigOutputs []float32 `json:"multisigOutputs,omitempty"`
}
