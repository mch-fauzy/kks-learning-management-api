# Keraton Karya Solusi Learning Management System (LMS) API

This API provides a platform for managing students, lecturers, courses, enrollments, and exams, making it easier to handle educational activities and data.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
- [Swagger Documentation](#swagger-documentation)
- [API Endpoints](#api-endpoints)
- [Error Handling](#error-handling)
- [Contributing](#contributing)

## Features

- View, add, update, and delete student records
- Manage lecturer information
- Create and manage courses, including assigning lecturers
- Record and grade exams
- Calculate student GPAs based on exam grades
- Pagination for managing large lists of students

## Getting Started

1. Clone this repository:

   ```
   git clone https://github.com/mch-fauzy/kks-learning-management-api.git
   ```

2. Navigate to the project directory:

   ```
   cd kks-learning-management-api
   ```

3. Create a `.env` file and set your MySQL DB configurations. Refer to `./infras/mysql.go` for the required parameters

4. Set up the database tables by running the SQL scripts located in the `./migrations` folder in sequence

5. Generate the necessary wire code:

   ```
   go generate ./...
   ```

6. To start the application, run the following command in the project root folder:

   ```
   go run .
   ```

7. The API will be accessible at [http://localhost:8080](http://localhost:8080)

## Swagger Documentation

To access the API documentation using Swagger, follow these steps:

1. Make sure the server is running locally
2. Open your web browser and go to [http://localhost:8080/swagger/doc.json/](http://localhost:8080/swagger/doc.json/)
3. You'll see the Swagger UI interface with a list of endpoints, request parameters, and example requests/responses
4. You can interact with the API directly from the Swagger interface

## API Endpoints

Once the application is up and running, you can interact with the API using the following endpoints:

### View All Student Information

- **Endpoint:** `Get /v1/students`
- **Description:** View a list of all student information
- **Query Parameters:** `page` (default: 1), `pageSize` (default: 10)

### View Student Information by ID

- **Endpoint:** `Get /v1/students/{studentId}`
- **Description:** View Student Information by ID, including course and date which enrolled by student
- **Path Parameters:** `studentId`

## Contributing

Contributions are welcome! If you want to contribute, please follow these steps:

- Create an issue detailing the feature or bug fix you intend to work on.
- Fork the repository and create a new branch for your feature.
- Implement your changes.
- Create a pull request and reference the issue.
