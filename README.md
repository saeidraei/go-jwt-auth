# JWT Auth Example in golang using clean architecture
users data can be stored in mysql and mongodb.By default we are using mongodb for storing user data you can set the `server.userDriver` flag to `mysql` in order to change it to mysql.
###how to run using docker-compose:
all you need to do is run `docker-compose up` in project root directory and after some time that images are pulled and some of them are build , project is up and running.

you can use the [postman docs](https://documenter.getpostman.com/view/3010056/SzKTvyo2?version=latest) to explore the api.

### How tables are being created in mysql?
we are using http://github.com/golang-migrate/migrate for handling migrations. in app's container migrations are being run every time it starts running.
if you are not using docker you can run the `migrate` command to run migrations.

### Clean Architecture :
Layers ( from the most abstract to the most concrete ) :
- domain : abstract data structures
- uc : "use cases", the pure business logic
- implem : implementations of the interfaces used in the business logic (uc layer)
- infra : setup/configuration of the implementation
