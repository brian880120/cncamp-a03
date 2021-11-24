# cncamp-a02

作业要求:

1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回 200

运行:

make dep: 下载依赖包

make run: 运行程序

API测试:

测试环境: 本地
测试工具: postman

endpoints:

对于所有endpoints的请求，返回结果应包含:

	response headers:
	- 应包含request header的所有field，并包含Version field

	server端日志:
	- 包含客户端IP, HTTP请求状态以及请求方式

endpoint测试

1. localhost:8090/healthz GET

	结果:

	response body:

	{
		"message": "Status OK"
	}

2. localhost:8090/badRequest GET

	response body:

	{
		"message": "Bad Request"
	}

	server端日志应包含: Request Status: Bad Request

3. localhost:8090/notFound GET

	response body:

	404 page not found

	server端日志应包含: Request Status: Not Found

4. localhost:8090/mockSignup POST

	request header: "Content-Type": "application/json"

	request body:

	{
		"name": "test",
		"email": "test@mail.com"
	}

	结果:

	response body:

	{
		"name": "test",
    		"email": "test@mail.com"
	}

