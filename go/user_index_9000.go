/*

MIT License

Copyright (c) 2023 Mercury Phones Limited

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package main

import (
        "cpresource"
        "fmt"
        "log"
        "net/http"
)

func main() {
        var startHTML string
        startHTML = cpresource.StartHTML()

        var endHTML string
        endHTML = cpresource.EndHTML()

        var domainName string
        domainName = cpresource.FQDN()
	
	var companyName string
        companyName = cpresource.CompanyName()
        
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
                fmt.Fprintf(w, startHTML)
                fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<table>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <th><h1>"+companyName+"</h1></th>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "</table>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<hr class=\"roundedbar\">")
		fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<div>")
                fmt.Fprintf(w, "<a href=\"https://"+domainName+"/user-sip-detail\" class=\"zsip\"><h2>SIP Details</h2></a> &nbsp &nbsp &nbsp")
                fmt.Fprintf(w, "<a href=\"https://"+domainName+"/user-sip-registration\" class=\"zsip\"><h2>SIP Registration</h2></a>  &nbsp &nbsp &nbsp")
                fmt.Fprintf(w, "<a href=\"https://"+domainName+"/user-alter-sip\" class=\"zsip\"><h2>Add / Alter SIP</h2></a> &nbsp &nbsp &nbsp")
                fmt.Fprintf(w, "</div>")
                fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<hr class=\"roundedbar\">")
                fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<div>")
                fmt.Fprintf(w, "<a href=\"https://"+domainName+"/user-number-route\" class=\"ynumber\"><h2>Number Route</h2></a> &nbsp &nbsp &nbsp")
                fmt.Fprintf(w, "<a href=\"https://"+domainName+"/user-alter-number\" class=\"ynumber\"><h2>Alter Number</h2></a>  &nbsp &nbsp &nbsp")
                fmt.Fprintf(w, "</div>")
                fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<hr class=\"roundedbar\">")
                fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<div>")
                fmt.Fprintf(w, "<a href=\"https://"+domainName+"/user-account-detail\" class=\"vaccount\"><h2>Account Details</h2></a> &nbsp &nbsp &nbsp")
                fmt.Fprintf(w, "<a href=\"https://"+domainName+"/user-alter-account\" class=\"vaccount\"><h2>Alter Account</h2></a>  &nbsp &nbsp &nbsp")
                fmt.Fprintf(w, "</div>")
                fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<hr class=\"roundedbar\"")
                fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<div>")
                fmt.Fprintf(w, "<a href=\"https://"+domainName+"/oauth2/sign_out\" class=\"wlogout\"><h2>Logout</h2></a> &nbsp &nbsp &nbsp")
                fmt.Fprintf(w, "</div>")
                fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, endHTML)

        })
        ipPort := "127.0.0.1:9000"
        fmt.Println("index9000 is running on IP address and port " + ipPort)
        // Start server on IP address and port specified above
        log.Fatal(http.ListenAndServe(ipPort, nil))
}

// Contributor(s):
// Elliot Keavney - elliot@mercuryphones.co.uk
