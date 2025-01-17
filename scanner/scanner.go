/*
==========
Cariddi
==========

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see http://www.gnu.org/licenses/.

	@Repository:  https://github.com/edoardottt/cariddi

	@Author:      edoardottt, https://www.edoardoottavianelli.it
*/

package scanner

//isEmailUrl
func isEmailUrl(inp string) bool {
	return inp[:7] == "mailto:"
}

//isFtpUrl
func isFtpUrl(inp string) bool {
	return inp[:4] == "ftp:"
}

//isHttpUrl
func isHttpUrl(inp string) bool {
	return inp[:5] == "http:"
}

//isHttpsUrl
func isHttpsUrl(inp string) bool {
	return inp[:6] == "https:"
}
