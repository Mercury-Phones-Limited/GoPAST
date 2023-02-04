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

	http.HandleFunc("/sip-trunk-detail", func(w http.ResponseWriter, r *http.Request) {

		var accountCode string
		accountCode = string(r.Host)
		
		// Open database connection.
		db, err := sql.Open("mysql", dbDetails)
		defer db.Close()

		// Handle error
		if err != nil {
			panic("Is the database on?")
		}

		// SQL query returns all
		result, err := db.Query("SELECT ps_auth.username, ps_auth.password FROM ps_auth WHERE ps_auth.username AND ps_endpoint.username = ?", accountCode)
		defer db.Close()

		// Handle error
		if err != nil {
			panic("Is the database on?")
		}

		fmt.Fprintf(w, startHTML)
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<table>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td><h1>(Company Name)</h1></td>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td><h2>SIP Trunk Details For Account "+accountCode+"</h2></td>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "</table>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<table>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td><b>SIP USERNAME</b></td>")
		fmt.Fprintf(w, "    <td><b>SIP PASSWORD</b></td>")
		fmt.Fprintf(w, "    <td><b>RESTRICTED TO IP</b></td>")
		fmt.Fprintf(w, "  </tr>")
		for result.Next() {
			var username string
			var password string
			var ip string
			
			err = result.Scan(&username, &password, &ip)

			// Handle error
			if err != nil {
				panic("")
			}
			fmt.Fprintf(w, "  <tr>")
			fmt.Fprintf(w, "    <td>"+username+"</td>")
			fmt.Fprintf(w, "    <td>"+password+"</td>")
			fmt.Fprintf(w, "    <td>"+ip+"</td>")
			fmt.Fprintf(w, "  </tr>")
		}
		fmt.Fprintf(w, "</table>")
		fmt.Fprintf(w, endHTML)
	})
	ipPort := "127.0.0.1:9001"
	fmt.Println("siptrunkdetail9001 is running on IP address and port " + ipPort)
	// Start server on IP address and port specified above
	log.Fatal(http.ListenAndServe(ipPort, nil))
}
