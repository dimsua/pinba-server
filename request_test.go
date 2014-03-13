package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var valid_data99 = []byte{0xa, 0x8, 0x70, 0x68, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x34, 0x12, 0xb, 0x6c, 0x6f, 0x76, 0x65, 0x2e, 0x6e, 0x67, 0x73, 0x2e, 0x72, 0x75, 0x1a, 0x15, 0x4c, 0x6f, 0x76, 0x65, 0x5f, 0x50, 0x61, 0x67, 0x65, 0x5f, 0x4d, 0x73, 0x67, 0x3a, 0x3a, 0x5f, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x20, 0xc0, 0x1, 0x28, 0x9a, 0x9d, 0x3, 0x30, 0x80, 0x80, 0xd0, 0x1, 0x3d, 0x88, 0x80, 0x83, 0x3d, 0x45, 0xcd, 0xcc, 0x4c, 0x3d, 0x4d, 0x0, 0x0, 0x0, 0x0, 0x50, 0x1, 0x50, 0x6, 0x50, 0x1, 0x50, 0x1, 0x50, 0x2, 0x50, 0x2, 0x50, 0x1, 0x50, 0x1, 0x50, 0x3, 0x50, 0x1, 0x50, 0x1, 0x50, 0x1, 0x50, 0x2, 0x50, 0x1, 0x50, 0x1, 0x50, 0x1, 0x50, 0x1, 0x50, 0x2, 0x50, 0x1, 0x50, 0x1, 0x50, 0x9, 0x50, 0x2, 0x50, 0x1, 0x50, 0x1, 0x50, 0x1, 0x50, 0x1, 0x50, 0x1, 0x50, 0x1, 0x50, 0x3, 0x50, 0x1, 0x50, 0x3, 0x50, 0x1, 0x5d, 0x5f, 0xb2, 0xf1, 0x39, 0x5d, 0xf8, 0x55, 0x39, 0x3a, 0x5d, 0x79, 0xb0, 0xc5, 0x3b, 0x5d, 0x3b, 0xaa, 0x9a, 0x3a, 0x5d, 0xf4, 0xc0, 0xc7, 0x3a, 0x5d, 0x94, 0x17, 0x99, 0x3a, 0x5d, 0x37, 0x8c, 0x82, 0x3a, 0x5d, 0x20, 0xf0, 0x40, 0x38, 0x5d, 0xcf, 0x32, 0xb, 0x3b, 0x5d, 0x61, 0xc1, 0x7d, 0x39, 0x5d, 0xe, 0xf5, 0x3b, 0x3a, 0x5d, 0xc3, 0x64, 0x2a, 0x3a, 0x5d, 0x97, 0x38, 0x72, 0x3b, 0x5d, 0x7, 0x42, 0xb2, 0x39, 0x5d, 0x1a, 0xc3, 0x9c, 0x39, 0x5d, 0x16, 0x31, 0xec, 0x3a, 0x5d, 0x7, 0x42, 0xb2, 0x38, 0x5d, 0xec, 0xde, 0xa, 0x3b, 0x5d, 0x5f, 0xb6, 0x9d, 0x3c, 0x5d, 0x5, 0x33, 0x26, 0x3a, 0x5d, 0x3, 0x27, 0xdb, 0x3b, 0x5d, 0xb5, 0x8a, 0x7e, 0x3a, 0x5d, 0x1, 0x18, 0xcf, 0x39, 0x5d, 0xbe, 0xc0, 0x2c, 0x3a, 0x5d, 0xca, 0xe1, 0x13, 0x3c, 0x5d, 0xcd, 0xac, 0xa5, 0x3a, 0x5d, 0x7f, 0x87, 0xa2, 0x3a, 0x5d, 0xe8, 0x69, 0xc0, 0x39, 0x5d, 0x42, 0xce, 0xfb, 0x3c, 0x5d, 0x47, 0x4, 0x63, 0x3a, 0x5d, 0x8e, 0x76, 0x5c, 0x3a, 0x5d, 0x1, 0x18, 0xcf, 0x39, 0x60, 0x2, 0x60, 0x3, 0x60, 0x2, 0x60, 0x2, 0x60, 0x3, 0x60, 0x3, 0x60, 0x2, 0x60, 0x4, 0x60, 0x3, 0x60, 0x3, 0x60, 0x2, 0x60, 0x3, 0x60, 0x3, 0x60, 0x3, 0x60, 0x3, 0x60, 0x2, 0x60, 0x4, 0x60, 0x3, 0x60, 0x2, 0x60, 0x3, 0x60, 0x3, 0x60, 0x3, 0x60, 0x3, 0x60, 0x3, 0x60, 0x3, 0x60, 0x2, 0x60, 0x2, 0x60, 0x3, 0x60, 0x2, 0x60, 0x4, 0x60, 0x4, 0x60, 0x3, 0x68, 0x8, 0x68, 0xa, 0x68, 0x8, 0x68, 0xa, 0x68, 0xe, 0x68, 0x8, 0x68, 0xa, 0x68, 0x8, 0x68, 0xa, 0x68, 0x8, 0x68, 0xa, 0x68, 0xe, 0x68, 0x8, 0x68, 0xa, 0x68, 0x15, 0x68, 0x8, 0x68, 0xa, 0x68, 0x18, 0x68, 0x8, 0x68, 0xa, 0x68, 0x15, 0x68, 0x8, 0x68, 0x1c, 0x68, 0xa, 0x68, 0x8, 0x68, 0x1c, 0x68, 0xa, 0x68, 0x8, 0x68, 0xa, 0x68, 0x8, 0x68, 0x1c, 0x68, 0xa, 0x68, 0x8, 0x68, 0xa, 0x68, 0xe, 0x68, 0x8, 0x68, 0x1c, 0x68, 0xa, 0x68, 0x8, 0x68, 0x1c, 0x68, 0xa, 0x68, 0x8, 0x68, 0xa, 0x68, 0x18, 0x68, 0x8, 0x68, 0xa, 0x68, 0x15, 0x68, 0x8, 0x68, 0x1c, 0x68, 0xa, 0x68, 0x8, 0x68, 0xa, 0x68, 0x8, 0x68, 0x1c, 0x68, 0xa, 0x68, 0x8, 0x68, 0x1c, 0x68, 0xa, 0x68, 0x8, 0x68, 0x1c, 0x68, 0xa, 0x68, 0x8, 0x68, 0x1c, 0x68, 0xa, 0x68, 0x8, 0x68, 0x1c, 0x68, 0xa, 0x68, 0x8, 0x68, 0x1c, 0x68, 0xa, 0x68, 0x8, 0x68, 0xa, 0x68, 0x8, 0x68, 0xa, 0x68, 0x8, 0x68, 0x1c, 0x68, 0xa, 0x68, 0x8, 0x68, 0xa, 0x68, 0x18, 0x68, 0x8, 0x68, 0xa, 0x68, 0x15, 0x68, 0x18, 0x68, 0x33, 0x68, 0x8, 0x68, 0xa, 0x68, 0x8, 0x68, 0xa, 0x68, 0x15, 0x70, 0x9, 0x70, 0xb, 0x70, 0xc, 0x70, 0xd, 0x70, 0xf, 0x70, 0x10, 0x70, 0x11, 0x70, 0x12, 0x70, 0x13, 0x70, 0x14, 0x70, 0xd, 0x70, 0xf, 0x70, 0x9, 0x70, 0x14, 0x70, 0x16, 0x70, 0x12, 0x70, 0x17, 0x70, 0x19, 0x70, 0x9, 0x70, 0x1a, 0x70, 0x1b, 0x70, 0x1a, 0x70, 0x1d, 0x70, 0x1e, 0x70, 0x1a, 0x70, 0x1f, 0x70, 0x1e, 0x70, 0x10, 0x70, 0x20, 0x70, 0x1a, 0x70, 0x1f, 0x70, 0x21, 0x70, 0xc, 0x70, 0x22, 0x70, 0xf, 0x70, 0x1a, 0x70, 0x23, 0x70, 0x24, 0x70, 0x1a, 0x70, 0x23, 0x70, 0x1e, 0x70, 0x25, 0x70, 0x24, 0x70, 0x26, 0x70, 0x9, 0x70, 0x1a, 0x70, 0x27, 0x70, 0x1a, 0x70, 0x28, 0x70, 0x1e, 0x70, 0x25, 0x70, 0x29, 0x70, 0x1a, 0x70, 0x2a, 0x70, 0x1e, 0x70, 0x1a, 0x70, 0x2a, 0x70, 0x29, 0x70, 0x1a, 0x70, 0x2b, 0x70, 0x1e, 0x70, 0x1a, 0x70, 0x2a, 0x70, 0x24, 0x70, 0x1a, 0x70, 0x2b, 0x70, 0x29, 0x70, 0x1a, 0x70, 0x28, 0x70, 0x21, 0x70, 0x25, 0x70, 0x2c, 0x70, 0x25, 0x70, 0x2d, 0x70, 0x1a, 0x70, 0x2b, 0x70, 0x24, 0x70, 0x2e, 0x70, 0x2f, 0x70, 0x30, 0x70, 0x9, 0x70, 0x31, 0x70, 0x32, 0x70, 0x30, 0x70, 0x32, 0x70, 0x31, 0x70, 0x34, 0x70, 0x9, 0x70, 0x14, 0x70, 0x35, 0x7a, 0x2, 0x6e, 0x6f, 0x7a, 0x7, 0x69, 0x73, 0x5f, 0x61, 0x6a, 0x61, 0x78, 0x7a, 0x2, 0x35, 0x34, 0x7a, 0x6, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x7a, 0x7, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x7a, 0x4, 0x74, 0x79, 0x70, 0x65, 0x7a, 0x4, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x4, 0x75, 0x73, 0x65, 0x72, 0x7a, 0x5, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x7a, 0x7, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x7a, 0x9, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7a, 0x9, 0x6d, 0x65, 0x6d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x7a, 0x8, 0x6d, 0x65, 0x6d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x7a, 0x3, 0x67, 0x65, 0x74, 0x7a, 0x5, 0x63, 0x61, 0x63, 0x68, 0x65, 0x7a, 0x3, 0x68, 0x69, 0x74, 0x7a, 0x7, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7a, 0x9, 0x4c, 0x6f, 0x76, 0x65, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x7a, 0x8, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x7a, 0x5, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x7a, 0x5, 0x72, 0x65, 0x64, 0x69, 0x73, 0x7a, 0x6, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x7a, 0x13, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2d, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x36, 0x33, 0x38, 0x31, 0x7a, 0x7, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x7a, 0x8, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x7a, 0x4, 0x4c, 0x6f, 0x76, 0x65, 0x7a, 0x5, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x7a, 0x27, 0x31, 0x37, 0x32, 0x2e, 0x31, 0x36, 0x2e, 0x37, 0x36, 0x2e, 0x32, 0x33, 0x32, 0x3a, 0x32, 0x37, 0x30, 0x31, 0x37, 0x2c, 0x31, 0x37, 0x32, 0x2e, 0x31, 0x36, 0x2e, 0x37, 0x36, 0x2e, 0x32, 0x33, 0x31, 0x3a, 0x32, 0x37, 0x30, 0x31, 0x37, 0x7a, 0x2, 0x6e, 0x73, 0x7a, 0xa, 0x4c, 0x6f, 0x76, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x7a, 0x5, 0x71, 0x75, 0x65, 0x72, 0x79, 0x7a, 0xb, 0x4c, 0x6f, 0x76, 0x65, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x7a, 0x10, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x7a, 0x6, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x7a, 0x8, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x67, 0x65, 0x74, 0x7a, 0xc, 0x4c, 0x6f, 0x76, 0x65, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x7a, 0x5, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7a, 0x8, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x7a, 0x8, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x7a, 0x39, 0x31, 0x37, 0x32, 0x2e, 0x31, 0x36, 0x2e, 0x37, 0x36, 0x2e, 0x32, 0x33, 0x32, 0x3a, 0x32, 0x37, 0x30, 0x32, 0x32, 0x2c, 0x31, 0x37, 0x32, 0x2e, 0x31, 0x36, 0x2e, 0x37, 0x36, 0x2e, 0x34, 0x36, 0x3a, 0x32, 0x37, 0x30, 0x32, 0x32, 0x2c, 0x31, 0x37, 0x32, 0x2e, 0x31, 0x36, 0x2e, 0x37, 0x36, 0x2e, 0x34, 0x37, 0x3a, 0x32, 0x37, 0x30, 0x32, 0x32, 0x7a, 0x14, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x7a, 0x6, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x7a, 0x16, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x76, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x7a, 0x13, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x7a, 0xc, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x7a, 0x9, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x7a, 0x6, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x79, 0x7a, 0x5, 0x66, 0x65, 0x74, 0x63, 0x68, 0x7a, 0x7, 0x6e, 0x67, 0x73, 0x5f, 0x63, 0x6d, 0x73, 0x7a, 0x5, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x7a, 0xc, 0x64, 0x62, 0x72, 0x65, 0x61, 0x64, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x7a, 0x7, 0x64, 0x62, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x7a, 0x6, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x7a, 0xe, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x36, 0x33, 0x37, 0x39, 0x80, 0x1, 0xc8, 0x1, 0x88, 0x1, 0x80, 0xa0, 0xd7, 0x9, 0xa0, 0x1, 0x1, 0xa0, 0x1, 0x3, 0xa0, 0x1, 0x5, 0xa0, 0x1, 0x7, 0xa8, 0x1, 0x0, 0xa8, 0x1, 0x2, 0xa8, 0x1, 0x4, 0xa8, 0x1, 0x6, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0xa, 0xd7, 0x23, 0x3c, 0xb5, 0x1, 0xa, 0xd7, 0x23, 0x3c, 0xb5, 0x1, 0xa, 0xd7, 0x23, 0x3c, 0xb5, 0x1, 0xa, 0xd7, 0x23, 0x3c, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0xa, 0xd7, 0x23, 0x3c, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0xa, 0xd7, 0x23, 0x3c, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0xa, 0xd7, 0x23, 0x3c, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x1, 0x0, 0x0, 0x0, 0x0}
var invalid_data = []byte{0xa, 0x8, 0x70, 0x68, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x32, 0x12, 0xb, 0x6c, 0x6f, 0x76, 0x65, 0x2e, 0x6e, 0x67, 0x73, 0x2e, 0x72, 0x75, 0x1a, 0xc, 0x2f, 0x61, 0x64, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x68, 0x70, 0x20, 0xdc, 0x5, 0x28, 0x8b, 0x1, 0x30, 0x80, 0x80, 0x40, 0x3d, 0x4e, 0xe, 0x9f, 0x3b, 0x45, 0x0, 0x0, 0x0, 0x0, 0x4d, 0x0, 0x0, 0x0, 0x0, 0x50, 0x1, 0x50, 0x3, 0x50, 0x1, 0x50, 0x1, 0x5d, 0xb8, 0x93, 0x88, 0x3a, 0x5d, 0x1e, 0xe1, 0xb4, 0x3a, 0x5d, 0x20, 0xf0, 0x40, 0x39, 0x5d, 0xde, 0x21, 0x45, 0x39, 0x60, 0x2, 0x60, 0x2, 0x60, 0x2, 0x60, 0x2, 0x68, 0x8, 0x68, 0xa, 0x68, 0x8, 0x68, 0xa, 0x68, 0x8, 0x68, 0xa, 0x68, 0x8, 0x68, 0xa, 0x70, 0x9, 0x70, 0xb, 0x70, 0xc, 0x70, 0xd, 0x70, 0xc, 0x70, 0xe, 0x70, 0xf, 0x70, 0x10, 0x7a, 0x2, 0x6e, 0x6f, 0x7a, 0x7, 0x69, 0x73, 0x5f, 0x61, 0x6a, 0x61, 0x78, 0x7a, 0x2, 0x36, 0x36, 0x7a, 0x6, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x7a, 0x7, 0x61, 0x64, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x7a, 0x4, 0x74, 0x79, 0x70, 0x65, 0x7a, 0x5, 0x67, 0x75, 0x65, 0x73, 0x74, 0x7a, 0x4, 0x75, 0x73, 0x65, 0x72, 0x7a, 0x5, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x7a, 0x6, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x7a, 0x9, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7a, 0x4, 0x6c, 0x6f, 0x61, 0x64, 0x7a, 0x8, 0x6d, 0x65, 0x6d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x7a, 0x3, 0x67, 0x65, 0x74, 0x7a, 0x3, 0x61, 0x64, 0x64, 0x7a, 0x7, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7a, 0x5, 0x6d, 0x61, 0x67, 0x69, 0x71, 0x80, 0x1, 0x94, 0x3, 0x88, 0x1, 0x80, 0x80, 0xde, 0x2, 0xa0, 0x1, 0x1, 0xa0, 0x1, 0x3, 0xa0, 0x1, 0x5, 0xa0, 0x1, 0x7, 0xa8, 0x1, 0x0, 0xa8, 0x1, 0x2, 0xa8, 0x1, 0x4, 0xa8, 0x1, 0x6, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb5, 0x1, 0x0}

func TestCorrectDecoding(t *testing.T) {
	metrics, err := Decode(1394615846, valid_data99)
	if assert.Nil(t, err) {
		assert.Equal(t, metrics[len(metrics)-3].Name, "php.time")
		assert.Equal(t, metrics[len(metrics)-2].Name, "php.cpu")
		assert.Equal(t, metrics[len(metrics)-1].Name, "php.cnt")
	}
	assert.Equal(t, len(metrics), 99)
}

func TestIndexOutOfRange(t *testing.T) {
	_, err := Decode(1394615846, invalid_data)
	assert.NotNil(t, err)
}

func TestConvertToString(t *testing.T) {
	metrics, err := Decode(1394615846, valid_data99)
	assert.Nil(t, err)
	assert.Equal(t, len(metrics), 99)
	for idx, t := range metrics {
		fmt.Printf("#%d: %s", idx, t.String())
		break
	}
}

func BenchmarkDecode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Decode(time.Now().Unix(), valid_data99)
	}
}
