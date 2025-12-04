package main

import (
	"fmt"
	"strings"
)

func main() {
	s1, s2, s3 := "slo", "weslo", "osloslo"
	fmt.Println(strings.Index(s2, s1), "index substring")
	fmt.Println(strings.Contains(s2, s1), "contains substring")

	fmt.Println(strings.Split(s3, "l"))
	parts := []string{s1, s2, s3}
	fmt.Println(strings.Join(parts, " "))
	fmt.Println(strings.Replace(s2, s1, "replacement", -1))
	fmt.Println(s3[len(s3)-1], "byte")
	fmt.Println(string(s3[len(s3)-1]))
	fmt.Println(s3[:len(s3)-1], "splice [start:end]")
}
