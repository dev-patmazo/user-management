# User Management REST API

This is a simple REST API for managing users, built with Go, Gorilla Mux, and GORM.

## Requirements

- Go 1.21.3 or later
- MySQL

## Dependencies

- Gorilla Mux v1.8.1
- godotenv v1.5.1
- GORM v1.25.9
- GORM MySQL Driver v1.5.6

## Setup

1. Clone the repository:
```
git clone https://github.com/dev-patmazo/user-management.git
```

2. Navigate to the project directory.
```
cd user-management
```

3. Install the dependencies
```
go mod download
```

4. Setup environment file
```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=user_management
```

5. Run the application
```
go run main.go
```