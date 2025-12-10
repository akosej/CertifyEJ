@echo off
REM Script de ayuda para generar certificados
echo =========================================
echo  Generador de Certificados - Menu
echo =========================================
echo.
echo 1. Generar certificados del archivo por defecto (certificados.json)
echo 2. Generar certificados del ejemplo simple
echo 3. Ver ayuda
echo 4. Salir
echo.
set /p opcion="Seleccione una opcion (1-4): "

if "%opcion%"=="1" goto opcion1
if "%opcion%"=="2" goto opcion2
if "%opcion%"=="3" goto opcion3
if "%opcion%"=="4" goto salir

:opcion1
echo.
echo Generando certificados desde certificados.json...
generador-certificados.exe
pause
goto fin

:opcion2
echo.
echo Generando certificados desde ejemplo-simple.json...
generador-certificados.exe -json ejemplo-simple.json -output certificados_ejemplo
pause
goto fin

:opcion3
echo.
generador-certificados.exe -h
pause
goto fin

:salir
echo.
echo Saliendo...
exit

:fin
echo.
echo Proceso completado.
pause
