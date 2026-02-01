# Check if 'air' is installed
if (Get-Command "air" -ErrorAction SilentlyContinue) {
    Write-Host "Starting with hot-reload (air)..." -ForegroundColor Green
    air
} else {
    Write-Host "Hot-reload tool 'air' not found." -ForegroundColor Yellow
    Write-Host "Falling back to standard 'go run .'" -ForegroundColor Yellow
    Write-Host "Tip: To enable hot-reload, run: go install github.com/air-verse/air@latest" -ForegroundColor Gray
    go run .
}
