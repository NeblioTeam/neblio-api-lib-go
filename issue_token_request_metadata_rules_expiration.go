/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Object describing expiration rules of the token
type IssueTokenRequestMetadataRulesExpiration struct {

	// Blockheight at wh
	ValidUntil float32 `json:"validUntil,omitempty"`

	// Whether this rule can be modified in future transactions
	Locked bool `json:"locked,omitempty"`
}