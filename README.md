```
docker-compose up -d

go run main.go -command insert | sort | uniq -c | grep -vP '^\s+1'

go run main.go -command update | sort | uniq -c | grep -vP '^\s+1'
```