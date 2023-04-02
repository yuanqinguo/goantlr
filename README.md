# goantlr
基于antlr4的golang实现的基础框架，支持常用运算，日期函数，支持业务内嵌函数扩展和其他常用函数扩展

## 技术背景

ANTLR (ANother Tool for Language Recognition) 是一个强大的文本或者二进制文件解析工具，被广泛应用于构建语法以及各种工具和框架。如

- Twitter’s search query language
- MySQL Workbench
- Hibernate
- Trino (SQL query engine)
- 等等

antlr可以对某一种语言描述，根据对应的语法，生成一颗语法树，在这颗语法树中包含了语言描述与对应语法规则的关系。当你去遍历这颗语法树时，可以灵活处理遍历前和遍历后的规则，实现某种效果。（可以理解为根据一组确定的语法规则，处理一段数据，如实现某种声明、运算、调用等，从而得到某种结果）

antlr可以用来做各种各样的事情，比如本文档的基于antlr4构建的规则引擎框架。 除此之外，还有海量的场景可以用到antlr，比如

- vscode中的插件，基于antlr做语法规则检查
- grafana中的查询语法解析
- 格式转换工具，如json、yaml、xml等互相转换
- 等等

## 相关资料

https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/

https://lsongseven.github.io/posts/antlr4/

https://github.com/bilibili/gengine/blob/main/README_zh.md

https://xie.infoq.cn/article/40bfff1fbca1867991a1453ac

## 业务背景

随着业务的不断发展和深入抽象，可以发现很多业务场景的功能代码都可以抽象成 规则+计算=>指标 的逻辑模式，这种模式可以应用于很多场景，如：

1. 考勤规则场景
   1. 如超过上班时间5分钟不算迟到，10分钟以上算迟到
   2. 如超过下班时间30分钟后，算加班等等
2. 运维监控场景
   1. 如在CPU使用率超过90%时，持续时长2分钟以上，触发告警
   2. 在CPU和内存使用率均超过90%时，持续时长2分钟以上，触发扩容
3. 薪酬计算场景
   1. 如某员工全勤时且该部门是有全勤奖的，发放全勤奖工资
   2. 如某员工在出差时需要发放补贴，但需要基于出差地点，岗位职级等发放不同规格的补贴金额
4. 等等

以上这些都是基于某类规则产生逻辑计算，再基于逻辑计算产生某个指标结果或这行某类动作->执行

**思考：**

在传统的产品和技术业务思维下，一般要满足以上的如考勤规则场景，首先会想到的是，放开选项配置或者输入，让用户决定多少分钟算迟到，下班时间多少分钟后算加班，这种方法，在大部分客户都是标准化或者统一化的情况下，可以满足很多客户需求，但如果某一天有个客户说： 上班5分钟不算迟到，仅对程序员这个岗位有效，对运营岗位没有5分钟，此时我们会继续加配置，或者继续基于一个考勤组分开配置，好像也可以接受，但下一个客户说，我们考勤组内还有外包和兼职，他们可以在30分钟内都不算迟到...以此类推，慢慢的产品的配置，下拉，勾选等配置越来越多，产品开始臃肿，代码开始维护困难。。。。

## 解决什么问题

基于以上，在大部分业务场景下，都是因为某些条件，所以要做另外一些事情，而这些条件对不同的用户，客户来说可能都不一样。在没有各类系统之前，大家都会使用excel进行各类数据的处理，比如考勤，基于表格处理后再基于excel中的函数进行以上规则配置或者lookup操作

该项目主要解决以下问题：

1. 基于Antlr4的词法，语法分析，提供自定义的规则配置，并将规则执行的结果返回
2. 实现常见的IF,IFS,SUM,AVG等计算规则
3. 支持基于该项目框架和Antlr4，快速扩展强业务耦合的规则语法函数

## 语法函数清单

### 判断

- IF  IF判断，如果为true，返回第一个，如果为false，返回第二个
- IFS  IFS多判断，按顺序判断，遇到任何一个为true的，返回后面的值
- OR  OR 满足条件1或者满足条件2
- AND  AND 满足条件1并且满足条件2
- \>   =  <  != >=  <=   大于，等于，小于， 不等于，大于等于， 小于等于

### 字符串类

