# eth scanner
## Specify path to ethereum (optional)
If you already have ethtest synced on your PC - specify path to .ethereum folder
By default it is $HOME/.ethereum

```
vim docker-compose
```

## Run geth and scanner
```
# get vendors
make vendor
# build bin
make build
# run db and init database structure
make initdb
# run
docker-compose up -d
```

### Use API

```
GET http://localhost:8181/getLast
```

```
curl -d '{ "from":"0x4ddf45717d2a95fc005d4b33c21606f8524c10d4", "to":"0xccc8d252b6b6f14f4df66ab42e718ca1d5668d7d", "amount": 0.000000000000000001 }' -H "Content-Type: application/json" -X POST http://localhost:8181/sendEth
```