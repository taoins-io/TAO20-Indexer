# TAO-20 Indexer

Docs: https://taoins.io/tao20

## About TAO-20
Tao20 is a token standard based on TaoIns Protocol. As the inaugural test-level application of the TaoIns Protocol, inspired by BRC-20, we‘ve defined a TAO-20 standard for a Fungible Token on Bittensor. In this application:

Asset-type is specified as 0x2.

Content-type is defined as 0x0.

The 216 bits data-value is segmented into four parts:

Ticker occupies 32 bits, representing for the name of the fungible token asset.

Operation occupies 64 bits, representing for the users' execution of the fungible token asset, including deploy, mint, transfer and cancel

Amount occupies 48 bits.

AdditionalData occupies 72 bits. This field is used in deploy, transfer and cancel operation.

Special attention should be given to the fact that in the data-value. Ticker and Operation are parsed in ASCII format, while Amount and AdditionalData are parsed as hex numbers. In case of insufficient data for any field, uniform zero-bit padding is required. Additionally, it is crucial to enforce the restriction that ticker must not have leading zeros.

## About TaoIns Protocol

The TaoIns protocol is an innovative asset standard protocol on BitTensor. Since BitTensor doesnʼt have smart contracts or other type running environments for asset issuing, we intend to define an asset standard to embed the asset creations and operations and included the content into BitTensor blockchain.

The TaoIns protocol implements a technical design similar to Ordinals on Bitcoin, by engraving operational information into the receiving addresses of BitTensor transactions.

A BitTensor address consists a mutable 256-bits filed, i.e the PublicKey. In TaoIns protocol, this 256-bits space is divided into four segments: 32bits as the TaoIns indicator(0xffffffff), 4 bits for asset type, 4 bits for content type, and the remaining 216 bits designated for arbitrary data values. Essentially, each address represents unique information or operation. Any user wishing to perform a specific operation or convey a message related to a particular address simply needs to initiate a $TAO transaction to that address.

## About this repo
This repo contains the sample code for Tao20 token indexing.


## Setup
* Please note, these steps require the Go language version 1.21 or above.

1、Download the project to your local machine, modify the configuration files to suit your own settings(`./config/resources/config.toml`), and then run `go install` to install dependencies.

2、Run the `go build main.go` to create an executable file.

3、Run the `go run main` to start the service.
