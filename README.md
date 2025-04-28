# Judgino 🚀

A clean, fast, scalable **Online Judge System** built with **Go** and **Docker**.  
Perfect for handling thousands of submissions efficiently with a microservices structure.

## 🏗️ Architecture

The project follows a **Clean Architecture** pattern, ensuring a clear separation of concerns and maintainability. The system is composed of several independent services, which can be scaled horizontally:

- **API Service (`serve`)**: Main server responsible for handling submission operations, interacting with the database, and serving APIs.
- **Code Runner (`code-runner`)**: Executes user-submitted code within a secure Docker container.
- **Admin Setup (`create-admin`)**: Tool for setting up the initial admin user.

---
---

## 📦 Project Structure

- **serve**: Main API server for handling submissions, users, and general operations.
- **code-runner**: Executes users' submitted code securely inside Docker containers.
- **create-admin**: One-time tool for creating admin users.
- **PostgreSQL**: Used as the database for persisting users, submissions, and other data.

---

## 🛠️ Technologies Used

- **Go** (Golang) — Backend logic
- **Docker** — Containerization
- **Docker Compose** — Service orchestration
- **PostgreSQL** — Database
- **HTML templates** — Frontend rendering
- **Unix Socket** — For Docker-in-Docker execution

---

## ⚙️ Executing Commands

Once the services are up and running, you can execute commands directly within the containers.

### Execute the Services

To run the **Service** interactively inside the container, you can use the following command:

```bash
docker exec judgino serve --cfg ./assets/config/development.yaml
docker exec judgino code-runner --cfg ./assets/config/code-runner.yaml
docker exec judgino create-admin --cfg --username admin123 --password secret123

