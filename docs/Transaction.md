# Transaction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | The transaction amount in BTC | [optional] [default to null]
**Fee** | **float32** | The amount of the fee in BTC. This is negative and only available for the send category of transactions. | [optional] [default to null]
**Confirmations** | **int64** | The number of confirmations | [optional] [default to null]
**Blockhash** | **string** | The block hash | [optional] [default to null]
**Blockindex** | **int64** | The index of the transaction in the block that includes it | [optional] [default to null]
**Blocktime** | **int32** | The time in seconds since epoch (1 Jan 1970 GMT) | [optional] [default to null]
**Txid** | **string** | The transaction id | [optional] [default to null]
**Txhash** | **string** | The transaction hash | [optional] [default to null]
**Version** | **int32** |  | [optional] [default to null]
**Size** | **int32** |  | [optional] [default to null]
**Vsize** | **int32** |  | [optional] [default to null]
**Locktime** | **int32** |  | [optional] [default to null]
**Time** | **int32** | The transaction time in seconds since epoch (1 Jan 1970 GMT) | [optional] [default to null]
**Timereceived** | **int32** | The time received in seconds since epoch (1 Jan 1970 GMT) | [optional] [default to null]
**Bip125Replaceable** | **string** | Whether this transaction could be replaced due to BIP125 (replace-by-fee); may be unknown for unconfirmed transactions not in the mempool | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


