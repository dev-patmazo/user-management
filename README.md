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

## Database
1. Install MySQL database
2. Create database userdb

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
DB_USER=admin
DB_PASSWORD=password
DB_HOST=localhost:3306
DB_NAME=usersdb
```

5. Run the application
```
go run main.go
```

6. Open db.sql and insert the initial user data to mysql database.

7. Play around with the request using the collection thunder-collection_sample-project
you can import it to your desired rest client tool.