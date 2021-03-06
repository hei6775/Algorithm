# KMP算法

&emsp;&emsp;常用于字符串匹配

`Tstr`目标字符串，`Pstr`模式字符串

## 核心就是最长前后缀长度

&emsp;&emsp;如果字符串A和B，存在A=BS，其中S是任意的非空字符串，那就称B为A的前缀。例如，`Harry`的前缀包括
{”H”, ”Ha”, ”Har”, ”Harr”}，我们把所有前缀组成的集合，称为字符串的前缀集合。同样可以定义后缀A=SB， 其中
S是任意的非空字符串，那就称B为A的后缀，例如，`Potter`的后缀包括{”otter”, ”tter”, ”ter”, ”er”, ”r”}，
然后把所有后缀组成的集合，称为字符串的后缀集合。要注意的是，字符串本身并不是自己的后缀。

&emsp;&emsp;有了这个定义，就可以说明`PMT`中的值的意义了。PMT中的值是字符串的前缀集合与后缀集合的交集中最长元
素的长度。例如，对于`aba`，它的前缀集合为{”a”, ”ab”}，后缀 集合为{”ba”, ”a”}。两个集合的交集为{”a”}，那
么长度最长的元素就是字符串”a”了，长 度为1，所以对于”aba”而言，它在PMT表中对应的值就是1。再比如，对于字符串
”ababa”，它的前缀集合为{”a”, ”ab”, ”aba”, ”abab”}，它的后缀集合为{”baba”, ”aba”, ”ba”, ”a”}， 两个
集合的交集为{”a”, ”aba”}，其中最长的元素为”aba”，长度为3。

主要：

- 一个关键点就是：PMT数组

- 另一个关键点就是：当i和j不匹配的时候，将模式字符串向后移动`next[j-1]`位


## Links

[很详尽KMP算法（厉害）](https://www.cnblogs.com/ZuoAndFutureGirl/p/9028287.html)

[ 转自知乎作者海纳 ](https://www.zhihu.com/question/21923021/answer/281346746)
