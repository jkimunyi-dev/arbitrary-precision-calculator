# ARCHITECTURE.md

## Project Architecture Overview

The **Arbitrary Precision Calculator** is designed with modularity, scalability,
and clarity in mind. The application is structured to separate concerns,
ensuring the core functionality is clean, reusable, and easily testable.

---

## Key Components

### 1. **Main Application**

- **File:** `main.go`
- **Responsibilities:**
  - Handles command-line arguments.
  - Manages the REPL (Read-Eval-Print Loop) initialization.
  - Interfaces with the core arithmetic operations.
  - Displays help and version information.

---

### 2. **REPL (Read-Eval-Print Loop)**

- **File:** `cmd/repl/repl.go`
- **Responsibilities:**
  - Provides an interactive interface for the user.
  - Parses user input for commands and operations.
  - Executes arithmetic calculations using core methods.
  - Displays results or errors in a user-friendly manner.

---

### 3. **Arithmetic Core**

- **File:** `internal/arithmetic.go`
- **Responsibilities:**
  - Implements core functionality for arbitrary-precision integer calculations.
  - Handles operations like addition, subtraction, multiplication, division,
    modulo, exponentiation, and factorial.
  - Validates and processes large integers as strings.
  - Ensures accurate results without using external libraries.

---

### 4. **Utilities**

- **File:** `internal/advanced_operation.go`
- **Responsibilities:**
  - Contains helper functions for parsing and validating input.
  - Converts between different numerical bases (binary, hexadecimal, etc.).
  - Handles error formatting and reporting.

---

### 5. **Constants**

- **File:** `internal/constants.go`
- **Responsibilities:**
  - Stores constant values used throughout the application, such as operation
    symbols and error messages.

---

---

### 5. **Tests**

- **File:** `internal/unit_test.go`
- **Responsibilities:**
  - Performs unit tests

- **File:** `internal/advanced_operations_test.go`
- **Responsibilities:**
  - Performs tests for `bash advanced_operation.go`

---

## Data Flow

1. **Input Handling:**
   - Command-line arguments or REPL input are parsed to extract operands and
     operations.
   - Input is validated for correctness (e.g., numeric format, valid
     operations).

2. **Operation Execution:**
   - Parsed input is passed to the core arithmetic methods.
   - Calculations are performed using string-based arithmetic for precision.

3. **Output Rendering:**
   - Results are formatted and displayed back to the user via REPL or
     command-line.

---

## Error Handling

- Input validation ensures only correct data is processed.
- Detailed error messages are provided for invalid operations, incorrect input
  formats, or overflow issues.
- REPL gracefully handles errors without terminating the session.

---

## Future Enhancements

- Add support for floating-point calculations.
- Extend functionality to support non-decimal bases, logarithms, and
  trigonometric operations.
- Introduce performance optimizations for larger datasets.
- Implement a GUI wrapper for better user experience.

---

## Directory Structure

```bash
.
├── ARCHITECTURE.md
├── cmd
│   └── repl
│       └── repl.go                  # REPL implementation
├── cross_platform_build.sh          # Cross platform build implementations(Windows ,Linux and Mac)
├── go.mod
├── go.sum
├── IMPLEMENTATION.md                # Step by step guide
├── internal
│   ├── advance_operations.go        # Implementation for Advanced operations
│   ├── advance_operations_test.go   # Tests for advanced operations
│   ├── internal.go                  # Implementation for Arithmetic operations
│   └── unit_test.go                 # Unit tests for arithmetic logic
├── main.go                          # Entry point of the application
└── README.md
```
