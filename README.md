#### TukSearch

这只是一个简易的, 基于 [bleve](http://www.blevesearch.com/), [ledisdb](http://ledisdb.com/) 和 [hprose](http://hprose.com/) 的全文检索玩具.

=====

做这个的初衷只是为了让我的小项目可以有全文检索这样高大上的功能, 而我又不打算去搭一套Elasticsearch 之类.  

=====

写这个玩具的大半时间花在踩坑以及造轮子上.  

比如代码中包含了 bleve 的一部分实现, 是因为:  

1.  bleve 的高亮 在处理 slice 的值时会出现一些问题, 我正在协助作者解决这里的问题.
2.  原生的 bleve 为每个 index 创建一个目录来保存 meta 文件

这是一个粗暴的做法, 我会尝试用其他办法解决这个问题, 而不是像这样"为了能用上,管他呢".  

当然, 这个玩具 "为了能用上" 还使用了一些其他不太好看的解决方式, 我会尝试一一修正.

=====

很高兴这个玩具依赖的三个项目有两个是国人开源的, 而即使在 bleve 中也使用了国人开源的 [结巴分词 Go 语言版](https://github.com/wangbin/jiebago)