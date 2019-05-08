/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

type GetTokenHoldersResponse struct {
	// TokenId of the token
	TokenId string `json:"tokenId,omitempty"`
	Holders []GetTokenHoldersResponseHolders `json:"holders,omitempty"`
	// How many decimal points the token is divisble to
	Divibility float32 `json:"divibility,omitempty"`
	// Whether new issuances of this token are locked
	LockStatus bool `json:"lockStatus,omitempty"`
	// Whether the tokesn are aggregatable
	AggregationPolicy string `json:"aggregationPolicy,omitempty"`
	// A UTXO of this token
	SomeUtxo string `json:"someUtxo,omitempty"`
}
