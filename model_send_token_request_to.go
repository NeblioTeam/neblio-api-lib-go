/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

type SendTokenRequestTo struct {
	// Address to transfer tokens to
	Address string `json:"address,omitempty"`
	// Number of tokens to send
	Amount float32 `json:"amount,omitempty"`
	// ID of token we are sending
	TokenId string `json:"tokenId,omitempty"`
}
