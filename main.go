package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
What is `go mod tidy` : add module reqs and sums?
*/

/*
Set up Aggregate Data type
Set var name in the json : var `title` accesses the give value
Needed for browser

Browser decodes JSON into browser-compatible info.
*/

/*
➜  ~ curl http:/localhost:5100
404 page not found
➜  ~ curl http:/localhost:5100/
404 page not found

Woah -> use HTTP : not HTTPS here

*/

// Can we do other data types here?
// gaaah commmify it too?
type NetWorth struct {
	Asset       string `json:"asset"`
	Liabilities string `json:"liabilities"`
	Equity      string `json:"equity"`
}

type Asset struct {
	Asset string `json:"asset"`
}

type Liability struct {
	Liabilities string `json:"liabilities"`
}

type Equity struct {
	Equity string `json:"equity"`
}

/*
How to even position the struct here?
*/
// HariNetWorth := &NetWorth{
//  Asset:       "200000",
//  Liabilities: "100000",
//  Equity:      "100000",
// }

/*
Why is httpRequest a pointer, but httpResponse an actual value
*/
func viewAsset(w http.ResponseWriter, r *http.Request) {
	HariAssets := &Asset{
		Asset: "200000",
	}
	// The f*** why do I have to specify if it's application/json?
	// Setting the type for JSON is important
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(HariAssets)

}

// /*
// It's almost like the server has to inform the client ... the browser ... on what to render too!
// */
func viewLiabilities(w http.ResponseWriter, r *http.Request) {
	HariLiability := &Liability{
		Liabilities: "100000",
	}
	// Canonicalized. WOW
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(HariLiability)

}

func viewEquities(w http.ResponseWriter, r *http.Request) {
	HariEquity := &Equity{
		Equity: "100000",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HariEquity)

}

/*
GAAAH in Enterprise Network can not accept network connections.
Must test on private LAN.
Do you want the application “main” to accept incoming network connections?
Local host no work on Enteprrise network :-)

curl http:/localhost:5050
curl : The remote name could not be resolved: 'http'

GAAAH powershell resolution
curl.exe http://localhost:5100
404 page not found
*/
func main() {

	// Set up our endpoints here ( can we do elsewhere? )
	// Pattern match endpoints
	// Dispatch pattern based on endpoint hit :-)
	// If empty method : is a page
	// We need an endpoint there ain't no endpoint-less webpages.
	// The `/` auto resolves away to no characters.
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte("<h1 style='color: steelblue'>THIS IS YOUR HOMEPAGE</h1>"))
		})
	http.HandleFunc("/assets", viewAsset)
	http.HandleFunc("/liabilities", viewLiabilities)
	http.HandleFunc("/equity", viewEquities)

	// os.Exit(1) : the OS needed to interface and crash a program deterministically?
	// At least the error is non-nil ( if err ) :-)
	// Wait no need to spec Localhost here!
	// Wait do we seriously need to render page to browser?
	log.Fatal(http.ListenAndServe(":5100", nil))

}
