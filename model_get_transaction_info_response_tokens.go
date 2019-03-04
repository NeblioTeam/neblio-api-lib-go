/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

type GetTransactionInfoResponseTokens struct {
	// ID of the token
	TokenId string `json:"tokenId,omitempty"`
	// Number of tokens
	Amount float32 `json:"amount,omitempty"`
	// TXID the token was issued in
	IssueTxid string `json:"issueTxid,omitempty"`
	// Decimal places the token is divisible to
	Divisibility float32 `json:"divisibility,omitempty"`
	// Whether issuance of more tokens is locked
	LockStatus bool `json:"lockStatus,omitempty"`
	// Whether the tokens are aggregatable
	AggregationPolicy string `json:"aggregationPolicy,omitempty"`
}
