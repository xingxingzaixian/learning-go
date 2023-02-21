使用Go实现nginx静态Web服务器和文件服务器

### [Gin实现静态Web服务器](code/ginStatic/main.go)

- Static：静态Web服务器，目录下如果包含index.html，将作为静态网页进行访问。非常简单的就实现了nginx静态服务器的功能。
- StaticFS：静态文件服务器，相当于Web目录访问

需要注意的是Gin会有路由冲突，因此Static不能使用/作为路由，所以为了可以通过根目录访问，就需要进行根路由转发。

### [Http实现静态Web服务器](code/httpStatic/main.go)
使用原生http服务器实现静态文件服务器需要注意文件的访问需要使用`http.StripPrefix`进行前缀处理，前缀必须和路由相同，才能正常访问。
