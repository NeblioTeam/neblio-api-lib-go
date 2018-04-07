/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BurnTokenResponse struct {

	// Unsigned, raw transaction hex of the transaction to burn the token
	TxHex string `json:"txHex,omitempty"`

	// Array of indexes transfering NTP1 tokens
	Ntp1OutputIndexes []float32 `json:"ntp1OutputIndexes,omitempty"`

	// Array of indexes of multisig outputs
	MultisigOutputs []float32 `json:"multisigOutputs,omitempty"`
}
