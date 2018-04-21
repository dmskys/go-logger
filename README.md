### go-logger 是golang 的日志库 ，基于对golang内置log的封装。
用法类似java日志工具包log4j

**打印日志有5个方法 Debug，Info，Warn, Error ,Fatal  日志级别由低到高**

设置日志级别的方法为：logger.SetLevel() 如：logger.SetLevel(logger.WARN)
则：logger.Debug(....),logger.Info(...) 日志不会打出，而 
 logger.Warn(...),logger.Error(...),logger.Fatal(...)日志会打出。
设置日志级别的参数有7个，分别为：ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF
其中 ALL表示所有调用打印日志的方法都会打出，而OFF则表示都不会打出。

***

日志文件切割有两种类型：1为按日期切分。2为按日志大小切分。
按日期切分时：每天一个备份日志文件，后缀为 .yyyy-MM-dd 
过0点是生成前一天备份文件

按大小切分是需要3个参数，1为文件大小，2为单位，3为文件数量
文件增长到指定限值时，生成备份文件，结尾为依次递增的自然数。
文件数量增长到指定限制时，新生成的日志文件将覆盖前面生成的同名的备份日志文件。


### 根据原日志修改为简单全局的设置，并且简单的调用

**示例**：

    //--------------------- 此段代码可在配置文件里配置全局日志设置
  	//设置log目录
  	GOPATH := os.Getenv("GOPATH")
  	fmt.Println("GOPATH is ", GOPATH)
  	LogDir := GOPATH + "/bin/logs"
  	fmt.Println("logdir is ", LogDir)
  	//设置全局日志目录（必须设置全局日志目录）
  	logger.LogDir = LogDir
  	//设置全局日志级别 (默认不设置为 DEBUG 模式)
  	logger.LogLevel = logger.DEBUG
  	//设置全局log最大文件个数 (默认不设置为 10)
    logger.LogMaxNumber = 10
    //设置全局log最大文件size (默认不设置为 10)
    logger.LogMaxSize = 10
    //设置全局log最大文件size的单位计数 (默认不设置为 MB)
    logger.LogUnit = logger.MB
     //---------------------


    //在具体的代码实现调用 方法与原来的代码一致
  	//获取日志对象，并指定存取的文件
    log := logger.GetLogger("config.log")
    //使用Error
    log.Error("load config is error! errorinfo is ", err)
    //使用Debug
    log.Debug(fmt.Sprintf("config load succeessfully:%v", config))



