package main

import (
	"fmt"
	"net/url"
)

func main() {

	line := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(line)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}
