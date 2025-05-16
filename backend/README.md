```dotenv
HOST=0.0.0.0
PORT=8080
SECRET_KEY=SECRET_KEY
ALLOW_ORIGINS=http://localhost:5173
JWT_SECRET=JWT_SECRET1234
DATABASES=ocserv-test.db
```

```bash
env $(cat .env | xargs) go run cmd/main.go

swag init && go run main.go -debug -migrate && go run main.go -debug

# test
go install github.com/vektra/mockery/v2@latest

mockery --all --output=./mocks --outpkg=mocks --recursive

```

# Tips
```bash
ocpasswd -c /etc/ocserv/ocpasswd USERNAME

sudo docker exec -it ocserv_api bash 

echo -e "1234\n1234\n" | ocpasswd -c /etc/ocserv/ocpasswd test
```

```text
worker[test]: 172.17.0.1 worker-auth.c:1731: failed authentication for 'test'

main[test]:172.17.0.1:55064 user logged in

main[test]:192.168.100.21:41448 user disconnected (reason: user disconnected, rx: 483392, tx: 2330)

main[test]:172.17.0.1:56906 user disconnected (reason: server disconnected, rx: 378308, tx: 944)
```