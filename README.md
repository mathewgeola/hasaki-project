# hasaki-project

## 部署

```bash
go mod init github.com/mathewgeola/hasaki-project

go get -u github.com/spf13/cobra

go build -o hp
```

## 发布

```bash
make install && hp vcommit && hp tag && hp push
```
