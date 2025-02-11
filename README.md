# Monkey Interpreter in Go

Welcome to the **Monkey Programming Language Interpreter**! This project is an implementation of the Monkey programming language, built from scratch in Go, inspired by the book ["Writing An Interpreter In Go"](https://interpreterbook.com/). This language was designed specifically for the book and showcases a clean and beginner-friendly approach to understanding interpreters.

## Features

The Monkey programming language includes the following features:

- **Variables and Statements**: Use `let` statements to bind values to names.
- **Expressions**: Monkey supports arithmetic expressions and boolean logic.
- **Data Types**: Includes integers, booleans, strings, arrays, and hashes.
- **Functions**: First-class citizens with support for higher-order functions and closures.
- **Control Flow**: Includes `if` expressions and recursion.

Here’s a taste of Monkey with separate instructions on each line:

```monkey
let version = 1;
let name = "Monkey programming language";
let myArray = [1, 2, 3, 4, 5];
let coolBooleanLiteral = true;
let awesomeValue = (10 / 2) * 5 + 30;
```

Monkey also supports defining functions, recursion, and higher-order functions. Here, each line is a separate instruction:

```monkey
let fibonacci = fn(x) { if (x == 0) { 0 } else { if (x == 1) { return 1 } else { fibonacci(x - 1) + fibonacci(x - 2) } } };
let result = fibonacci(myArray[4]);
```

## Key Concepts

1. **First-Class Functions**: Functions in Monkey are first-class citizens. You can assign them to variables, pass them as arguments, and return them from other functions.

```monkey
let newGreeter = fn(greeting) { fn(name) { puts(greeting + " " + name) } };
let hello = newGreeter("Hello");
hello("world!");
```

2. **Higher-Order Functions**: You can pass functions as arguments and return them from other functions. Here’s an example using `map`:

```monkey
let map = fn(arr, f) { let iter = fn(arr, accumulated) { if (len(arr) == 0) { accumulated } else { iter(rest(arr), push(accumulated, f(first(arr)))) } }; iter(arr, []) };
let result = map([1, 2, 3], fn(x) { x * 2 });
```

## Getting Started

To run the Monkey interpreter, make sure you have Go installed. Clone this repository and build the project:

```bash
git clone https://github.com/shotowon/Monkey-Interpreter.git
cd Monkey-Interpreter
go build ./cmd/interpreter
./monkey
```

Once you have the interpreter running, you can type Monkey code directly into the prompt:

```bash
>> let x = 5;
>> let y = x * 2;
>> y;
10
```

## Features to Explore

- **Arithmetic operations**: `+`, `-`, `*`, `/`
- **Boolean operations**: `==`, `!=`, `<`, `>`
- **Array manipulation**: Indexing and operations like `len()`, `push()`, `first()`, `rest()`
- **Hash (Dictionary-like structures)**: Key-value pairs with string keys
- **Functions**: Anonymous functions, recursion, and closures
- **Control flow**: `if`, `else`, and return statements

## Example Code

Here’s a more complex example that shows Monkey's ability to handle recursion, arrays, and higher-order functions with all instructions on separate lines:

```monkey
let fibonacci = fn(x) { if (x == 0) { 0 } else { if (x == 1) { return 1 } else { fibonacci(x - 1) + fibonacci(x - 2) } } };
let numbers = [1, 2, 3, 4, 5];
let result = fibonacci(numbers[4]);
```