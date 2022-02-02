# golang-restful-api
GoLang restful API demonstrating CRUD operations

## Installation
Download and install GoLang if you don't have [https://go.dev/doc/install](https://go.dev/doc/install)

#### Checkout and Run
```sh
$ git clone git@github.com:mshafiq9/golang-restful-api.git
$ cd golang-restful-api
$ go run .
```

#### Testing
Note: During testing also cross verify console output when hitting different endpoints.

```
# Home page
http://localhost:10000/

# Get all articles
http://localhost:10000/articles

# Get one article
http://localhost:10000/article/2
```

```sh
# Create:
curl -i -X POST -H 'Content-Type: application/json' -d '{"Id": "3", "Title": "Newly Created Post 3", "desc": "The description for my new post 3", "content": "my articles content 3"}' http://localhost:10000/article

# Read
curl -i -X GET http://localhost:10000/article/2

# Update:
curl -i -X PUT -H 'Content-Type: application/json' -d '{"Id": "3", "Title": "Updated Created Post 3", "desc": "Updated description for my new post 3", "content": "Updated my articles content 3"}' http://localhost:10000/article/3

# Delete:
curl -i -X DELETE http://localhost:10000/article/3
```

## References
https://tutorialedge.net/golang/creating-restful-api-with-golang/

