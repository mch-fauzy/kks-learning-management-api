# Name: Muchamad Fauzy

# Keraton Karya Solusi LMS API

A Learning Management System (LMS) API for Managing Study Environment

## Table of Contents

- [Setup and Installation](#setup-and-installation)
- [Swagger Documentation](#swagger-documentation)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)

## Setup and Installation

1. Clone the repository and navigate to the root folder

2. Create a `.env` file and set your MySQL DB configurations. Refer to `./infras/mysql.go` for the required parameters

3. Set up the database tables by running the SQL scripts located in the `./migrations` folder in sequence

4. Generate the necessary wire code:

   ```
   go generate ./...
   ```

5. To start the application, run the following command in the project root folder:

   ```
   go run .
   ```

6. The API will be accessible at [http://localhost:8080](http://localhost:8080)

## Swagger Documentation

To access the API documentation using Swagger, follow these steps:

1. Make sure the server is running locally
2. Open your web browser and go to [http://localhost:8080/swagger/doc.json/](http://localhost:8080/swagger/doc.json/)
3. You'll see the Swagger UI interface with a list of endpoints, request parameters, and example requests/responses
4. You can interact with the API directly from the Swagger interface

## API Endpoints

Once the application is up and running, you can interact with the API using the following endpoints:

### View Student Information by ID

- **Endpoint:** `Get /v1/students/{studentId}`
- **Description:** View Student Information by ID, including course and date which enrolled by student
- **Path Parameters:** `studentId`

### Start Concurrency

- **Endpoint:** `Post /startConcurrency`
- **Description:** Start concurrency to print random number

### Stop Concurrency

- **Endpoint:** `Post /stopConcurrency `
- **Description:** stop the concurrency process

### Stop Program

- **Endpoint:** `Post /stopProgram`
- **Description:** Force stop the program

## Contributing

Contributions are welcome! If you want to contribute, please follow these steps:

- Create an issue detailing the feature or bug fix you intend to work on.
- Fork the repository and create a new branch for your feature.
- Implement your changes.
- Create a pull request and reference the issue.
