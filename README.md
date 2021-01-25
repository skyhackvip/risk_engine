# 风控决策引擎系统
### 决策引擎系统介绍

风控决策引擎系统是在大数据支撑下，根据行业专家经验制定规则策略、以及机器学习/深度学习/AI领域建立的模型运算，对当前的业务风险进行全面的评估，并给出决策结果的一套系统。

决策引擎，常用于金融反欺诈、金融信审等互金领域，由于黑产、羊毛党行业的盛行，风控决策引擎在电商、支付、游戏、社交等领域也有了长足的发展，刷单、套现、作弊，凡是和钱相关的业务都离不开风控决策引擎系统的支持保障。决策引擎和规则引擎比较接近（严格说决策引擎包含规则引擎，之前也有叫专家系统，推理引擎），它实现了业务决策与程序代码的分离。

关于如何实现决策引擎的文章市面极少见，实践生产落地的经验分享也基本没有。我会结合工作实践及个人思考，从业务抽象建模，产品逻辑规划以及最终技术架构和代码实现等方面给出全方位的解决方案。

### 开源声明
本项目用于学习和参考，不能直接用于生产环境。代码会不定期迭代更新，可以加关注定期查看。如有交流欢迎加微信号留言。

### 服务测试
- go build engine.go
- ./engine

请求接口：
```json
curl -XPOST  -v  http://localhost:8889/run -d'{"flow":"flow_conditional","features":{"feature_1":18,"feature_2":30,"feature_3":20,"feature_4":30}}' -H'Context-Type:application/json'
```
flow:决策流，存储在test/yaml中

接口返回：
```json
{"flow":"flow_conditional","result":{"NextNodeName":"","NextCategory":"","Decision":null,"Track":["start_1","conditional_1","ruleset_2","end_2"],"Detail":[{"NodeName":"conditional_1","Factor":{"feature_4":{"Name":"feature_4","Type":0,"Value":30,"Default":null}},"Hits":null,"Decision":"ruleset_2"},{"NodeName":"ruleset_2","Factor":{"feature_1":{"Name":"feature_1","Type":0,"Value":18,"Default":null},"feature_2":{"Name":"feature_2","Type":0,"Value":30,"Default":null}},"Hits":null,"Decision":0}]}}
```


### 决策引擎架构图
![决策引擎架构图](https://i.loli.net/2021/01/21/bOR1tyVPnCZNGoi.png)

### 代码解读
[智能风控决策引擎系统可落地实现方案（一）规则引擎实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483738&idx=1&sn=111609f176f11de8357c51a820b089b5&chksm=e8215e4adf56d75c2e6e8b81b89c1faabab667f493ce809cb749994cc9cd776342fd17d4172e&token=227666410&lang=zh_CN#rd)

[智能风控决策引擎系统可落地实现方案（二）决策流实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483770&idx=1&sn=3166a6617ddb6b628261b8b7ff84cfac&chksm=e8215e6adf56d77cb76de41b63e63759221932f030e315acebbc4025939b2e02b354a9072ecc&scene=178#rd)

[智能风控决策引擎系统可落地实现方案（三）模型引擎实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483789&idx=1&sn=ddb5f31edfd3174d4551fecc3f120f42&chksm=e8215e9ddf56d78b520f7ab5c8db7e978b3078a1e2511d424ff272ac6c509fd4c13d893dfc09&token=1795265687&lang=zh_CN#rd)

[智能风控决策引擎系统可落地实现方案（四）风控决策实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483825&idx=1&sn=3ebf7c8ad42f870e48db56ca6bb99ade&chksm=e8215ea1df56d7b7d9b1c653c61ef011d72d46d090845d91deba39f635d03ce1282eaa433485&token=1795265687&lang=zh_CN#rd)

[智能风控决策引擎系统可落地实现方案（五）评分卡实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483860&idx=1&sn=45bfbf4e436001dc060d5d4718688e9b&chksm=e8215ec4df56d7d2396c6024b49fc67eb25ee5754da9ddd40365f72abd5c1535a45218ea79b1&token=1239858205&lang=zh_CN#rd)

[智能风控决策引擎系统可落地实现方案（六）风控监控大盘实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483882&idx=1&sn=cb1142ea342b03f2f4ada44383e4bcbe&chksm=e8215efadf56d7ecae2159b7f742678d6036e6df046513ccce0efb052029d13b4c7b67ae1bc6&token=290046129&lang=zh_CN#rd)


扫码关注微信订阅号支持：

![技术岁月](https://i.loli.net/2021/01/21/orQm9BUkEqKAR6x.jpg)
