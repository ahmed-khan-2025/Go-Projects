# Go AJAX CRUD 🚀

A simple full-stack CRUD application built with Go (Golang) backend and AJAX-based frontend communication.

---

## 📌 Overview

This project demonstrates how a Go backend API can interact with a frontend using AJAX calls to perform real-time Create, Read, Update, and Delete (CRUD) operations without page reloads.

---

## ✨ Features

* REST API built with Go
* AJAX-based frontend requests
* Create, Read, Update, Delete operations
* JSON data handling
* Simple and responsive UI
* Real-time updates without page refresh

---

## 🛠️ Tech Stack

* Go (Golang)
* HTML, CSS, JavaScript
* AJAX (Fetch API / XMLHttpRequest)
* REST API

---

## 📂 Project Structure

```text id="structure_ajax"
go-ajax-crud/
│
├── main.go
├── handlers/
├── static/
│   ├── index.html
│   ├── script.js
│   └── style.css
└── models/
```

---

## 🚀 How to Run

### 1. Clone the repository

```bash id="clone_ajax"
git clone https://github.com/<your-username>/go-projects.git
cd go-projects/go-ajax-crud
```

---

### 2. Run the application

```bash id="run_ajax"
go run main.go
```

---

### 3. Open in browser

```text id="browser"
http://localhost:8080
```

---

## 🔄 API Endpoints

| Method | Endpoint    | Description     |
| ------ | ----------- | --------------- |
| GET    | /items      | Get all items   |
| POST   | /items      | Create new item |
| PUT    | /items/{id} | Update item     |
| DELETE | /items/{id} | Delete item     |

---

## 🎯 What I Learned

* Building REST APIs in Go
* Handling JSON data
* Frontend-backend communication using AJAX
* CRUD operations design
* Basic full-stack application structure

---

## 📌 Future Improvements

* Add database (PostgreSQL / SQLite)
* Add authentication (JWT)
* Improve UI design
* Add validation and error handling

---

## ⭐ Purpose

This project is part of a Go learning series demonstrating backend development and API integration skills.
