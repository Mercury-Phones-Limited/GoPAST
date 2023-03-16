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

	http.HandleFunc("/user-sip-post", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, startHTML)
		fmt.Fprintf(w, "&nbsp &nbsp &nbsp")
		fmt.Fprintf(w, "<a href=\"https://"+domainName+"/\" class=\"sipbutton\">Main Menu</a>")
		fmt.Fprintf(w, "&nbsp &nbsp &nbsp")
		fmt.Fprintf(w, "<a href=\"https://"+domainName+"/user-sip-detail\" class=\"sipbutton\">SIP Detail</a>")
		fmt.Fprintf(w, "&nbsp &nbsp &nbsp")
		fmt.Fprintf(w, "<a href=\"https://"+domainName+"/user-sip-registration\" class=\"sipbutton\">SIP Registration</a>")
		fmt.Fprintf(w, "&nbsp &nbsp &nbsp")
		fmt.Fprintf(w, "<a href=\"#\" onclick=\"window.close('https://"+domainName+"'); window.open('https://"+domainName+"/oauth2/sign_out'); window.open('https://github.com/logout');\" class=\"logoutbutton\"> Logout</a>")
		fmt.Fprintf(w, "&nbsp &nbsp &nbsp")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<table>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td><h3>&nbsp &nbsp New SIP &nbsp &nbsp</h3></td>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "</table>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<div class=newsip>")
		fmt.Fprintf(w, "<form method=\"POST\" action=\"/order-sip\">")
		fmt.Fprintf(w, "  <label><b>SIP username<br>(Required, No Spaces)</b></label>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <input name=\"username\" type=\"text\" value=\"\" />")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <label><b>IP address<br>(Required)</b></label>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <input name=\"ip\" type=\"text\" value=\"\" />")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <label for=\"subnet_prefix\"><b>Subnet prefix<br>(Required)</b></label>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "    <select id=\"subnet_prefix\" name=\"subnet_prefix\">")
		fmt.Fprintf(w, "      <option value=\"\">Please select an option</option>")
		fmt.Fprintf(w, "      <option value=\"32\">32</option>")
		fmt.Fprintf(w, "      <option value=\"31\">31</option>")
		fmt.Fprintf(w, "      <option value=\"29\">29</option>")
		fmt.Fprintf(w, "      <option value=\"28\">28</option>")
		fmt.Fprintf(w, "      <option value=\"27\">27</option>")
		fmt.Fprintf(w, "      <option value=\"128\">128</option>")
                fmt.Fprintf(w, "      <option value=\"127\">127</option>")
                fmt.Fprintf(w, "      <option value=\"126\">126</option>")
                fmt.Fprintf(w, "      <option value=\"125\">125</option>")
                fmt.Fprintf(w, "      <option value=\"124\">124</option>")
                fmt.Fprintf(w, "      <option value=\"123\">123</option>")
                fmt.Fprintf(w, "    </select>")
                fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <label><b>Note<br>(Required, No Spaces)</b></label>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <input name=\"note\" type=\"text\" value=\"\" />")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <input type=\"submit\" value=\"submit\" />")
		fmt.Fprintf(w, "</form>")
		fmt.Fprintf(w, "</div>")
		fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<hr class=\"roundedbar\">")
                fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<table>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td><h3>&nbsp &nbsp Update SIP &nbsp &nbsp")
		fmt.Fprintf(w, "    <br>")
		fmt.Fprintf(w, "    &nbsp &nbsp (Blank field will keep the value unchanged) &nbsp &nbsp</h3></td>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "</table>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<div class=updatesip>")
		fmt.Fprintf(w, "<form method=\"POST\" action=\"/user-\">")
		fmt.Fprintf(w, "  <label><b>SIP username<br>(Required)</b></label>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <input name=\"username\" type=\"text\" value=\"\" />")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <label><b>New IP address<br>(Optional)</b></label>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <input name=\"ip\" type=\"text\" value=\"\" />")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <label><b>New IP address subnet or prefix<br>(Optional)</b></label>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "    <select id=\"subnet_prefix\" name=\"subnet_prefix\">")
		fmt.Fprintf(w, "      <option value=\"\">Do not alter</option>")
		fmt.Fprintf(w, "      <option value=\"32\">32</option>")
		fmt.Fprintf(w, "      <option value=\"31\">31</option>")
		fmt.Fprintf(w, "      <option value=\"29\">29</option>")
		fmt.Fprintf(w, "      <option value=\"28\">28</option>")
		fmt.Fprintf(w, "      <option value=\"27\">27</option>")
		fmt.Fprintf(w, "      <option value=\"128\">128</option>")
                fmt.Fprintf(w, "      <option value=\"127\">127</option>")
                fmt.Fprintf(w, "      <option value=\"126\">126</option>")
                fmt.Fprintf(w, "      <option value=\"125\">125</option>")
                fmt.Fprintf(w, "      <option value=\"124\">124</option>")
                fmt.Fprintf(w, "      <option value=\"123\">123</option>")
                fmt.Fprintf(w, "    </select>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <label><b>Generate new password<br>(Optional)</b></label>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "    <select id=\"generate_password\" name=\"generate_password\">")
		fmt.Fprintf(w, "      <option value=\"\">Do not alter</option>")
		fmt.Fprintf(w, "      <option value=\"no\">no</option>")
		fmt.Fprintf(w, "      <option value=\"yes\">yes</option>")
                fmt.Fprintf(w, "    </select>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <label><b>Replace note<br>(Optional, No Spaces)</b></label>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <input name=\"note\" type=\"text\" value=\"\" />")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <input type=\"submit\" value=\"submit\" />")
		fmt.Fprintf(w, "</form>")
		fmt.Fprintf(w, "</div>")
		fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<hr class=\"roundedbar\">")
                fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<table>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td><h3>&nbsp &nbsp Delete SIP &nbsp &nbsp")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "</table>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<div class=deletesip>")
		fmt.Fprintf(w, "<form method=\"POST\" action=\"/user-\">")
		fmt.Fprintf(w, "  <label><b>SIP username to delete</b></label>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <input name=\"username\" type=\"text\" value=\"\" />")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <br>")
		fmt.Fprintf(w, "  <input type=\"submit\" value=\"submit\" />")
		fmt.Fprintf(w, "</form>")
		fmt.Fprintf(w, "</div>")
		fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, endHTML)
	})
	ipPort := "127.0.0.1:9003"
	fmt.Println("user_order_form_9003 is running on IP address and port " + ipPort)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(ipPort, nil))
}

// Contributor(s):
// Elliot Keavney - elliot@mercuryphones.co.uk
