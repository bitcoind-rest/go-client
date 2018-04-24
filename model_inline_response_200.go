/*
 * Bitcoind
 *
 * The REST API can be enabled with the `-rest` option. The interface runs on the same port as the JSON-RPC interface, by default port `8332` for **mainnet**, port `18332` for **testnet**, and port `18443` for **regtest**.
 *
 * API version: 0.16
 * Contact: johan@lepetitbloc.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type InlineResponse200 struct {
	ChainHeight int32 `json:"chainHeight,omitempty"`
	ChaintipHash string `json:"chaintipHash,omitempty"`
	Bitmap string `json:"bitmap,omitempty"`
	Utxos []UTxO `json:"utxos,omitempty"`
}