# A scalable REST API Video Hosting

#### Available methods:
* POST      /user-interface/       POSTUser
* GET       /user-interface/:id    GETUser
* GET       /user-interface/all    GETAllUsers
* PUT       /user-interface/:id    PUTUser
* PATCH     /user-interface/:id    PATCHUser
* DELETE    /user-interface/:id    DELETEUser

* POST      /auth/sign-up/         POSTUser
* POST      /auth/sign-in/         POSTUser

#### First Starting:
1. Create file ".env" in a root of your app directory and put line "DB_PASSWORD=your_db_pass" in it.
2. Create Database "vhosting_db" in your DBMS and execute SQL query file "postgres_db_up.sql" in that database.
