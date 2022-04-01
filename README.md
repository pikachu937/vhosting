# A Video Hosting project that can be scalable (vhservice)

#### Available methods:
* POST      /user-edit/       CreateUser
* GET       /user-edit/:id    GetUser
* PUT       /user-edit/:id    UpdateUser
* PATCH     /user-edit/:id    PartiallyUpdateUser
* DELETE    /user-edit/:id    DeleteUser

#### First Starting:
1. Create file ".env" in a root of your app directory and put line "DB_PASSWORD=your_db_pass" in it.
2. Create Database "vhs_db" in your DBMS and execute SQL query file "postgres_db_up.sql" in that database.
