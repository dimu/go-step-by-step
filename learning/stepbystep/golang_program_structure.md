
## golang names

go命名规则: 以字母或者下划线开始，剩余部分为数字，字母或者下划线的组合。

go语言命名对大小写敏感，例如：countNum与CountNum为不同命名。

### go keyword（total：25）

``` 
break default func interface select case defer go map struct chan else goto package switch
const fallthrough if range type continue for import return var  
```

其中比较少见的为: **select defer go chan fallthrough range type**

### go predeclared names

预声明的命名，有点类似C语言的保留字，如下：

 * Constants：ture, false, iota, nil
 * Types: int, int8, int16, int62, uint, unit8, uint16, uint32, uint64, uintptr, 
	float32, float64, complex128, complex64, bool, byte, rune, string, error
 * Functions: make, len, cap, new, append, copy, close, delete, complex, real, imag, panic, recover

**特别注意**

第一个字母大小写决定变量或者方法的包可见性，如果是大写，包外部可以访问，若是小写，只能包内部访问。

例如在自定义包中访问fmt的Println方法，但是不能够访问fmtInteger等方法.

## declarations

go中主要有四种声明类型：

* var
* const
* type
* func

## variables

语法如下：

```
	var name type = expression
```

其中type，expression可以二缺一，但是不能都缺省，go有类型推断机制。

go语言采用缺省值机制： **zero-value machanism** 

* 数值缺省值为0
* 布尔值缺省值为false
* 字符串缺省值为""
* interfaces & reference type(slice,pointer,map,channel,function)缺省值为nil

**同一类型变量声明**

`var i,j,k int`

**不同类型变量声明**

`var a, b, c = true, 1.1, "abc"`

**根据方法调用初始化变量**

`var f, err = os.Open("c:\windows\xx.dll")`

上面在调动文件打开方法时，返回一个文件对象赋值给f，返回一个异常对象赋值给err

### short variable declarations

短变量声明的语法为

`name := expression`

例如：

`t := 0.123`

短变量声明通常在声明局部变量时使用，var变量通常用于需要显示指定变量类型或者变量初始值不重要，后续赋值时使用

### Pointers

Go语言中也支持指针，&为变量的地址， *为变量的值

`x := 1 p:=`