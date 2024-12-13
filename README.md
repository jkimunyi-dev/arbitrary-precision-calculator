# Arbitrary Precision Calculator

## Overview

The Arbitrary Precision Calculator is a powerful command-line tool and REPL
(Read-Eval-Print Loop) for performing mathematical operations on arbitrarily
large integers. It is designed for precision and flexibility, supporting a wide
range of operations without relying on libraries for its core functionality.

---

## Features

- **Arbitrary Precision**: Handle extremely large integers with precision.
- **Basic Arithmetic**: Addition, Subtraction, Multiplication, Division.
- **Advanced Operations**: Modulo, Exponentiation.
- **Interactive REPL**: Real-time calculations in an interactive terminal
  session.
- **Multiple Number Formats**: Supports Decimal, Binary, and Hexadecimal input.
- **Extensibility**: Easily extendable for additional operations.

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/jkimunyi-dev/arbitrary-precision-calculator.git
   ```

2. Navigate to the project directory:
   ```bash
   cd arbitrary-precision-calculator
   ```

3. Build the application:
   ```bash
   gcc *.c -o calc
   ```

---

## Usage

### Running the Calculator

The calculator can be used in two modes:

1. **Command-Line Mode**: Perform single calculations directly from the
   terminal.
2. **Interactive REPL Mode**: Enter multiple calculations interactively in a
   session.

### Command-Line Mode

To run a single calculation:

```bash
./calc [num1] [operation] [num2]
```

#### Supported Operations

| Operation      | Symbol | Description                                  |
| -------------- | ------ | -------------------------------------------- |
| Addition       | `+`    | Adds two numbers                             |
| Subtraction    | `-`    | Subtracts two numbers                        |
| Multiplication | `*`    | Multiplies two numbers                       |
| Division       | `/`    | Divides two numbers (quotient and remainder) |
| Modulo         | `%`    | Computes the remainder of division           |
| Exponentiation | `^`    | Raises a number to a power                   |

Examples:

```bash
./calc 1000 + 2000
./calc 2 ^ 10
```

### Interactive REPL Mode

To start the REPL mode, simply run the program without arguments:

```bash
./calc
```

In REPL mode, you can enter calculations interactively, such as:

```bash
1000 + 2000
5 ^ 3
exit
```

### Options

| Option  | Flag              | Description                 |
| ------- | ----------------- | --------------------------- |
| Help    | `--help`, `-h`    | Show usage instructions     |
| Version | `--version`, `-v` | Display version information |

Example:

```bash
./calc --help
./calc --version
```

---

## Documentation

- [**Architecture Overview**](./ARCHITECTURE.md): Detailed explanation of the
  project’s architecture.
- [**Implementation Details**](./IMPLEMENTATION.md): Comprehensive guide to the
  codebase and functionality.

---

## Known Limitations

- Currently only supports integers; floating-point numbers are not handled.
- Non-decimal bases (binary, hexadecimal) require specific input formats.

---

## Future Enhancements

- Add support for floating-point arithmetic.
- Implement additional operations like logarithms and trigonometric functions.
- Optimize performance for extremely large calculations.

---

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

---

## License

This project is licensed under the MIT License.
