# 统一鉴权
## 鉴权流程
鉴权流程主要分以下三个过程
* **过程一**：`account`服务在处理`/login`接口时会生成一个关于该登录用户账号的`token`，将其存储到后端`redis`中。
* **过程二**：鉴权服务(`account-rpc`)的提供一个rpc接口来对其他服务传入的参数(`accountName`+`token`)来在`redis`服务端进行查询，返回成功或者失败，成功表示通过鉴权，失败表示没通过鉴权
* **过程三**：其他服务，进行后端请求时都需要上传一个(`accountName`+`token`)，然后通过`account-rpc`服务完成鉴权流程。通过才允许访问其要请求的资源，不通过则不允许其要访问的资源。

以下将对上面三个过程进行详细介绍

### 过程一
登录接口(`/login`)的原过程如下：
```go
func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
        encryptedPassword := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password)
        record, err := l.svcCtx.TbUserAccountModel.FindOne(l.ctx, req.AccountName)

        resp = new(types.LoginResp)
        if err != nil {
            resp.Result = _const.ApiFailed
            resp.Message = fmt.Sprintf("login failed, err: %v", err)
            return
        }

        if record.Password.String != encryptedPassword {
            resp.Result = _const.ApiFailed
            resp.Message = fmt.Sprintf("login failed, password wrong")
            return
        }

        tokenData := TokenData{
            AccountName: req.AccountName,
            Password:    req.Password,
        }
        accessToken, err := token.GenerateToken(l.svcCtx.Config.TokenSecretKey, tokenData, TokenExpireTime)
        if err != nil {
            resp.Result = _const.ApiFailed
            resp.Message = fmt.Sprintf("token generate failed, err: %v", err)
            return
        }

        resp.Result = _const.ApiSuccess
        resp.Token = accessToken
        return
    }    
```

这里我们需要引入`redis`,具体实现拆分成以下几个过程。
1. 在`Docker`环境中部署`redis`：具体配置详见`docker-compose.yml`中关于`redis`相关部分的配置
2. 抽出一个公共模块用于`redis`数据的存储和获取，详见`/pkg/store/redis.go`中的实现
3. `redis`的配置,详见`/cmd/account/etc/account-api.yaml`中关于`redis`部分的配置 
4. 在`account`服务的启用`redis`，详见`cmd/account.go`中`store.MustUseRedisStore(c.RedisConf)`的调用和`cmd/account/internal/config/config.go`中的`redis.RedisConf`
5. 在登录逻辑中调用`redis`的存储逻辑，详见`cmd/account/internal/logic/loginlogic.go`中关于`redis`部分的逻辑，即` err = store.Set(req.AccountName, accessToken, 3600*15)`的调用

### 过程二
1. 配置远程`redis`的连接参数
2. 修改`cmd/account-rpc/internal/logic/validatetokenlogic.go`文件中`ValidateToken`函数即可。主要的功能就是：
* 从`redis`中拿`token`数据
* 验证该`token`是否和客户端上传的数据匹配即可

### 过程三
重新写一个叫做`userinfo-api`的服务，端口地址暂时可以定为`8006`，该服务的主要功能就是用来修改某账号的用户信息。该服务本身不和`account-rpc`服务进行通信。这里暂时随便提供一个叫做`/hello`的接口，就用来向客户端返回字符串`Hello, visitor!`。

前端的访问逻辑是通过访问地址:`http://localhost:8888/userinfo/hello`来访问。

在前面文档`06-gateway-config-add-nginx-proxy.md`中说过，我们打算用`nginx`来对后端服务进行一个代理,而`8888`端口正是`Docker`环境下，提供给外部客户端的访问容器环境`nginx服务`的端口。而`/userinfo/hello`路径会被`nginx服务`，转发到我们这里的`userinfo-api`服务，通过在`nginx`的配置文件中, 我们在请求由`userinfo-api`服务处理之前先将请求转发给`account-rpc`服务进行一个身份鉴权，身份鉴权成功才继续交由`userinfo-api`服务处理，身份鉴权失败，则直接返回给客户端。那么也就是说，如果鉴权成功的话，客户端访问`http://localhost:8888/userinfo/hello`在响应中会得到字符串`Hello, visitor!`。

具体在`Docker`环境中如何部署`userinfo-api`参照文章`06-account-docker-image-build.md`。

