package techpalace
import (
    "strings"
)

var welcome  string = "Welcome to the Tech Palace,"
var border string = "*"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    customerName := strings.ToUpper(customer)
	return welcome + " " + customerName
}

func AddNewline(text string) string{
    text +="\n"
    return text
}

func AddBorderPart(number int) string{
    return strings.Repeat(border, number)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	result := AddBorderPart(numStarsPerLine)
    result = AddNewline(result)
    result += welcomeMsg
    result = AddNewline(result)
    result += AddBorderPart(numStarsPerLine)
    return result    
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.Trim(oldMsg, "\n* ")
}
