# ⚡ Thunder Client API Collection

This document contains sample Thunder Client API requests to test your backend.

All protected routes require:
```
Authorization: Bearer <your_jwt_token>
```

---

## 🔐 1. Login

**POST** `/api/login`

```json
{
  "email": "alice.doc@hospital.com",
  "password": "123456"
}
```

📌 Returns a JWT token on success.

---

## 👨‍⚕️ 2. Create Patient (Receptionist)

**POST** `/api/patients`

Headers:
```
Authorization: Bearer <token>
Content-Type: application/json
```

```json
{
  "name": "John Doe",
  "age": 30,
  "gender": "Male",
  "contact": "9998887777",
  "medical_history": "None"
}
```

---

## 📋 3. Get All Patients

**GET** `/api/patients`

Headers:
```
Authorization: Bearer <token>
```

---

## 📄 4. Get Patient by ID

**GET** `/api/patients/1`

---

## ✏️ 5. Update Full Patient Record (Receptionist)

**PUT** `/api/patients/1`

```json
{
  "name": "Updated Name",
  "age": 40,
  "gender": "Male",
  "contact": "1234567890",
  "medical_history": "Updated history"
}
```

---

## 🩺 6. Update Medical History Only (Doctor)

**PUT** `/api/patients/1/medical`

```json
{
  "medical_history": "Diabetes, High BP"
}
```

---

## 🗑️ 7. Delete Patient (Receptionist)

**DELETE** `/api/patients/1`

---

