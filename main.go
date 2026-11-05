package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// 故意保留可扫模式，便于 gosec / SARIF 验证（弱哈希 G401、硬编码 G101 等）
const demoToken = "demo-token-for-gosec-demo"

func main() {
	sum := md5.Sum([]byte("hello"))
	fmt.Println(hex.EncodeToString(sum[:]))
	_ = demoToken
}
