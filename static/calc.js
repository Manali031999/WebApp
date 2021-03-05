const calculator = {
  displayValue: '0',
  firstOperand: null,
  waitingForSecondOperand: false,
  operator: null,
};

function inputDigit(digit) {
  const { displayValue, waitingForSecondOperand } = calculator;

  if (waitingForSecondOperand === true) {
    calculator.displayValue = digit;
    console.log(calculator.displayValue);
    calculator.waitingForSecondOperand = false;
  } else {
    calculator.displayValue = displayValue === '0' ? digit : displayValue + digit;
    console.log(calculator.displayValue);
  }
}

function inputDecimal(dot) {
  // If the `displayValue` does not contain a decimal point
  if (!calculator.displayValue.includes(dot)) {
    // Append the decimal point
    calculator.displayValue += dot;
  }
}

function handleOperator(nextOperator) {
const { firstOperand, displayValue, operator } = calculator
const inputValue = parseFloat(displayValue);

if (operator && calculator.waitingForSecondOperand)  {
  calculator.operator = nextOperator;
  return;
}

if (firstOperand == null) {
  calculator.firstOperand = inputValue;
} else if (operator) {
    const currentValue = firstOperand || 0;
    console.log("Before Fetch")
    const result = Fetch(operator,currentValue,inputValue)
    result.then((data)=>{
      document.getElementById("displayresult").value = data.result;
      calculator.displayValue = String(data.result);
      calculator.firstOperand = data.result;
    })
}

calculator.waitingForSecondOperand = true;
calculator.operator = nextOperator;
}


async function Fetch(operator,currentValue,inputValue){
  const settings = {
    method: 'POST',
    body: JSON.stringify({
        "operator": operator,
        "op1": currentValue.toString(),
        "op2":inputValue.toString(),
        "result":"",
    }),
    headers: {
         Accept: 'application/json',
        'Content-Type': 'application/json',
    }
  };
  console.log("Fetch Called")
  const response = await fetch('/Get',settings)
  const data = await response.json();
  console.log(data)
  return data;
}

function resetCalculator() {
  calculator.displayValue = '0';
  calculator.firstOperand = null;
  calculator.waitingForSecondOperand = false;
  calculator.operator = null;
}

function updateDisplay() {
  const display = document.querySelector('.calc');
  display.value = calculator.displayValue;
}

updateDisplay();

const keys = document.querySelector('.calculator-keys');
keys.addEventListener('click', (event) => {
  const { target } = event;
  if (!target.matches('button')) {
    return;
  }

  if (target.classList.contains('operator')) {
    handleOperator(target.value);
    console.log(target.value);
    updateDisplay();
    return;
  }

  if (target.classList.contains('decimal')) {
    inputDecimal(target.value);
    updateDisplay();
    return;
  }

  if (target.classList.contains('all-clear')) {
    resetCalculator();
    updateDisplay();
    return;
  }

  inputDigit(target.value);
  updateDisplay();
});
