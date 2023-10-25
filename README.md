//目录介绍
一、cmd:项目启动目录

    1、admin:运维配置接口，与前端交互
    2、api:业务服务接口
    3、task:异步任务服务
二、config:配置阿波罗的文件目录

三、deploy:部署文件目录

四、initialize:初始化文件目录

五、internal:内部的

    1、configuration:代码中的配置文件
    2、constant:常量定义
    3、dal:数据访问层
        invoker:调用外部服务
        repositry:放DB连接数据
    4、dto:数据传输对象
    5、infrastructure:基础设施层
    6、usecase:相当于service层
    7、domain:业务逻辑层
    8、util:工具层
六、apps:应用入口层（路由层，接收请求）

七、protocol:协议包层，用于定义协议
    
