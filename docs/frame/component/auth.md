## Oauth2 Server

### AuthorizeRequest
```go
userJson, err = json.Marshal(user)
if err != nil {
	return fmt.Errorf("AuthorizeRequest user json marshal error, err: %w",err)
}

ar := &server.AuthorizeRequest{}
ar = invoker.Sso.HandleAuthorizeRequest(server.AuthorizeRequestParam{
    ClientId:     reqView.Params.ClientId,
    RedirectUri:  reqView.Params.RedirectUri,
    Scope:        reqView.Params.Scope,
    State:        reqView.Params.State,
    ResponseType: reqView.Params.ResponseType,
})

ar.Build()
```

### AccessRequest

## SSO流程
* 第一步授权（浏览器）
  * client => server authorize?clientId=xxx&redirect_uri=xxx
  * client <= code <= server authorize 

* 第二步获取token（接口）
  * client => code => server getAccess
  * client <= access <= server getAccess   
  
* 第三步同域下，client刷新sub token（接口）
  * client获取cookie的token，如果获取不到
  * client获取cookie根域的parentToken  
  * client => parentToken => server refreshAccess
  * client <= access <= server refreshAccess 
  
* 第三步同域下，server续期parent token（接口） 
  * server获取cookie的parentToken
  * server renew, parentToken 增加租期