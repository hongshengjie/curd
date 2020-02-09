# curd
自动生成curd 代码


## 目标
1. 根据表名,字段名 自动生成 golang 结构体，避免手写
2. 根据主键id生成 查询，更新，删除，插入 golang curd代码，避免手写（非常枯燥）
3. 根据表索引生成单个查询，或者是批量查询代码
4. 不要求生成高级的查询，
5. 如果有自定义复杂的需求，拷贝一下前面步骤生成的代码，更改其中的方法名，参数，sql语句即可满足


## 优点
1. 解放双手，避免手写无聊的一样的代码
2. 快速开发，灵活拷贝定制
3. 统一代码风格，所有项目的curd代码风格都是一致，防止不同的人用不同的orm，或者奇怪的代码风格导致交接维护困难

## 限制
1. 不支持复合主键

 
## 安装

```

git clone https://github.com/hongshengjie/curd.git

cd  curd 

go install 

```

## 使用方法

```

curd  -dsn='root:1234@tcp(127.0.0.1:3306)/my-app?parseTime=true'  -schema=my-app -table=all_type_table | gofmt 

curd  -dsn='root:1234@tcp(127.0.0.1:3306)/my-app?parseTime=true'  -schema=my-app -table=all_type_table -tmpl=table_bilibili.tmpl | gofmt  > example/bilibili/all_type_table.go

```


