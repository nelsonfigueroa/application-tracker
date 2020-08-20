# Application Tracker

A simple JSON API built with Golang to track applications. At the moment, it does not save data to disk.

Third party libraries:
- [mux](https://github.com/gorilla/mux): For routing purposes

# Endpoints

| HTTP Verb | Endpoint          | Description              |
|-----------|-------------------|--------------------------|
| GET       | /applications     | Get all applications     |
| GET       | /applications/:id | Get a single application |
| POST      | /applications     | Create an application    |
| DELETE    | /applications/:id | Delete an application    |

The Application model consists of the following attributes:

| Attribute | Type     |
|-----------|----------|
| Id        | `string` |
| Date      | `string` |
| Company   | `string` |
| Position  | `string` |
| Location  | `string` |

# Running Locally

Assuming Go is installed, simply run

```
go run main.go
```