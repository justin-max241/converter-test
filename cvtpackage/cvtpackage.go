package cvtpackage

import (
	"fmt"
	"regexp"
	"strings"
)

func ConvertToCustomList(html string) string {
	// Define a regular expression to match the content of <li> elements
	re := regexp.MustCompile(`<li>(.*?)</li>`)
	matches := re.FindAllStringSubmatch(html, -1)

	// Extract the content of each <li> element
	var items []string
	for _, match := range matches {
		if len(match) > 1 {
			items = append(items, "* "+match[1]+"\\\\")
		}
	}

	// Join the items with a newline as the separator
	return strings.Join(items, "\n")
}

func ConvertToHTML(input string) string {
	// Split the input string by the delimiter "\\"
	items := strings.Split(input, "\\")

	// Trim whitespace and remove empty strings
	var trimmedItems []string
	for _, item := range items {
		trimmedItem := strings.TrimSpace(item)
		if len(trimmedItem) > 0 {
			// Remove the leading '* ' from each item
			if strings.HasPrefix(trimmedItem, "* ") {
				trimmedItem = strings.TrimPrefix(trimmedItem, "* ")
			}
			trimmedItems = append(trimmedItems, trimmedItem)
		}
	}

	// Start building the HTML list
	var builder strings.Builder
	builder.WriteString("<ul>\n")
	for _, item := range trimmedItems {
		builder.WriteString("  <li>")
		builder.WriteString(item)
		builder.WriteString("</li>\n")
	}
	builder.WriteString("</ul>")

	return builder.String()
}

// Seperiert die markierten Worte vom restlichen Text
func Seperate(t string, a []string, b []string, o string) [][]string {
	// String an den definierten Stellen zerteilen (String wird zum Slice)
	baseSlice := strings.Split(t, o)

	// Iteriert durch den baseSlice und baut zwei Arrays (Markierte Wörter, Normaler Text) durch Modulo
	for i := 0; i < len(baseSlice); i++ {
		w := baseSlice[i]
		if i%2 == 0 {
			a = append(a, w)
		} else if i%2 == 1 {
			b = append(b, "<br>"+w+"</br>")
		} else {
			fmt.Println("Error!")
		}
	}

	// Beide Slices zu einem 2D-Slice kombinieren für den Rückgabewert
	var returnValue = [][]string{a, b}
	returnValue = append(returnValue, a, b)

	return returnValue
}

// Beide Slices wieder zu einem großen Slice zusammenfügen
func SliceConnect(s [][]string) []string {
	connectedSlice := []string{}

	// Durch beide Slices iterieren und abwechselnd zum verbundenen Slice hinzufügen
	for i := 0; i < len(s[0]); i++ {
		connectedSlice = append(connectedSlice, s[0][i])
		if i < len(s[1]) {
			connectedSlice = append(connectedSlice, s[1][i])
		}
	}

	return connectedSlice
}
