using System;
using System.Windows.Forms;

namespace CalculatorApp
{
    public partial class Form1 : Form
    {
        private double firstNumber;
        private double secondNumber;
        private string currentOperation;
        private bool isOperationClicked;

        public Form1()
        {
            InitializeComponent();
        }

        private void buttonClick(object sender, EventArgs e)
        {
            Button button = (Button)sender;
            if (isOperationClicked)
            {
                textBoxResult.Clear();
                isOperationClicked = false;
            }

            textBoxResult.Text += button.Text;
        }

        private void operatorClick(object sender, EventArgs e)
        {
            Button button = (Button)sender;

            if (double.TryParse(textBoxResult.Text, out firstNumber))
            {
                currentOperation = button.Text;
                isOperationClicked = true;
            }
        }

        private void buttonEqualClick(object sender, EventArgs e)
        {
            if (double.TryParse(textBoxResult.Text, out secondNumber))
            {
                double result = 0;
                switch (currentOperation)
                {
                    case "+":
                        result = firstNumber + secondNumber;
                        break;
                    case "-":
                        result = firstNumber - secondNumber;
                        break;
                    case "*":
                        result = firstNumber * secondNumber;
                        break;
                    case "/":
                        if (secondNumber == 0)
                        {
                            MessageBox.Show("Cannot divide by zero.");
                            return;
                        }
                        result = firstNumber / secondNumber;
                        break;
                }

                textBoxResult.Text = result.ToString();
                firstNumber = result;
            }
        }

        private void buttonClearClick(object sender, EventArgs e)
        {
            textBoxResult.Clear();
            firstNumber = 0;
            secondNumber = 0;
            currentOperation = string.Empty;
            isOperationClicked = false;
        }
    }
}
