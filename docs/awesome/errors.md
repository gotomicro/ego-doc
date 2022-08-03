# 错误处理
## 要点
* Error需要常量，不能是变量
* Error需要wrap，方便排查问题
* Error需要用IS判断，不能用=判断
* Error需要收敛

## Error是常量
第一个问题是io.EOF公共变量-导入io包的任何代码都可能更改的值io.EOF。事实证明，在大多数情况下，这并不是什么大问题，但可能数据被人篡改，引发不必要的问题。
```go
fmt.Println(io.EOF == io.EOF) // true
x := io.EOF
fmt.Println(io.EOF == x)      // true
	
io.EOF = fmt.Errorf("whoops")
fmt.Println(io.EOF == io.EOF) // true
fmt.Println(x == io.EOF)      // false
```

正确的用法，应该如下所示
```go
const eof = Error("eof")

func (r * Reader) Read([] byte) (int, error){ 
        return 0, eof 
} 

func main () { 
        var r Reader 
        _,err: = r.Read([] byte {})
        fmt.Println(err == eof)// true 
}
```

## Error需要wrap
GO1.13支持了error wrap。我们可以在错误以下方法，将原始错误进行包装。fmt.Errorf里是%w
```
err = fmt.Errorf("wrap error %w", err)
```
这里需要提醒一点，go官方的error wrap没有堆栈信息，还是比较坑爹

## Error需要IS
以往我们对错误判断都是=，但是如果使用了wrap，在用=是无法相等的：
```
selectErr := fmt.Errorf("select info err: %w", gorm.IsNotRecord)
fmt.Println(selectErr == gorm.IsNotRecord) // false
fmt.Println(errors.Is(selectErr, gorm.IsNotRecord)) // true
```

## Error需要收敛

## Error说明
目前官方error没有支持堆栈，可能使用pkg/errors排查问题更方便。
但ego为了支持官方后续升级，还是决定使用官方error用法。

# 引用文献
* [常量error](https://dave.cheney.net/2016/04/07/constant-errors)
* [1.13 Error Wrap深度分析](https://www.cnblogs.com/sunsky303/p/11571440.html)
