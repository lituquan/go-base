# Simple Full-Text Search engine
## 参考
Sample code for https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine/.

## 设计
    代码搜索
### 1.扫描指定文件夹
        读取指定后缀文件
        写入倒排文件和索引表
        （这里用内存查询，为了保存加载的索引，json格式化写入文件）

        分词部分用的
### 2.搜索关键词
    http://127.0.0.1:8080/
    http://127.0.0.1:8080/files/all
    http://127.0.0.1:8080/search/:key
### 3.暴露接口
    使用gin框架
