# o_clinic

o_clinic is a Go-based application designed to manage clinic operations, including doctor management, patient management, and appointment scheduling.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/the-sabra/o_clinic.git
    cd o_clinic
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up the database:
    ```sh
    # Ensure your database is running and accessible
    # Run the migration scripts to set up the database schema COMING SOON
    ```

4. Create a `.env` file in the root directory and add the following:
    ```env
    PORT=8080
    DBConnString="sqlserver://sa:YourStrong!Password@localhost:1433?database=task&encrypt=disable"
    ```

## Usage

1. Run the application:
    ```sh
    go run cmd/api/main.go
    ```

## Configuration

The application uses environment variables for configuration. These variables can be set in a `.env` file in the root directory.

- `PORT`: The port on which the server will run.
- `DBConnString`: The connection string for the database.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Swagger Documentation

To generate Swagger documentation, run the following command:
```bash
swag init -g cmd/api/main.go -o internal/handler/docs --parseDependency --parseInternal
```

To Access the Swagger end Point
```
 HOST/swagger/
```