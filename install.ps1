# Configuration
$INSTALL_DIR = "$env:ProgramFiles\tchat"  # Standard installation location
$CONFIG_DIR = "$HOME\.config\tchat"
$CLIENT_DIR = "./client" # Directory containing the client source code

# Utils
func Write-ColoredOutput {
    param(
        [Parameter(Mandatory)]
        [ValidateSet('info', 'error')]
        $Type

        [Parameter(Mandatory)]
        $Message
    )

    $Color = switch ($Type) {
        'info' { 'Green' }
        'error' { 'Red' }

        Write-Host "tChat >> " -ForegroundColor $Color -NoNewline
        Write-Host $Message
    }
}

# Error Handling
function handle_error {
    Write-ColoredOutput "error" "Installation failed. Please check the error messages above."
    exit 1
}


# Check OS (PowerShell is Windows-based)
Write-ColoredOutput "info" "OS check is implicit on PowerShell..."

Write-ColoredOutput "info" "Checking for git..."
if (-not (Get-Command git -ErrorAction SilentlyContinue)) {
    Write-ColoredOutput "error" "git could not be found. Please install it ."
    exit 1
}

Write-ColoredOutput "info" "Checking for go..."
if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-ColoredOutput "error" "go could not be found. Please install it."
    exit 1
}

# Main Installation Logic
Write-ColoredOutput "info" "Building client..."
Push-Location $CLIENT_DIR  # Equivalent to cd in PowerShell
try {
    go build -o $CLIENT_DIR\bin\tchat $CLIENT_DIR\main.go
    Write-ColoredOutput "info" "Build successful!"
} finally {
    Pop-Location
}

Write-ColoredOutput "info" "Installing tchat..."

Write-ColoredOutput "info" "Creating configuration directory: $CONFIG_DIR"
New-Item -ItemType Directory -Force -Path $CONFIG_DIR | Out-Null

Copy-Item -Path .\bin\tchat.exe -Destination $INSTALL_DIR 

Write-ColoredOutput "info" "Installation successful!"
