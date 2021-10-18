package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	emailVerifier "github.com/AfterShip/email-verifier"
	"github.com/julienschmidt/httprouter"
)

var (
	verifier = emailVerifier.
		NewVerifier().
		EnableSMTPCheck().
		Proxy("socks5://user:password@213.186.119.58:13660?timeout=5s")
)

func GetEmailVerification(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	domain := "outlook.com"
	username := "sangtran.d14ptit"
	ret, err := verifier.CheckSMTP(domain, username)
	if err != nil {
		fmt.Println("check smtp failed: ", err)
		return
	}

	fmt.Println("smtp validation result: ", ret)

}

func main() {
	router := httprouter.New()

	router.GET("/v1/:email/verification", GetEmailVerification)

	log.Fatal(http.ListenAndServe(":8080", router))
}
