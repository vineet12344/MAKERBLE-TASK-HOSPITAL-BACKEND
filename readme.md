# 🏥 Hospital Management Backend (Golang + Gin + PostgreSQL)

This is a backend system for a hospital management application built as part of a backend development internship task. It provides secure login and role-based access for doctors and receptionists, supports patient record management, and uses best practices such as JWT authentication, middleware-based authorization, and PostgreSQL as a relational database.

---

## 🚀 Live Deployment

- 🔗 **Base API URL**: [https://makerble-task-hospital-backend.onrender.com](https://makerble-task-hospital-backend.onrender.com)
- 📘 **Swagger Docs**: [https://makerble-task-hospital-backend.onrender.com/swagger/index.html](https://makerble-task-hospital-backend.onrender.com/swagger/index.html)

---

## 🚀 Features

- 🔐 **JWT Authentication** for login
- 👩‍⚕️ **Doctor** access:
  - View patients
  - Update medical history only
- 🧾 **Receptionist** access:
  - Create, view, update, and delete patient records
- 🛡️ **Role-based access control** using middleware
- 🗄️ PostgreSQL database with GORM ORM
- 🧪 Unit tested utils and middleware
- 🧱 Clean project structure following production patterns

---

## 📁 Folder Structure

```bash
HOSPITAL-BACKEND/
├── controllers/              # Route handler logic
├── docs/                     # API endpoint documentation
│   ├── endpoints.md
│   └── thunderclient.md
├── initializers/             # DB and env setup
├── middleware/               # JWT + Role checking middleware
├── models/                   # GORM models (User, Patient)
├── routes/                   # All route definitions
├── utils/                    # Token and password helpers
├── .env                      # Local environment config (ignored in git)
├── .env.example              # Sample config for setup
├── .gitignore
├── go.mod / go.sum           # Dependencies
├── main.go                   # App entry point
├── Dockerfile
├── .dockerignore
└── readme.md                 # Project documentation
```

---

## 🛠️ Tech Stack

- **Language:** Golang
- **Framework:** Gin Web Framework
- **ORM:** GORM
- **Database:** PostgreSQL
- **Authentication:** JWT
- **Password Security:** bcrypt
- **Testing:** `testing`, `httptest`, `stretchr/testify`

---

## 📦 Setup Instructions

### 1. Clone the Repo

```bash
git clone https://github.com/yourusername/hospital-backend.git
cd hospital-backend
```

### 2. Create `.env` File

Copy `.env.example` and update credentials:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=yourpassword
DB_NAME=hospital
JWT_SECRET=your_jwt_secret
PORT=8080
```

### 3. Run the App

```bash
go mod tidy
go run main.go
```

> ✅ Auto-migration will create tables on first run.

---

## 🧪 Run Tests

```bash
go test ./... -v
```

Tested Components:
- Token generation/validation
- Password hashing/comparison
- Auth & role middleware

---

## 🔐 API Endpoints

### Authentication

| Method | Endpoint        | Description                       |
|--------|------------------|-----------------------------------|
| POST   | `/api/signup`    | Register new user (doctor/receptionist) |
| POST   | `/api/login`     | Login with credentials            |

### Patients (Protected)

| Method | Endpoint                     | Role         | Description                      |
|--------|-------------------------------|--------------|----------------------------------|
| POST   | `/api/patients`              | Receptionist | Register a new patient           |
| GET    | `/api/patients`              | Any          | Get all patients                 |
| GET    | `/api/patients/:id`          | Any          | Get a patient by ID              |
| PUT    | `/api/patients/:id`          | Receptionist | Update entire patient record     |
| PUT    | `/api/patients/:id/medical`  | Doctor       | Update medical history only      |
| DELETE | `/api/patients/:id`          | Receptionist | Delete a patient record          |

> 🔒 All protected routes require:
> ```
> Authorization: Bearer <your_jwt_token>
> ```

---

## 📄 Sample .env File

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=your_db_password
DB_NAME=hospital
JWT_SECRET=supersecretjwtkey
PORT=8080
```

---

## 🧾 Internship Task Coverage

✅ Single login endpoint for doctor & receptionist  
✅ Signup for both roles  
✅ JWT-based authentication  
✅ RBAC middleware (doctor vs. receptionist)  
✅ REST API with GORM & PostgreSQL  
✅ Auto-migration  
✅ Test cases for critical logic  
✅ Thunder Client + Swagger documentation  
✅ Clean README and `.env.example`  
✅ Dockerized & hosted on Render

---

## 👤 Author

**Vineet Salve**  
B.Tech, Computer Engineering   
GitHub: [vineet12344](https://github.com/vineet12344)

---

## 📎 License

MIT (or as applicable for the internship submission)
