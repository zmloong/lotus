package core

type S_Category string //服务类别 例如 网关服务 游戏服务 业务服务   主要用于服务功能分类
type M_Modules string  //模块类型
type S_Comps string    //服务器组件类型
type ErrorCode int32   //错误码
type Event_Key string  //事件Key
type Rpc_Key string    //RPC
type Redis_Key string  //Redis缓存
type SqlTable string   //数据库表定义
type CustomRoute uint8 //自定义网关
