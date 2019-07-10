# usersvr
用户服务  golang学习
项目参考goa[官方example](https://github.com/goadesign/examples)学习框架使用

### 使用方法
#### 1.环境 golang 1.12+
#### 2.安装三方库
```
cd usersvr
go get -u goa.design/goa/v3
go get -u goa.design/goa/v3/...
go mod download
```
#### 3.在项目根目录，创建配置文件config.yaml,内容如下：
```
server:
  http: "http://127.0.0.1:8000"
  grpc: "grpc://localhost:8080"
  debug: true

postgresql:
  host: "127.0.0.1"
  port: "5432"
  user: "iot"
  dbname: "iot"
  password: "123456"
  sslmode: "disable"

redis:
  addr: "localhost:6379"
  password: ""
  db: 0

email:
  user: "your@email.com"
  pass: "password"
  host: "smtp.email.com"
  port: "465"

jwt: 
  secret: "secret"
```

#### 4.编译并运行
```
cd cmd/usersvr
go build
./usersvr
```

#### 5.注意
- 本项目使用的框架是goa,如需要请前往[gos官方文档](https://goa.design/)查看帮助
