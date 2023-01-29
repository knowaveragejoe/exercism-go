package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var message string = strings.Repeat("*", numStarsPerLine)
	message += "\n" + welcomeMsg + "\n"
	message += strings.Repeat("*", numStarsPerLine)

	return message
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var message string = strings.ReplaceAll(oldMsg, "*", "")

	message = strings.TrimSpace(message)

	return message
}
