# Social Backend (Go + Gin + GORM + JWT)

A production-ready backend starter template built with **Go**, featuring:

- Full authentication system (Register / Login)
- JWT Access Token + Refresh Token
- Role-based authorization (Admin / User)
- User Profile (View + Update)
- Secure MySQL connection (Aiven TLS)
- Clean Architecture (Service / Repository / Handler / Middleware)
- Ready for production deployment

---

## 🚀 Features

### 🔐 Authentication System  
- Register  
- Login  
- Hash password (bcrypt)  
- JWT Authentication  
- Access Token (15 min)  
- Refresh Token (7 days)  
- Refresh Token endpoint  
- Logout-ready architecture  

### 🛡 Authorization  
- Role-based access  
- Admin-only routes  

### 👤 User System  
- Get profile  
- Update profile (avatar)  
- Soft delete (GORM)  

### 🗄 Database  
- Aiven MySQL  
- TLS connection (CA certificate)  
- AutoMigrate  
- Repository layer  

### 🧱 Architecture (Clean Architecture)
internal/
├── auth/ # JWT logic
├── middleware/ # Auth & Admin middleware
├── handler/ # HTTP controllers
├── service/ # Business logic
├── repository/ # DB access
├── model/ # GORM models
├── router/ # Routes setup
└── database/ # MySQL init

yaml
Copy code

---

## ⚙️ Tech Stack

- **Go 1.21+**
- **Gin** (Web framework)
- **GORM** (ORM)
- **MySQL (Aiven Cloud)** with TLS
- **JWT (golang-jwt v4)**
- **bcrypt (password hashing)**

---

## 📦 Installation

Clone the project:

```sh
git clone https://github.com/<yourname>/<repo>.git
cd <repo>
Install dependencies:

sh
Copy code
go mod tidy
🔧 Environment Variables
Create .env in project root:

env
Copy code
MYSQL_DSN="avnadmin:xxxxxx@tcp(mysql-xxxx.aivencloud.com:10104)/defaultdb?charset=utf8mb4&parseTime=True&loc=Local&tls=aiven"

JWT_ACCESS_SECRET="your_access_secret"
JWT_REFRESH_SECRET="your_refresh_secret"
Make sure ca.pem exists in the project root (Aiven CA).

🗄 MySQL Setup
AutoMigrate will automatically create the users table.

▶️ Running the Server
sh
Copy code
go run cmd/server/main.go
Server runs at:

arduino
Copy code
http://localhost:8080
🧪 Testing Guide (非常详细) 🔥
Use Thunder Client / Postman / Insomnia to test.

1️⃣ Health Check
bash
Copy code
GET /health
Response:

json
Copy code
{ "status": "ok" }
2️⃣ Register
arduino
Copy code
POST /auth/register
Body:

json
Copy code
{
  "email": "test@example.com",
  "password": "123456"
}
3️⃣ Login
bash
Copy code
POST /auth/login
Body:

json
Copy code
{
  "email": "test@example.com",
  "password": "123456"
}
Response:

json
Copy code
{
  "access_token": "xxx",
  "refresh_token": "yyy",
  "user": {
      "id": 1,
      "email": "test@example.com",
      "role": "user"
  }
}
Save both tokens.

4️⃣ Refresh Token
bash
Copy code
POST /auth/refresh
Body:

json
Copy code
{
  "refresh_token": "<your refresh token>"
}
Response:

json
Copy code
{
  "access_token": "new_access_token"
}
5️⃣ Get Profile (Requires Access Token)
bash
Copy code
GET /api/profile
Headers:

makefile
Copy code
Authorization: Bearer <access_token>
6️⃣ Update Profile
bash
Copy code
PUT /api/profile
Headers:

makefile
Copy code
Authorization: Bearer <access_token>
Body:

json
Copy code
{
  "avatar": "https://i.imgur.com/xxxxx.png"
}
7️⃣ Admin-Only Route
pgsql
Copy code
GET /admin/stats
Headers:

makefile
Copy code
Authorization: Bearer <access_token>
If role != admin, response:

json
Copy code
{ "error": "admin only" }
To test Admin:

Go to MySQL → set:

sql
Copy code
UPDATE users SET role='admin' WHERE email='test@example.com';
Re-login → test again.

📌 API List
Auth
Method	Endpoint	Description
POST	/auth/register	Register a new user
POST	/auth/login	Login & get tokens
POST	/auth/refresh	Refresh access token

User
Method	Endpoint	Description
GET	/api/profile	Get user profile
PUT	/api/profile	Update user profile

Admin
Method	Endpoint	Description
GET	/admin/stats	Admin-only data

📈 Future Enhancements (Roadmap)
 Email verification (OTP)

 Forgot password

 Redis-based session blacklisting

 Pagination

 Rate limiting

 File upload (avatar)

 Swagger API docs

 Docker Compose

👤 Author
Than — Go Backend Developer
Feel free to fork and extend.

📄 License
MIT License.