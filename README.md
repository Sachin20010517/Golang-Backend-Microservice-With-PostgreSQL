# Golang-Backend-Microservice-With-PostgreSQL

A RESTful API for managing employees using Supabase PostgreSQL as the database and the Fiber framework in Go for the backend. This application supports employee creation, retrieval, updating, and deletion with proper data validation and error handling.

## Features

- **Create Employee**: Add a new employee to the database.
- **Retrieve Employees**: Fetch all employees or find a specific employee by `employee_id`.
- **Update Employee**: Update the details of an existing employee.
- **Delete Employee**: Soft-delete an employee by marking them as deleted.

---

## Table of Contents

- [Prerequisites](#prerequisites)
- [Project Setup](#project-setup)
- [Database Configuration](#database-configuration)
- [Endpoints](#endpoints)
- [Folder Structure](#folder-structure)
- [Contributing](#contributing)
- [License](#license)

---

## Prerequisites

Ensure you have the following installed:

1. [Go](https://golang.org/dl/) (v1.18 or higher)
2. [PostgreSQL](https://www.postgresql.org/)
3. [Supabase](https://supabase.com/) account for hosting the database (optional)
4. [Fiber](https://gofiber.io/) framework
5. `go mod` for dependency management

---

## Project Setup

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-username/employee-management.git
   cd employee-management

