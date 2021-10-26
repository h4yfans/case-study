
### Backend Case Study


#### Requirements 
* Go Language
* Postgres or MongoDB
* Docker
* Git
* Go Module

You need to develop a REST API which is used to manage a user management service. It has some basic CRUD operations on user system.

Let's say, it's listening on localhost at port 8080;

These are example requests and responses it needs to provide;

---

#### Add a user

##### Sample Request
```
curl -X PUT \
 -d '{"name": "Test", "email": "test@example.com", "password": "securepasswd"}' \
 -H 'Content-Type: application/json' \
  http://localhost:8080/users
```

##### Errors:

| Status Code | Description | Sample Response |
| --  | -- | -- |
| 200 | Success | {"id": 1, "name": "Test", "email": "test@example.com"} |
| 400 | When request body or parameters wrong | {"error": "Bad request"}|
| 403 | If user already exists | {"error": "User with that email already exists"} |
| 500 | When something unexpected happens | {"error": "server error"} |

---

#### Edit a user's attributes

##### Sample Request
```
curl -X PATCH \
 -d '{"name": "No name", "password": "strongpasswd"}' \
 -H 'Content-Type: application/json' \
  http://localhost:8080/users/1
```

##### Errors:

| Status Code | Description | Sample Response |
| --  | -- | -- |
| 200 | Success | {"id": 1, "name": "No name", "email": "test@example.com"} |
| 400 | When request body or parameters wrong | {"error": "Bad request"}|
| 404 | If user not found | {"error": "User with that id does not exist"} |
| 500 | When something unexpected happens | {"error": "server error"} |

---

#### Delete a user

##### Sample Request
```
curl -X DELETE \
  http://localhost:8080/users/1
```

##### Errors:

| Status Code | Description | Sample Response |
| --  | -- | -- |
| 200 | Success |  |
| 400 | When request body or parameters wrong | {"error": "Bad request"}|
| 404 | If user not found | {"error": "User with that id does not exist"} |
| 500 | When something unexpected happens | {"error": "server error"} |

---

#### Find a user with ID

##### Sample Request
```
curl -X GET \
  http://localhost:8080/users/1
```

##### Errors:

| Status Code | Description | Sample Response |
| --  | -- | -- |
| 200 | Success | {"id": 1, "name": "No name", "email": "test@example.com"} |
| 400 | When request body or parameters wrong | {"error": "Bad request"}|
| 404 | If user not found | {"error": "User with that id does not exist"} |
| 500 | When something unexpected happens | {"error": "server error"} |


---

#### Get All Users

##### Sample Request
```
curl -X GET \
  http://localhost:8080/users
```

##### Errors:

| Status Code | Description | Sample Response |
| --  | -- | -- |
| 200 | Success | [{"id": 1, "name": "No name", "email": "test@example.com"}, {"id": 2, "name": "Test 2", "email": "test2@example.com"}] |
| 400 | When request body or parameters wrong | {"error": "Bad request"}|
| 404 | If user not found | {"error": "User with that id does not exist"} |
| 500 | When something unexpected happens | {"error": "server error"} |


--- 

### Expectations

* This service should be written in Go lang, you can use any framework or library you need. 
* There should be a database system to store user data, preferably Postgres or MongoDB. 
* Dockerfile should be provided in the project.
* Code Quality and Design

#### Extras (Optional)

* Kubernetes deployment setup (yaml file[s])
* Unit tests


### Timeline

You have 4 days, Good luck.


