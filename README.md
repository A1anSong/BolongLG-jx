# 江西省公共资源交易平台金融服务系统金融机构侧接入指引（Jiangxi Letter of Guarantee）

# Nginx相关配置
配置文件修改内容，bt面板可通过伪静态设置：
```nginx
location /api {
  proxy_set_header Host $http_host;
  proxy_set_header  X-Real-IP $remote_addr;
  proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
  proxy_set_header X-Forwarded-Proto $scheme;
  rewrite ^/api/(.*)$ /$1 break;  #重写
  proxy_pass http://127.0.0.1:9999; # 设置代理服务器的协议和地址
}
```

设置vue-router为`createWebHistory`(去掉Hash模式下的路径'#'号)，会遇到页面刷新404，需要添加nginx配置：
```nginx
location / {
  try_files $uri $uri/ /index.html;
}
```

需要用gin来serve静态文件参考将`NoRoute`设置到`index.html`路径，embed同理
```golang
ginRouter.NoRoute(func(c *gin.Context) {
		c.File("../web/dist/index.html")
	})
```

二级目录刷新空白的bug，修改`vite.config.js`：
```javascript
return {
    base: '/',
}
```

# 开发备注
gva版本 `v2.5.5` @ Commit Hash（SHA）：[0ef9aac](https://github.com/flipped-aurora/gin-vue-admin/tree/0ef9aac32e1cc966476cd03b8138c32b4b50ce54)
