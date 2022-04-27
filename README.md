# Video Hosting Web-Service (REST API)

#### Available requests:

* POST      /auth/sign-in            CreateUser
* POST      /auth/change-password    CreateUser
* POST      /auth/sign-out           CreateUser
* POST      /user-interface          CreateUser
* GET       /user-interface/:id      GetUser
* GET       /user-interface/all      GetAllUsers
* PATCH     /user-interface/:id      PartiallyUpdateUser
* DELETE    /user-interface/:id      DeleteUser

#### First Starting:

1. Create file ".env" in directory "./configs/" and post variables from example ".env.example".
2. Create database "vhosting" in your DBMS and create tables by executing SQL query file "vhosting_database.sql".
