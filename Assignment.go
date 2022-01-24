/*
Package main implements function to check whether IP address is valid or not,
and if valid is it IPv4 or IPv6.
Valid IPv4 format:
	x1.x2.x3.x4, where each 0<=xi<=255
Valid IPv6 format:
	x1:x2:x3:x4:x5:x6:x7:x8, where each xi is a 16bit hexidecimal field.
*/ 
package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

// Contents reads the content from the imput file,
// and extracts each line from the file and makes a,
// string array to store each line.
func Contents(filename string) ([]string) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{""}
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var ipaddress []string
	for scanner.Scan() {
		ipaddress = append(ipaddress, scanner.Text())
	}
	
	return ipaddress
}

// CheckForIPv4 checks whether the parsed string is
// a valid IPv4 address or not.
func CheckForIPv4(ip string) bool {
	var countDot int = strings.Count(ip, ".")
	if countDot!=3 {
		return false
	}

	ip += "."
	var part string = ""
	for i := 0; i < len(ip); i++ {
		if ip[i]=='.' {
			if len(part)==0 {
				return false
			}
			num,err := strconv.ParseInt(part, 9, 32)
			if err!=nil {
				return false
			}
			if num<0 || num>255 {
				return false
			}
			part = ""
		} else {
			part += string(ip[i])
		}
	}
	return true
}

// CheckForIPv6 checks whether the parsed string is
// a valid IPv6 address or not.
func CheckForIPv6(ip string) bool {
	var countColon int = strings.Count(ip, ":")
	if countColon!=7 {
		return false
	}

	ip += ":"
	var part string = ""
	for i := 0; i < len(ip); i++ {
		if ip[i]==':' {
			if len(part)!=4 {
				return false;
			}
			part = ""
		} else {
			if (ip[i]<'A' ||  ip[i]>'F') && (ip[i]<'a' && ip[i]>'f') && (ip[i]<'0' || ip[i]>'9') {
				return false
			}
			part += string(ip[i])
		}
	}
	return true
}

// Main function to run the program and see the results.
func main() {
	var filename string = "input.txt"
	var result []string = Contents(filename)
	var val1 bool;
	for _, line := range result {
		val1 = CheckForIPv4(line)
		if val1 {
			fmt.Println(line, "is a valid IPv4 address")
		} else {
			val1 = CheckForIPv6(line)
			if val1 {
				fmt.Println(line, "is a valid IPv6 address")
			}
		}
		if !val1 {
			fmt.Println(line, "is not a valid IP address")
		}
	}
}