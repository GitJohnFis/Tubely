# Contributing to Tubely

Thank you for your interest in contributing to Tubely!

## Setting Up the Development Environment

1. **Install Go**:
   - Ensure you have Go installed (version 1.21 or later recommended). Download it from https://golang.org/dl/.
   - Verify installation by running `go version` in your terminal.

2. **Clone the Repository**:
   ```bash
   git clone https://github.com/GitJohnFis/Tubely.git
   cd Tubely
   ```

3. **Install Dependencies**:
- Run the following command to download dependencies:
```bash
go mod download
```



## Running Tests Locally
To run the test suite locally, follow these steps:

1. Ensure you're in the project root directory (```Tubely```).

2. Run the tests using:
```bash
go test ./... -v
```

The ```-v``` flag provides verbose output for better debugging.
Tests are located in the ```tests/``` directory and follow Go's testing conventions.


3. If new tests are added, ensure they are placed in the ```tests/``` directory with filenames ending in ```_test.go```.


## Contributing Guidelines

- Create a new branch for your changes (e.g., ```git checkout -b feature/ci-tests```).
- Write clear, modular tests following Go's testing best practices (see https://golang.org/pkg/testing/).
- Update this file if new setup steps or dependencies are introduced.
- Submit a pull request to the ```main``` branch for review.

For questions, reach out to the project maintainer or check the repository at https://github.com/GitJohnFis/Tubely/.