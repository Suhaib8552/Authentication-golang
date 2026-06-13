<h1>Authentication using Golang & React</h1>

A full-stack authentication system built with Golang (Gin), PostgreSQL, JWT Authentication, and React.

<h2>Features:</h2>
<ul>
  <li>User Registration</li>
<li>User Login</li>
<li>JWT Token Generation</li>
<li>JWT Authentication Middleware</li>
<li>Protected Profile Route</li>
<li>User Logout</li>
<li>Token Blacklisting</li>
<li>Password Hashing with bcrypt</li>
<li>PostgreSQL Database Integration</li>
<li>React Frontend Integration</li>
</ul>

<h2>Tech Stack</h2>

<h3>Backend</h3>
<ul>
<li>Golang</li>
<li>Gin Web Framework</li>
<li>PostgreSQL</li>
<li>JWT (JSON Web Tokens)</li>
<li>bcrypt</li>
  </ul>
<h3>Frontend</h3>
<ul>
<li>React</li>
<li>React Router</li>
<li>Fetch API</li>
</ul>

<h2>Environment Variables</h2>

Create a .env file in the project root.

DB_HOST=localhost

DB_PORT=5432

DB_USER=postgres

DB_PASSWORD=your_password

DB_NAME=mydb

JWT_SECRET=your_secret_key

<h2>Installation</h2>

<h3>Clone Repository</h3>

git clone https://github.com/Suhaib8552/Authentication-golang.git
cd auth-go

<h3>Install Dependencies</h3>
go mod tidy
Run PostgreSQL

Ensure PostgreSQL is installed and running.

<h3>Create a database:</h3>

CREATE DATABASE mydb;

<h3>Run Application</h3>

go run main.go

<h3>Server starts on:</h3>

http://localhost:8000

<h2>API Endpoints</h2>
<h3>Register User</h3>
POST /api/v1/auth/register

<h4>Request:</h4>

{
  "username": "suhaib",
  "email": "suhaib@gmail.com",
  "password": "password123"
}

<h4>Response:</h4>

{
  "message": "User registered successfully",
  "userId": 1
}
<h3>Login User</h3>
POST /api/v1/auth/login

<h4>Request:</h4>

{
  "email": "suhaib@gmail.com",
  "password": "password123"
}

<h4>Response:</h4>

{
  "Status": "Successfully generated the token",
  "token": "<jwt_token>"
}
<h3>Get Profile</h3>
GET /api/v1/auth/profile

<h4>Headers:</h4>

Authorization: Bearer <jwt_token>

<h4>Response:</h4>

{
  "id": 1,
  "username": "suhaib",
  "email": "suhaib@gmail.com"
}
<h3>Logout</h3>
POST /api/v1/auth/logout

<h4>Headers:</h4>

Authorization: Bearer <jwt_token>

Response:

{
  "message": "Logged out successfully"
}
