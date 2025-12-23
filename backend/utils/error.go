package utils

import (
	"fmt"
)

type InvalidKeyLength struct {
	Length int
}

func (i *InvalidKeyLength) Error() string {
	return fmt.Sprintf("invalid key length. Key must be either 16, 24 or 32 bytes, Got: %d", i.Length)
}

type KeyNotFound struct {
}

func (i *KeyNotFound) Error() string {
	return "key not found"
}

type InvalidDstBlock struct {
	Dst int
	Src int
}

func (i *InvalidDstBlock) Error() string {
	return fmt.Sprintf("length of dst block is less than length of src block. Dst: %d bytes, Src: %d bytes", i.Dst, i.Src)
}

type InvalidPadding struct {
	Length    int
	BlockSize int
}

func (i *InvalidPadding) Error() string {
	return fmt.Sprintf("incorrect padding length. Padding must be between 1 and %d, Got: %d", i.BlockSize, i.Length)
}

type InvalidSrcBlock struct {
	Src int
	Block int
}

func (i *InvalidSrcBlock) Error() string {
	return fmt.Sprintf("length of src block must be multiple of block size. Src: %d bytes, Block Size: %d", i.Src, i.Block)
}
