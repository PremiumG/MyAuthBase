#!/bin/bash
if [ "$EUID" -ne 0 ]; then
    echo "Please run as root (use sudo)"
    exit 1
fi
#install updates etc
echo "Updating package list..."
apt update

echo "Upgrading packages..."
apt upgrade -y


#check for .env file
echo "Retrive .env file"
if [ -f .env ]; then
    echo "Loading configuration from .env"
    export $(grep -v '^#' .env | xargs)
else
    echo "No .env file found, using defaults"
fi

#get .env file
APP_NAME="${APP_NAME:-AuthBase}"
SERVICE_USER="${SERVICE_USER:-AppServiceUser}"
LOG_DIR="${LOG_DIR:-/var/log/AuthBase}"
DB_DIR="${DB_DIR:-/var/lib/AuthBase/db}"

CURRENT_USER=$(whoami)  #placeholder

# Check Go installation
if ! command -v go &> /dev/null; then
    echo "Go not found. Installing..."
    install_go
else
    GO_VERSION=$(go version | awk '{print $3}')
    echo "Found Go: $GO_VERSION"
fi

# Check if we're in Go project directory
if [ ! -f "go.mod" ]; then
    echo "ERROR: go.mod not found. Run from project root."
    exit 1
fi

#Check and/or install sqlite3
echo "Checking if sqlite3 is installed..."
if command -v sqlite3 &> /dev/null; then
    echo "sqlite3 is already installed: $(sqlite3 --version)"
else
    echo "Installing sqlite3..."
    apt install -y sqlite3
    
    if command -v sqlite3 &> /dev/null; then
        echo "sqlite3 installed successfully: $(sqlite3 --version)"
    else
        echo "Failed to install sqlite3"
        exit 1
    fi
fi

#service users and directories/files
echo "Setting up service user and directories..."
if ! id "$SERVICE_USER" &>/dev/null; then
    useradd --system --shell /bin/false --create-home "$SERVICE_USER"
    echo "Created service user: $SERVICE_USER"
else
    echo "Service user $SERVICE_USER already exists"
fi

# if ! getent group "AUTH_BASE_GROUP" > /dev/null 2>&1; then
#     groupadd --system "AUTH_BASE_GROUP"
#     echo "Created group: AUTH_BASE_GROUP"
# else
#     echo "Group AUTH_BASE_GROUP already exists"
# fi

# Create directory if it doesn't exist
if [ ! -d "$DB_DIR" ]; then
    mkdir -p "$DB_DIR"
    echo "Created directory: $DB_DIR"
fi

# Create directory if it doesn't exist
if [ ! -d "$LOG_DIR" ]; then
    mkdir -p "$LOG_DIR"
    echo "Created directory: $LOG_DIR"
fi

#Ownership and gorups
chown -R "$SERVICE_USER:$SERVICE_USER" "$LOG_DIR"
chown -R "$SERVICE_USER:$SERVICE_USER" "$DB_DIR"

# Set proper permissions
chmod 757 "$DB_DIR" #change to 750 or 700 later
chmod 757 "$LOG_DIR" #change to 750 or 700 later

echo "Permissions set:"
echo "  $DB_DIR -> $SERVICE_USER:$SERVICE_USER (757)"  #change to 750 or 750 later
echo "  $LOG_DIR -> $SERVICE_USER:$SERVICE_USER (757)" #change to 750 or 750 later


install_go() {
    echo "Installing latest Go..."
    
    # Get latest Go version
    LATEST_GO=$(curl -s https://golang.org/VERSION?m=text)
    GO_URL="https://golang.org/dl/${LATEST_GO}.linux-amd64.tar.gz"
    
    echo "Downloading $LATEST_GO..."
    wget -q $GO_URL -O /tmp/go.tar.gz
    
    # Remove old installation
    sudo rm -rf /usr/local/go
    
    # Extract new version
    sudo tar -C /usr/local -xzf /tmp/go.tar.gz
    
    # Add to PATH for current session
    export PATH=$PATH:/usr/local/go/bin
    
    # Add to profile for future sessions
    if ! grep -q "/usr/local/go/bin" ~/.bashrc; then
        echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    fi
    
    # Clean up
    rm /tmp/go.tar.gz
    
    echo "Go installed: $(go version)"
}