Aim :- Creating an Rest API where an user can sign-up the required fields - id, name, mobile number, address, username, password 

Setting up the Golang Project
Defining the Employee Entity
Connecting to the database
Routing
Implementing CRUD in Golang Rest API
Testing CRUD Operations(Using mockgen)


API Server technology stack is

Server code: golang
REST Server: gin framework
Database: MySQL
ORM: gorm v2
Unit Testing: go test and testify & mockegen
DockerCompose

Usage
$ git clone https://github.com/int1359/RESTApi-GIN-MySQL.git



config:- This package contains a variable DB of type pointer to sql.DB, which is a database connection handle. It also defines a function DbURL() which returns a string that represents a database URL built from environment variables.

The util.GetEnvVariable() function retrieves the value of an environment variable, which is used to build the database URL string that is returned by the DbURL() function.

model:- This package contains Go structs that represent different entities in the application.
* "Employee" struct represents a Employee entity with the following fields:
* Success Response for each EndPoints (GetAllEmployeeResponse,CreateEmployeeResponse...) along with ErrorResponse

routes:- This package is responsible for setting up the routing of a web application using the Gin framework.Mapping of each end point to the contoller function.

The code imports necessary packages for the application's controller, data access object (DAO), and service layer.
The function "SetupRouter" creates instances of the Employee DAO, Service, and Controller, and sets up the following RESTful API endpoints for the Employee resource:
* GET /Employee: returns a list of all Employee
* GET /Employee/:id: returns a Employee with a given id
* PUT /Employee/:id: returns a Updated Employee in a given id
* POST /Employee: creates a new Employee
Finally, the function returns a Gin Engine instance with the routes set up, which can be used to start the application.

dao:-  The EmployeeDao interface defines the four methods for accessing Employee data, GetAllEmployees(), CreateEmployee(), GetEmployeeByID(), and UpdateEmployee(). 
The EmployeeDaoImpl struct implements these methods and provides the actual logic to interact with a database. The methods use the config.DB object to execute SQL queries to retrieve and manipulate the student data.
The EmployeeDao interface specifies the methods to be implemented for accessing and manipulating Employee data. The implementation of the interface is provided in the struct EmployeeDaoImpl.

service:-This layer call the dao layer for the specific Endpoints. Also for createEmployee() service Validation for Employee field is also added.

controller:-This layer is Responsible for handaling the Response and Errors returned by the service layer

main:- It is the entry point of the program. The package has the following functions:
* init(): This function initiallizes the database connection and sets it up to use the MySQL driver. The function also checks if the connection was successful. If it fails, it logs an error.
* main(): This is the main function that runs when the program is executed. It sets up the routing of the application using the "SetupRouter" function of the "routes" package and starts the application with the "Run" method.



