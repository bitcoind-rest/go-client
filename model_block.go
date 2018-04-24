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

type Block struct {
	// The block hash
	Hash string `json:"hash,omitempty"`
	// The number of confirmations
	Confirmations int64 `json:"confirmations,omitempty"`
	// The block stripped size
	Strippedsize int64 `json:"strippedsize,omitempty"`
	// The block size
	Size int64 `json:"size,omitempty"`
	// The block weight
	Weight int64 `json:"weight,omitempty"`
	// The block height (or index)
	Height int64 `json:"height,omitempty"`
	// The block version
	Version int64 `json:"version,omitempty"`
	// The block version (in hex)
	VersionHex string `json:"versionHex,omitempty"`
	// The block merkle root
	Merkleroot string `json:"merkleroot,omitempty"`
	// The list of transactions in the block
	Tx []Transaction `json:"tx,omitempty"`
	// The block time
	Time int32 `json:"time,omitempty"`
	// The block median time
	Mediantime int32 `json:"mediantime,omitempty"`
	// The block nonce
	Nonce int32 `json:"nonce,omitempty"`
	Bits string `json:"bits,omitempty"`
	Difficulty float32 `json:"difficulty,omitempty"`
	Chainwork string `json:"chainwork,omitempty"`
}
