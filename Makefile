# 运行后端
server:
	 go run main.go

# 运行前端
browser:
	docker run  --name nginx -d -p 80:80  -v $(CURDIR)/nginx.conf:/etc/nginx/nginx.conf -v $(CURDIR)/view/:/undercover/ nginx

# 开始
# http://localhost/undercover
