# AuthBase

A modern, secure authentication system built with Go and Gin framework, featuring passwordless magic link authentication and a clean web interface.

## 🚀 Features
- **Only frontend was vibecoded**
- **Passwordless Authentication**: Secure magic link-based login system
- **SQLite Database**: Lightweight, embedded database for user management
- **JWT Tokens**: Secure session management with JSON Web Tokens
- **Admin Dashboard**: Protected admin area for authenticated users
- **Responsive UI**: Vibecoded. Priority was GO backend
- **Email Integration**: Automated magic link delivery via email
- **Service-Ready**: Production deployment scripts

## 🛠️ Tech Stack

- **Backend**: Go 1.24.4 with Gin web framework
- **Database**: SQLite with GORM ORM
- **Authentication**: JWT tokens + Magic links
- **Frontend**: HTML/CSS with responsive design
- **Email**: Built-in email service for magic link delivery
- **Logging**: Structured logging with file rotation

## 🚀 Quick Start

### Production Deployment

1. **Run the installation script** (as root):
   ```bash
   sudo ./firstInstall.sh
   ```

2. **Move to production directory (where you can develop)**:
   ```bash
   sudo mkdir -p /var/lib/AuthBase/app
   sudo cp -r . /var/lib/AuthBase/app/
   ```

3. **Build the application**:
   ```bash
   cd /var/lib/AuthBase/app
   go build -o authbase cmd/web/main.go
   ```

4. **Run as service user**:
   ```bash
   sudo -u WebServiceUser ./authbase
   ```

## ⚙️ Configuration

### Environment Variables (.env)

```env
APP_NAME=AuthBase
TESTING=false
PORT=8080
HOST=localhost
DEBUG=false
EMAILER_SECRET_KEY=your-email-secret-key
JWT_SECRET=your-jwt-secret-key
SERVICE_USER=WebServiceUser
LOG_DIR=/var/log/AuthBase
DB_DIR=/var/lib/AuthBase/db
```

### Important Configuration Notes

- Set `DEBUG=false` in production
- Change `EMAILER_SECRET_KEY` and `JWT_SECRET` to secure values
- Update `baseURL` in `internal/utils/magicLink.go` for production domain (forgot to update this line XD)
- Adjust file permissions to 750/700 for production security

## 🏗️ Project Structure

```
AuthBase/
├── cmd/
│   ├── web/           # Main application entry point
│   └── webTest/       # Testing utilities
├── internal/
│   ├── controllers/   # HTTP request handlers
│   ├── db/           # Database connection and operations
│   ├── middleware/   # Authentication middleware
│   ├── models/       # Data models
│   ├── routes/       # Route definitions
│   └── utils/        # Utility functions (JWT, email, magic links)
├── templates/        # HTML templates
├── assetsssss/      # Static assets (CSS, images, etc.)
├── .env             # Environment configuration
├── firstInstall.sh  # Production installation script
├── go.mod           # Go module dependencies
└── README.md        # This file
```

## 🔐 Authentication Flow

1. **User Registration/Login**:
   - User enters email address
   - System generates secure magic link token
   - Magic link sent via email
   - Token expires after 5 minutes

2. **Magic Link Verification**:
   - User clicks magic link
   - System validates token
   - New user account created if needed
   - JWT token issued and stored in secure cookie

3. **Protected Access**:
   - JWT token validated on protected routes
   - Access to admin dashboard granted
   - Session managed via HTTP-only cookies

## 🛡️ Security Features

- **Secure Token Generation**: Cryptographically secure random tokens
- **Token Expiration**: Magic links expire after 5 minutes
- **JWT Security**: Secure cookie storage with HTTP-only flag
- **Email Validation**: Regex-based email format validation
- **Database Protection**: GORM prevents SQL injection
- **Service User**: Runs under dedicated system user with minimal privileges

## 📝 API Endpoints

### Public Routes
- `GET /` - Landing page
- `GET /signup` - Registration page
- `GET /login` - Login page
- `POST /magicLinkGet` - Request magic link
- `GET /verifymagicregister` - Verify magic link token
- `GET /emailsent` - Confirmation page

### Protected Routes
- `GET /admindashboard` - Admin dashboard (requires authentication)

## 🔧 Development


### Database Operations
The application automatically:
- Creates SQLite database on first run
- Runs migrations to set up user table
- Seeds admin user (admin@example.com)

## 📊 Logging

- **Location**: `/var/log/AuthBase/` (configurable)
- **Format**: Structured JSON logging
- **Rotation**: Automatic log file rotation
- **Levels**: Info, Error, Debug (when DEBUG=true)

## 🚨 Security Considerations

**Current Security Notes** (from existing documentation):
- JWT cookies should be marked as Secure in production
- Rate limiting should be implemented for DoS protection
- Server-side session management could be enhanced
- Debug mode should be disabled in production


### Common Issues

1. **Permission Errors**:
   ```bash
   sudo chown -R WebServiceUser:WebServiceUser /var/lib/AuthBase
   sudo chmod 750 /var/lib/AuthBase
   ```

2. **Database Connection Issues**:
   - Check DB_DIR permissions
   - Ensure SQLite3 is installed
   - Verify service user has write access

3. **Email Delivery Problems**:
   - Verify EMAILER_SECRET_KEY configuration
   - Check application logs for email errors
   - Ensure mail service is configured

4. **Port Already in Use**:
   ```bash
   sudo netstat -tlnp | grep :8080
   sudo kill -9 <PID>
   ```

### Logs Location
- Application logs: `/var/log/AuthBase/`
- System logs: `journalctl -u authbase` (if running as systemd service)
