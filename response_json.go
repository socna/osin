package osin

import (
	"encoding/json"
	"net/http"
	"fmt"
)

// OutputJSON encodes the Response to JSON and writes to the http.ResponseWriter
func OutputJSON(rs *Response, w http.ResponseWriter, r *http.Request) error {
	fmt.Println(r.URL.Query())
	
	// Add headers
	for i, k := range rs.Headers {
		for _, v := range k {
			w.Header().Add(i, v)
		}
	}
	
	fmt.Println(w.Header())

	if rs.Type == REDIRECT {
		// Output redirect with parameters
		u, err := rs.GetRedirectUrl()
		if err != nil {
			return err
		}
		fmt.Printf("redirect: %s", u)
		w.Header().Add("Location", u)
		w.WriteHeader(302)
	} else {
		// set content type if the response doesn't already have one associated with it
		if w.Header().Get("Content-Type") == "" {
			w.Header().Set("Content-Type", "application/json")
		}
		w.WriteHeader(rs.StatusCode)

		encoder := json.NewEncoder(w)
		fmt.Printf("redirect: %v", rs.Output)
		err := encoder.Encode(rs.Output)
		if err != nil {
			return err
		}
	}
	return nil
}
