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
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	var dbDetails string
	dbDetails = cpresource.DBReadDetails()

	var startHTML string
	startHTML = cpresource.StartHTML()

	var endHTML string
	endHTML = cpresource.EndHTML()

	var domainName string
	domainName = cpresource.FQDN()

	http.HandleFunc("/user-sip-registration", func(w http.ResponseWriter, r *http.Request) {

		var emailAddress string
		emailAddress = string(r.Header.Get("X-Forwarded-Email"))

		// Open a database connection.
		db, err := sql.Open("mysql", dbDetails)
		defer db.Close()

		// Handle error
		if err != nil {
			panic("Is the database on?")
		}
		
		// SQL query returns account number based on email address
		accountNumberResult, err := db.Query("SELECT account_number FROM account_number_lookup WHERE email_address = ?", emailAddress)
		defer db.Close()

		// Handle error
		if err != nil {
			panic("SQL query for account number not working")
		}

		for accountNumberResult.Next() {
			var accountNumber string
			err = accountNumberResult.Scan(&accountNumber)

			// Handle error
			if err != nil {
				panic("")
			}

		result, err := db.Query("SELECT endpoint, via_addr, via_port, user_agent FROM ps_contact, ps_endpoint WHERE ps_contact.endpoint = ps_endpoint.id AND ps_endpoint.customer_id = ?", accountNumber)
		defer db.Close()

		// Handle error
		if err != nil {
			panic("SQL query for data not working")
		}

		fmt.Fprintf(w, startHTML)
		fmt.Fprintf(w, "&nbsp &nbsp &nbsp")
		fmt.Fprintf(w, "<a href=\"https://"+domainName+"/\" class=\"sipbutton\">Main Menu</a>")
                fmt.Fprintf(w, "&nbsp &nbsp &nbsp")
		fmt.Fprintf(w, "<a href=\"https://"+domainName+"/user-sip-detail\" class=\"sipbutton\">SIP Detail</a>")
                fmt.Fprintf(w, "&nbsp &nbsp &nbsp")
		fmt.Fprintf(w, "<a href=\"https://"+domainName+"/user-sip-post\" class=\"addalterbutton\">Add / Alter SIP</a>")
		fmt.Fprintf(w, "&nbsp &nbsp &nbsp")
		fmt.Fprintf(w, "<a href=\"#\" onclick=\"window.close('https://"+domainName+"'); window.open('https://"+domainName+"/oauth2/sign_out'); window.open('https://github.com/logout');\" class=\"logoutbutton\"> Logout</a>")
		fmt.Fprintf(w, "&nbsp &nbsp &nbsp")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<table>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <th><h3>&nbsp &nbsp SIP Registration For Account "+accountNumber+" &nbsp &nbsp</h3></th>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <th><h3>&nbsp &nbsp user logged in is "+emailAddress+" &nbsp &nbsp</h3></th>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "</table>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<input type=\"text\" id=\"sipInput\" onkeyup=\"tableFunction()\" placeholder=\"Search for SIP username...\" title=\"Type in a sip username\">")
                fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<table>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <th><b>SIP USERNAME</b></th>")
		fmt.Fprintf(w, "    <th><b>IP ADDRESS</b></th>")
		fmt.Fprintf(w, "    <th><b>PORT</b></th>")
		fmt.Fprintf(w, "    <th><b>USER AGENT (PBX)</b></th>")
		fmt.Fprintf(w, "  </tr>")
		for result.Next() {
			var sip_username string
			var ip_addr string
			var port string
			var user_agent string

			err = result.Scan(&sip_username, &ip_addr, &port, &user_agent)

			// Handle error
			if err != nil {
				panic("")
			}
			fmt.Fprintf(w, "  <tr>")
			fmt.Fprintf(w, "    <td>"+sip_username+"</td>")
			fmt.Fprintf(w, "    <td>"+ip_addr+"</td>")
			fmt.Fprintf(w, "    <td>"+port+"</td>")
			fmt.Fprintf(w, "    <td>"+user_agent+"</td>")
			fmt.Fprintf(w, "  </tr>")
		}
		}
		fmt.Fprintf(w, "</table>")
		fmt.Fprintf(w, "<script>")
                fmt.Fprintf(w, "function tableFunction() {")
                fmt.Fprintf(w, "  var input, filter, table, tr, td, i, txtValue;")
                fmt.Fprintf(w, "  input = document.getElementById(\"sipInput\");")
                fmt.Fprintf(w, "  filter = input.value.toUpperCase();")
                fmt.Fprintf(w, "  table = document.getElementById(\"sipTable\");")
                fmt.Fprintf(w, "  tr = table.getElementsByTagName(\"tr\");")
                fmt.Fprintf(w, "  for (i = 0; i < tr.length; i++) {")
                fmt.Fprintf(w, "    td = tr[i].getElementsByTagName(\"td\")[0];")
                fmt.Fprintf(w, "    if (td) {")
                fmt.Fprintf(w, "      txtValue = td.textContent || td.innerText;")
                fmt.Fprintf(w, "      if (txtValue.toUpperCase().indexOf(filter) > -1) {")
                fmt.Fprintf(w, "        tr[i].style.display = \"\";")
                fmt.Fprintf(w, "      } else {")
                fmt.Fprintf(w, "        tr[i].style.display = \"none\";")
                fmt.Fprintf(w, "      }")
                fmt.Fprintf(w, "    }")
                fmt.Fprintf(w, "  }")
                fmt.Fprintf(w, "}")
                fmt.Fprintf(w, "</script>")
		fmt.Fprintf(w, endHTML)
	})
	ipPort := "127.0.0.1:9002"
	fmt.Println("sipregistration9002 is running on IP address and port " + ipPort)
	// Start server on port specified above
	log.Fatal(http.ListenAndServe(ipPort, nil))
}

// Contributor(s):
// Elliot Keavney - elliot@mercuryphones.co.uk
