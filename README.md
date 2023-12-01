# Hydro Minder App

![GitHub repo size](https://img.shields.io/github/repo-size/IvanChalukov/go-hydro-minder)
![GitHub last commit](https://img.shields.io/github/last-commit/IvanChalukov/go-hydro-minder)
![GitHub issues](https://img.shields.io/github/issues/IvanChalukov/go-hydro-minder)



### About the Project
Hydro Minder App is a simple tool designed to encourage regular water intake. It's a simple application that serves as a practical example to learn about timers, desktop notifications, and basic error handling in Go. The primary function of this app is to remind me to stay hydrated by sending a desktop notification every 60 minutes.

### Features
- Regular Reminders: The app reminds users to drink water every 60 minutes.
- Desktop Notifications: Utilizes the beeep library to display desktop notifications.
### Getting Started
To get a local copy up and running, follow these simple steps.

#### Prerequisites
- Go (Golang) - The project is developed using Go. Make sure you have Go installed on your machine. Download Go

#### Installation
- Clone the repository:
```sh
git clone https://github.com/IvanChalukov/go-hydro-minder.git
```

- Navigate to the project directory:
```sh
cd go-hydro-minder
```

- Run the application:
```sh
go run main.go
```
- Usage
After running the program, it will send a desktop notification every 60 minutes to remind you to drink water.