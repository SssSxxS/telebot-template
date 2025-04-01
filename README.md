# Telegram Bot Template

This project is a ready-to-use template for developing Telegram bots in Go. The template includes a basic project structure, configured components, and ready-to-use functionality.

## Technologies Used

- **[telebot v4](https://github.com/tucnak/telebot)** - Elegant library for creating Telegram bots
- **[GORM](https://gorm.io/)** - Powerful ORM for Go with SQLite support
- **[zerolog](https://github.com/rs/zerolog)** - Fast and structured logging system with zero memory allocation
- **[godotenv](https://github.com/joho/godotenv)** - Loading environment variables from .env file

## Implemented Features

### User System

- **User model** with fields:
  - Telegram ID
  - Username
  - Status (-1: banned, 0: blocked, 1: active)
  - Role (administrator/regular user)

### Middleware

- **UserTracker** - Automatic user tracking in the database
- **IsAdmin** - Administrator rights verification
- **IgnoreOld** - Ignoring outdated messages

### Role Separation

- Separate handlers for users and administrators
- Role-based access control system

### Database

- SQLite with GORM
- Automatic model migration
- Repository for user management

### Logging

- Structured logging using zerolog
- Output logs to console and file
- Automatic creation of log files by date

## Project Structure

```
.
├── cmd/
│   └── main.go           # Application entry point
├── data/
│   ├── db/               # Directory for database files
│   └── logs/             # Directory for log files
├── internal/
│   ├── bot/
│   │   ├── bot.go        # Bot initialization and configuration
│   │   ├── handlers/     # Command handlers
│   │   │   ├── admin/    # Handlers for administrators
│   │   │   └── user/     # Handlers for users
│   │   └── middleware/   # Middleware
│   ├── database/
│   │   ├── database.go   # Database initialization
│   │   ├── models/       # Data models
│   │   └── repositories/ # Repositories for data management
│   └── lib/
│       └── logger/       # Logging configuration
├── .env                  # File with environment variables
├── go.mod                # Project dependencies
└── go.sum                # Dependency checksums
```

_readme by claude-3.7-Sonnet <3_
