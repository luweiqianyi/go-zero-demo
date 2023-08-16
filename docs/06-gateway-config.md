# 统一网关
给该项目下的微服务设置一个统一网关，由该网关完成统一鉴权，或者同一功能处理。
这里暂时使用`nginx`做统一网关处理。详细步骤见如下章节。
## docker中部署nginx
### docker compose方式部署
本项目不提供`docker run`的方式来下载，构建容器的方式。很笨。
1. `docker-env`目录下新增`docker-compose.yml`文件，内容如下
    ```yml
    version: '3'
    services:
    nginx:
        image: nginx:1.24.0
        container_name: nginx-1.24.0
        ports:
        - 8888:8881
        volumes:
        - ./nginx.conf:/etc/nginx/nginx.conf
    ```
2. 部署
    ```shell
    docker-compose -p go-zero-demo up -d
    ```