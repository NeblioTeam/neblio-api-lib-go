/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GetTokenMetadataResponseMetadataOfIssuanceData struct {

	// Token symbol
	TokenName string `json:"tokenName,omitempty"`

	// Name of token issuer
	Issuer string `json:"issuer,omitempty"`

	// Token description
	Description string `json:"description,omitempty"`

	UserData *GetTokenMetadataResponseMetadataOfIssuanceDataUserData `json:"userData,omitempty"`
}
