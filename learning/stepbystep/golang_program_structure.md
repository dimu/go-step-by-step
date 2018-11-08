
## golang names

go命名:以字母或者下划线开始，剩余部分为数字，字母或者下划线的组合。

go语言为大小写敏感，countNum与CountNum为不同命名。

### go keyword（total：25）

``` 
break default func interface select case defer go map struct chan else goto package switch
const fallthrough if range type continue for import return var  
```

其中比较少见的为：select defer go chan fallthrough range type 

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



