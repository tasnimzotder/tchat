# Configuration (Adjust as needed)
$INSTALL_DIR = "C:\Program Files\tchat" 
$CONFIG_DIR = "$env:USERPROFILE\.config\tchat"
$CLIENT_DIR = "./client"  # Directory containing the client source code
$RELEASE_URL = "https://github.com/tasnimzotder/tchat/releases/latest/download"
$ASSET_WINDOWS_AMD64 = "tchat.windows.amd64.exe" 

# ------------------------
#  Functions (like Bash 'print')
# ------------------------

function Write-ColoredOutput($Type, $Message) {
    switch ($Type) {
        "info"  { Write-Host "tChat >> " -ForegroundColor Cyan $Message }
        "error" { Write-Host "tChat >> " -ForegroundColor Red $Message }
        Default { Write-Host "tChat >> " -ForegroundColor Green $Message }
    }
}

# ------------------------
#  "Error Handling"
# ------------------------

$ErrorActionPreference = "Stop"  # PowerShell's equivalent to 'set -e'

# ------------------------
#  Check OS 
# ------------------------

Write-ColoredOutput "info" "Checking OS..."

if (-not ($env:OS -eq "Windows_NT")) {
    Write-ColoredOutput "error" "This script is primarily intended for Windows systems."
    exit 1
}

# ------------------------
#  Main Installation Logic
# ------------------------

Write-ColoredOutput "info" "Installing tchat..."

# Create installation directory if needed
if (-not (Test-Path $INSTALL_DIR)) {
    New-Item -Path $INSTALL_DIR -ItemType Directory | Out-Null
}

# Download the binary (assuming there's a Windows-compatible asset)
Invoke-WebRequest -Uri ($RELEASE_URL + "/" + $ASSET_WINDOWS_AMD64) -OutFile "$INSTALL_DIR\tchat.exe" 

Write-ColoredOutput "info" "Installation successful!"
