sudo snap install golang --classic
sudo apt install nginx -y
sudo nano /etc/nginx/nginx.conf # Add the following to the end of the file
sudo nano /etc/nginx/sites-enabled/default
sudo nano /etc/security/limits.conf

ulimit -Sn
ulimit -Hn
git clone https://github.com/HarishTeens/todos-app.git
cd todos-app
go mod tidy
nano .env
go run *.go
PORT=8080 nohup go run *.go &
PORT=8081 nohup go run *.go &
PORT=8082 nohup go run *.go &
PORT=8083 nohup go run *.go &
PORT=8084 nohup go run *.go &
PORT=8085 nohup go run *.go &
PORT=8086 nohup go run *.go &
PORT=8087 nohup go run *.go &
PORT=8088 nohup go run *.go &
PORT=8089 nohup go run *.go &
PORT=8090 nohup go run *.go &
PORT=8091 nohup go run *.go &
PORT=8092 nohup go run *.go &
PORT=8093 nohup go run *.go &
PORT=8094 nohup go run *.go &
PORT=8095 nohup go run *.go &

PORT=8096 nohup go run *.go &
PORT=8097 nohup go run *.go &
PORT=8098 nohup go run *.go &
PORT=8099 nohup go run *.go &
PORT=8100 nohup go run *.go &
PORT=8101 nohup go run *.go &
PORT=8102 nohup go run *.go &
PORT=8103 nohup go run *.go &
PORT=8104 nohup go run *.go &
