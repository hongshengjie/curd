# crud

自动生成MySQL增删改查代码

## 安装

```
 go get github.com/hongshengjie/crud/crud
```

## 使用方法

```

crud  -dsn='root:1234@tcp(127.0.0.1:3306)/example?parseTime=true'   -table=user 

crud  -dsn='root:1234@tcp(127.0.0.1:3306)/example?parseTime=true' -table=all_type_table
```
