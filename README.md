# gee_study
实现egine引擎，让egine实现ServeHTTP,从而实现LAS的handler接口 
将router包分离,使得engine封装router  
封装了各种中间方法和请求信息的Context随请求而诞生，在ServeHTTP中新建  
路由的注册由router实现，并且router实现路由映射，并返回不存在的404  
