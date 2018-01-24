package main

import (
	"log"

	"dbms/lib/ldap"
)

func main() {
	client := &ldap.LDAPClient{
		//Base:         "ou=神州租车用户与计算机,dc=zuche,dc=intra",
		//Base:         "dc=zuche,dc=com",
		Base:         "",
		Host:         "10.2.208.11",
		Port:         389,
		UseSSL:       false,
		BindDN:       "uid=maolin.yang01@ucarinc.com",
		BindPassword: "Welcome2sohu~",
		UserFilter:   "(uid=%s)",
		GroupFilter:  "(memberUid=%s)",
		Attributes:   []string{"givenName", "sn", "mail", "memberOf", "sAMAccountName"},
	}
	// It is the responsibility of the caller to close the connection
	defer client.Close()

	ierr := client.SimpleAuth("maolin.yang01@ucarinc.com", "Welcome2sohu~")
	if ierr != nil {
		log.Fatalf("auth err%v", ierr)
	} else {
		log.Printf("simple auth ok")
	}

	ok, user, err := client.Authenticate("maolin.yang01@ucarinc.com", "Welcome2sohu~")
	if err != nil {
		log.Fatalf("Error authenticating user %s: %+v", "maolin.yang01@ucarinc.com", err)
	}
	if !ok {
		log.Fatalf("Authenticating failed for user %s", "maolin.yang01@ucarinc.com")
	}
	log.Printf("!!!!!User: %+v", user)

	groups, err := client.GetGroupsOfUser("maolin.yang01@ucarinc.com")
	if err != nil {
		log.Fatalf("Error getting groups for user %s: %+v", "maolin.yang01@ucarinc.com", err)
	}
	log.Printf("!!!!Groups: %+v", groups)
}
