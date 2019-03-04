/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

type RpcRequest struct {
	// JSON-RPC version
	Jsonrpc string `json:"jsonrpc"`
	// Identifier of RCP caller
	Id string `json:"id"`
	// Name of the Neblio RPC method to call
	Method string `json:"method"`
	// Array of string params that should be passed to the RPC method.
	Params []string `json:"params"`
}
