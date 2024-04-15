## Udacity project
### Project: CRM Backend

### Best Practices
* Success Criteria
  * Set up a proper Go environment 
  * Document the project in the README
  * Organize and write clean code
  * Build an intuitive user experience


* Specifications
  * The project requires only a simple go run command to launch the application.
  * The project README contains a description of the project, and also includes instructions for installation, launch, and usage.
    - Syntax and semantics of language features (e.g., functions, variables, loops, etc.) are well-formed and free from errors or warnings in the console
    - Data structures, handlers, routes, imports, and other assets are organized logically (e.g., grouped together) and are easy to find
  * Users can interact with the application (i.e., make API requests) by simply using Postman or cURL.


### Data
* Success Criteria   
   * Create a Customer struct
   * Create a mock "database" to store customer data
   * Seed the database with initial customer data
   * Assign unique IDs to customers in the database

* Specifications
  * Each customer includes:
    - ID
    - Name
    - Role
    - Email
    - Phone
    - Contacted (i.e., indication of whether or not the customer has been contacted)
    - Data is mapped to logical, appropriate types (e.g., Name should not be a bool).
  * Customers are stored appropriately in a basic data structure (e.g., slice, map, etc.) that represents a "database."
  * The "database" data structure is non-empty. That is, prior to any CRUD operations performed by the user (e.g., adding a customer), the database includes at least three existing (i.e., "hard-coded") customers.
  * Customers in the database have unique ID values (i.e., no two customers have the same ID value).



### Server
* Success Criteria
   * Serve the API locally


* The application handles the following 5 operations for customers in the "database":
    - Getting a single customer through a /customers/{id} path
    - Getting all customers through a the /customers path
    - Creating a customer through a /customers path
    - Updating a customer through a /customers/{id} path
    - Deleting a customer through a /customers/{id} path
* Each RESTful route is associated with the correct HTTP verb.


* The Handler interface is used to handle HTTP requests sent to defined paths. There are five routes that return a JSON response, and are each is registered to a dedicated handler:

      getCustomers()
      getCustomer()
      addCustomer()
      updateCustomer()
      deleteCustomer()

## Migrate
[Github](https://github.com/golang-migrate/migrate?tab=readme-ov-file#cli-usage)

### install windows
    $ scoop install migrate

### Installing Golang Migrate on Ubuntu

1. Let us setup the repository to install the migrate package.

        $ curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
2. Update the system by executing the following command.

        $ sudo apt-get update
3. Execute the following command in the terminal to install migrate.

       $ sudo apt-get install migrate


##
    $ migrate create -ext=sql -dir=sql/migrations -seq init
    $ migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/crm" -verbose up
    $ migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/crm" -verbose down


## Run go main
    $ docker-compose up -d
    $ go run .\main.go .\wire_gen.go