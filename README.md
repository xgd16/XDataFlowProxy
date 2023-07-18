# XDataFlowProxy

[生成KEY](https://github.com/xgd16/getKey/releases/tag/v1.0.0)

### 功能

1. 代理服务 可实现单用户对于指定地址按照设置规则进行顺序请求
2. 快速接入服务请求信息统计 (Prometheus)


### 在使用宝塔时结构为 ``nginx -> XDFP -> nginx -> php-fpm`` 请确保 ``proxy_set_header Host`` 使用 ``$host``

```nginx configuration
location / {
    proxy_pass http://127.0.0.1:9431;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header REMOTE-HOST $remote_addr;
    add_header X-Cache $upstream_cache_status;
    proxy_set_header X-Host $host:$server_port;
    proxy_set_header X-Scheme $scheme;
    proxy_connect_timeout 30s;
    proxy_read_timeout 86400s;
    proxy_send_timeout 30s;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";
}
```