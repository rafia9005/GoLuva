# Hi everyone :wave:

This project is a Go Fiber Starter template that includes ORM (Object-Relational Mapping), authentication, and JWT (JSON Web Token) features. With this starter, you can quickly and easily develop a secure and high-performance Go application.

<details>
<summary>
  How To Installation?
</summary>
  
### Docs
Getting Started

Follow these steps to get started with the Go Fiber Starter Template [Docs](https://www.efset.org/cert/5P5Pp1)


### Technologies

- Golang
- Go Fiber
- GORM
- x crypto bcrypt                                                      |
  
## Router 📕
| Method | Route                                                | QUERY                                             |
| ------------------ | ------------------------------------------------------ | ------------------------------------------------ |
| POST               | http://localhost:8000/login | [QUERY](https://github.com/aellopos) |
| POST               | http://localhost:8000/register | [QUERY](https://github.com/aellopos) |
| GET               | http://localhost:8000/users | [QUERY](https://github.com/aellopos) |
| GET              | http://localhost:8000/users/{id} | [QUERY](https://github.com/aellopos) |
| POST               | http://localhost:8000/users | [QUERY](https://github.com/aellopos) |
| DELETE             | http://localhost:8000/users/{id} | [QUERY](https://github.com/aellopos) |
| PUT             | http://localhost:8000/users/{id} | [QUERY](https://github.com/aellopos) |

</details>

### Installation

```bash
git clone https://github.com/rafia9005/GoLuva
cd GoLuva
go mod tidy
go run main.go
```

### Configuration on .env

```env
APP_DATABASE = "mysql://username:password@tcp(localhost:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
APP_PORT = "3000"
APP_SECRET = "your-very-secure-secreet"
```
