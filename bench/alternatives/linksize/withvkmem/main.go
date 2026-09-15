// Command withvkmem starts and stops one server: the smallest program that
// links vkmem, for the link-size delta.
package main

import "github.com/shibukawa/vkmem"

func main() {
	s, err := vkmem.Start()
	if err != nil {
		panic(err)
	}
	s.Close()
}
