package main

import (
	"fmt"
	"net/http"

	_ "github.com/shaj13/libcache/fifo"

	"github.com/shaj13/go-guardian/v2/auth/strategies/ldap"
)

func main() {
	cfg := ldap.Config{
		BaseDN:       "cn=users,cn=accounts,dc=demo1,dc=freeipa,dc=org",
		BindDN:       "uid=admin,cn=users,cn=accounts,dc=demo1,dc=freeipa,dc=org",
		URL:          "ldap://ipa.demo1.freeipa.org",
		BindPassword: "Secret123",
		Filter:       "(uid=%s)",
	}

	r, _ := http.NewRequest("GET", "/", nil)
	//r.SetBasicAuth("employee", "Secret123")

	r.Header.Add("Authorization", "Basic ZW1wbG95ZWU6U2VjcmV0MTIz")

	info, err := ldap.New(&cfg).Authenticate(r.Context(), r)
	fmt.Println(info, err != nil)
	fmt.Println(err)

}
