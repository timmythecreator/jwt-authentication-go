# JWT Authentication Go App using Gin and Gorm

A simple implementation of JWT authentication with three routes:

1. **POST `/signup`** - Register a user. Accepts two parameters: `email` and `password`.
2. **POST `/login`** - Log in a user. Accepts two parameters: `email` and `password`.
3. **GET `/validate`** - Validate the user who makes the request.

## Usage

1. **Clone the repository**
    ```sh
    git clone https://github.com/timmythecreator/jwt-authentication-go
    ```

2. **Create a database for users and link it in the `.env` file**

3. **Download all necessary modules**
    ```sh
    go mod tidy
    go mod verify
    ```

## .env File

Make sure to create a `.env` file in the root directory of your project with the following content:

```env
PORT=3000
JWT_SECRET=UNIQUE_TOKEN
DB="host=host-of-your-db user=username password=password dbname=db-name port=5432 sslmode=disable"
```

## API Endpoints
### Register User
**POST `/signup`**

**Request Body:**
```json
{
    "email": "user@example.com",
    "password": "password"
}
```

**Response:**
```json
{
    "message": "Signup successful"
}
```

### Log In User
**POST `/login`**

**Request Body:**
```json
{
    "email": "user@example.com",
    "password": "password"
}
```

**Response:**
```json
{
    "message": "Login successful",
    "token": "your_jwt_token"
}
```

### Validate User
**GET `/validate`**

**Headers:**

```http
Authorization: Bearer your_jwt_token
```

Response:
```json
{
    "message": "User validated"
}
```

## Dependencies
* [Gin](https://gin-gonic.com/)
* [JWT](https://github.com/golang-jwt/jwt)
* [Gorm](https://gorm.io/)
* [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
* [GoDotEnv](https://github.com/joho/godotenv)