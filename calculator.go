package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
	"strings"
	"strconv"
)

func main() {
	// Create a new application
	myApp := app.New()
	myWindow := myApp.NewWindow("Go Calculator")

	// Create a variable to hold the current expression
	var currentExpression string

	// Function to update the display
	updateDisplay := func() {
		display.SetText(currentExpression)
	}

	// Function to append character to the display
	appendToDisplay := func(value string) {
		currentExpression += value
		updateDisplay()
	}

	// Function to evaluate the current expression
	evaluateExpression := func() {
		// Remove any extra spaces
		expr := strings.TrimSpace(currentExpression)

		// Try to calculate the result using Go's `strconv` package (or some simple eval approach)
		result, err := evaluate(expr)
		if err != nil {
			currentExpression = "Error"
		} else {
			currentExpression = fmt.Sprintf("%f", result)
		}
		updateDisplay()
	}

	// Create the display label
	display := widget.NewLabel("0")

	// Create the number and operation buttons
	buttons := []string{
		"7", "8", "9", "/",
		"4", "5", "6", "*",
		"1", "2", "3", "-",
		"0", ".", "=", "+",
		"C",
	}

	// Create a grid layout for the buttons
	var buttonGrid []fyne.CanvasObject
	for _, btn := range buttons {
		btn := btn
		button := widget.NewButton(btn, func() {
			if btn == "=" {
				evaluateExpression()
			} else if btn == "C" {
				currentExpression = ""
				updateDisplay()
			} else {
				appendToDisplay(btn)
			}
		})
		buttonGrid = append(buttonGrid, button)
	}

	// Add the display and buttons to the window
	content := container.NewVBox(display, container.NewGridWithColumns(4, buttonGrid[0], buttonGrid[1], buttonGrid[2], buttonGrid[3],
		buttonGrid[4], buttonGrid[5], buttonGrid[6], buttonGrid[7],
		buttonGrid[8], buttonGrid[9], buttonGrid[10], buttonGrid[11],
		buttonGrid[12], buttonGrid[13], buttonGrid[14], buttonGrid[15]))
	myWindow.SetContent(content)

	// Set the window size and show
	myWindow.Resize(fyne.NewSize(300, 400))
	myWindow.ShowAndRun()
}

// Function to evaluate simple expressions (basic +, -, *, /)
func evaluate(expression string) (float64, error) {
	// Parse and evaluate the expression (simplified version without complex parsing)
	// You can use a library like `go-eval` for more complex evaluation
	// Here, we'll use strconv and try to handle simple float-based operations for demo purposes.

	// Here we can use a simple solution, but note that `strconv.ParseFloat` is not an expression parser.
	// For a more complete calculator, you would need a real expression parser.
	parts := strings.Fields(expression)
	if len(parts) == 3 {
		operand1, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return 0, err
		}
		operand2, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			return 0, err
		}
		operator := parts[1]

		var result float64
		switch operator {
		case "+":
			result = operand1 + operand2
		case "-":
			result = operand1 - operand2
		case "*":
			result = operand1 * operand2
		case "/":
			if operand2 == 0 {
				return 0, fmt.Errorf("cannot divide by zero")
			}
			result = operand1 / operand2
		default:
			return 0, fmt.Errorf("invalid operator")
		}

		return result, nil
	}

	return 0, fmt.Errorf("invalid expression")
}
