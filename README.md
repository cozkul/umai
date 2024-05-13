# Dockerized Stack

This project utilizes Docker for containerization, enabling easy deployment and scalability. The stack comprises PostgreSQL, Go, React, and Fiber (PGRF).

## Stack Overview

### Backend
- **Go**: Used for the backend server and API development.
- **Fiber**: A web framework inspired by Express.js, providing a robust server environment.
- **GORM**: An Object-Relational Mapper (ORM) for handling database interactions.
- **Fasthttp**: Serves as the HTTP engine for Fiber.
- **Air**: Enables hot-loading for development efficiency.
- CORS Configured to allow requests to local ports on the same machine - for the initial development

#### Authentication
- **JWT with Http-Only Cookie**: Simple to implement but vulnerable to CSRF attacks.
- **To Be Implemented**: Utilize short-lived access tokens with refresh tokens for improved security.

### Frontend
- **Yarn**: A fast and reliable package manager, an alternative to npm.
- **Vite**: Provides tooling for frontend development, including local hosting and templating with React and TypeScript.
- **React**: A popular library for building user interfaces.
- **Mantine**: A comprehensive component library for React, enhancing UI development.
- **SWR**: A React Hooks library for efficient data fetching.

## Getting Started

To get started with this project, ensure you have Docker installed on your machine. Clone the repository and follow the instructions below to set up the development environment.

### Prerequisites
- Docker
- Docker Compose

### Installation

1. **Clone the repository:**
    ```sh
    git clone https://github.com/cozkul/umai.git
    cd umai
    ```

2. **Start the Docker containers:**
    ```sh
    docker-compose up --build
    ```

3. **Access the application:**
    - Backend API: `http://localhost:5173`
    - Frontend: `http://localhost:3000`

### Development

Currently project is set for hot-loading for both front end and back end. Simply save changes and it will be hot-loaded.