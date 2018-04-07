/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type IssueTokenRequestMetadataEncryptions struct {

	// userData key to encrypt
	Key string `json:"key,omitempty"`

	// RSA public key used for encryption
	Pubkey string `json:"pubkey,omitempty"`

	// key format (pem or der)
	Format string `json:"format,omitempty"`

	// pkcs1 or pkcs8
	Type_ string `json:"type,omitempty"`
}
