# security

[![Go Report Card](https://goreportcard.com/badge/github.com/datewu/security?style=flat-square)](https://goreportcard.com/report/github.com/datewu/security)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/datewu/security)

## Description
This package is made for dealing with security stuff like hashing, encryption, x509/pem marshal/unmaral, etc.
It supports Md5/SHA512 hashing, x509 certificate encoding,AES-GCM 256-bit symmetric-key en/decryption etc.

## Usage
Import as a normal go package
```golang
import  "github.com/datewu/security"

```
## Example
Check these two `encrypt_test.go` and `hash_test.go` files.
