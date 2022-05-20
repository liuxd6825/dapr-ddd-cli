# dapr-ddd-cli

#### 目标
框架目标是简化DDD开发难度，使开发人员可直接进行业务开发，不需关心技术细节与实现。实现技术与业务分离，提升开发效率与质量。

#### 介绍

dapr ddd 命令行脚手架工具，通过定义DDDML文件，可以快速生成DDD业务代码。


#### 使用说明


1. 新建目录，在其中定义DDDML文件。可定义多个，脚手架会自动合并文件内容。
2. 执行dapr-ddd-cli脚手架。\


    $ dapr-ddd-cli init -model ./dddml -lang go -out . 

   
   - 参数 -model： dddml模型目录。\
   - 参数 -lang： 要生成代码的开发语言 可选go/java/csharp。目前仅支持go\
   - 参数 -out：生成代码的存放位置。


#### DDDML 示例


1. 请参见 DDDML目录.yaml文件 
2. DDDML的语法定义参考了《深入实践DDD以DSL驱动复杂软件开发》一书。


#### 参与贡献


1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