- CONCATSTR   字符串连接 CONCATSTR("A", "B", "C") = ABC
- FINDSTR  查找字符串，返回位置FINDSTR("AB", "B")=1找不到返回0

### 数学类

- `+，-，×, ÷，%  加减乘除，求余`
- `MAX  取最大值 MAX(1,2,3,4) > 4`
- `MIN  取最小值 MIN(1,2,3,4) > 1`
- `SUM  求和 SUM(1,2,3,4) = 10`
- `AVERAGE  求平均值 AVERAGE(1,2,3,4) = 2.5`
- `ROUND  四舍五入 ROUND(1.45, 1) = 1.5` 

### 日期类

- `TODAY  TODAY() = 2022-01-01 当天日期`
- `DATE     生成日期DATE(1999, 7, 20) = 1999-07-20`
- `DATEDIF   日期差DATEDIF("1969-7-16", "1969-7-24", "D") = 8`
- `DATEADD  日期增加/减少天数  日期增加/减少月数  日期增加/减少年数`

### 扩展函数

- HELLO_WORLD()   HELLO_WORLD('参数1', '参数2')

可以基于语法树文件自行扩展

### 结构图

![img](blob:https://test-c21k1x5ri5hp.feishu.cn/49382a0b-cc45-4d44-a0e4-8f9098bf13ca)

### 对象说明

执行器

> 接收数据，执行规则语法，bing对外提供bizHandle注入和执行结果返回

- dataMap执行规则语法所需要的数据集，该数据集为map[string]interface{}结构
- funcCallIndexMap每个to函数被执行的索引值，基于visitor模式实现时，我们在每个函数调用中，基于上一个token的执行结果，判断是否进行语法树递归调用
- bizHandle 该对象为一个golang的interface{}对象定义，用于扩展业务函数实现逻辑
- funcResultMaps该对象记录了token，函数被执行时的结果值

解析器

> 主要解决用于在调用执行器之前，检查语法规则是否存在语法错误，或分析该规则中被引用的数据集变量标识，函数标识，和业务函数标识

- varsMap  被引用的数据集变量标识
- funcMap  被引用的token和函数标识， 该对象中不应包含业务函数
- bizFuncMap 被引用的业务函数标识

### 使用

#### 说明

1. 所有被引用的数据变量均基于 #xx# 进行包裹， # 字符为语法文件中定义，并在代码中硬性使用，不支持修改
2. 由于负号"-"和减号"-"在token字符中存在冲突，所以在规则公式中若需要输入-1号，需要使用()包裹

#### 代码示例

参考测试用例

```golang
uFunHandle := NewUFunSimple("name-hello")
calcVisitor := NewExecutor(uFunHandle, nil)

// 扩展函数示例
calcVisitor.CalculateForDebug("HELLO_WORLD('参数1', '参数2')")

calcVisitor.CalculateForDebug(`DATEDIF("2017-12-3","2015-10-7","MD")`)
calcVisitor.CalculateForDebug(`DATEDIF("2015-10-7","2017-12-9","MD")`)
calcVisitor.CalculateForDebug(`DATEDIF("2015-10-7","2017-7-3","MD")`)
calcVisitor.CalculateForDebug(`DATEDIF("2015-10-7","2017-7-12","MD")`)
calcVisitor.CalculateForDebug(`OR(31>3.1)`)
calcVisitor.CalculateForDebug("1+(-2)")
calcVisitor.CalculateForDebug("1-(-2)")
calcVisitor.CalculateForDebug("MAX(10, 99, 100, 102)")
calcVisitor.CalculateForDebug("MIN(10, 99, 100, 103)")
calcVisitor.CalculateForDebug("MAX(88, MAX(10, 99), 100)")
calcVisitor.CalculateForDebug("SUM(1 * 10, 1 + 2 - 1, 1)")
calcVisitor.CalculateForDebug("AVERAGE(10, 99, 100, 103)")
calcVisitor.CalculateForDebug("ROUND(10.1544, 1)")
calcVisitor.CalculateForDebug("DATEDIF(\"1993-04-01\",\"2022-09-15\",\"YD\")")
calcVisitor.CalculateForDebug(`FINDSTR("上海","深圳人民的夜晚，上海人民的上海滩")`)
```

### 最后
希望对有需要应用antlr4的同学有带来帮助