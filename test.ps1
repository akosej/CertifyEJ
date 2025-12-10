# Script para generar un certificado de prueba rápido
# Ejecuta este script con: .\test.ps1

Write-Host "==================================" -ForegroundColor Cyan
Write-Host "  Test del Generador de Certificados" -ForegroundColor Cyan
Write-Host "==================================" -ForegroundColor Cyan
Write-Host ""

# Verificar que existe el ejecutable
if (!(Test-Path "generador-certificados.exe")) {
    Write-Host "ERROR: No se encuentra generador-certificados.exe" -ForegroundColor Red
    Write-Host "Ejecuta primero: go build -o generador-certificados.exe" -ForegroundColor Yellow
    exit 1
}

# Ejecutar con el ejemplo simple
Write-Host "Generando certificado de prueba..." -ForegroundColor Green
.\generador-certificados.exe -json ejemplo-simple.json -output certificados_test

if ($LASTEXITCODE -eq 0) {
    Write-Host ""
    Write-Host "✓ Certificado generado exitosamente!" -ForegroundColor Green
    Write-Host ""
    Write-Host "Revisa la carpeta: certificados_test" -ForegroundColor Yellow
    Write-Host ""
    
    # Abrir la carpeta de salida
    $respuesta = Read-Host "¿Deseas abrir la carpeta de certificados? (s/n)"
    if ($respuesta -eq "s" -or $respuesta -eq "S") {
        Start-Process "certificados_test"
    }
} else {
    Write-Host ""
    Write-Host "✗ Hubo un error al generar el certificado" -ForegroundColor Red
}
