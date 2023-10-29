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


//创建数据的时候应该验证是否已经有了

//把表拆成动态数据表和静态数据表，一个是数据经常变的一个是数据不怎么变的。一次修改多个表就需要考虑事务

//github添加协作人

//解决id不连续的问题


/**
//点线连接问题，选点线的时候需要判断电线是否存在，则需要查询点线，但是现在不能查电线
//解决方案，在invoker层用GRPC调用LLS的接口查询返回
**/
