@echo off
cd /d E:\woogle

docker compose up --build -d

echo.
echo WOOGLE STACK ONLINE
echo Gateway: http://localhost:8080/health
echo Dashboard: http://localhost:3000
echo Cortex: http://localhost:8090/health
pause