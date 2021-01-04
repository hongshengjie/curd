# crud

自动生成增删改查代码


## 目标
1. 根据表名,字段名 自动生成 golang 结构体，避免手写
2. 根据主键id生成 查询，更新，删除，插入 golang crud代码，避免手写（非常枯燥）
3. 支持更新指定字段
4. 灵活创建where条件语句
5. 支持事务
6. 支持 mariadb,mysql5, mysql8


## 限制
1. 不支持复合主键

 
## 安装

```

git clone https://github.com/hongshengjie/crud.git

cd  crud/crud 

go install 

```

## 使用方法

```

crud  -dsn='root:1234@tcp(127.0.0.1:3306)/example?parseTime=true'   -table=user 

crud  -dsn='root:1234@tcp(127.0.0.1:3306)/example?parseTime=true' -table=all_type_table
```
