# Ran10ByRand7Go

[leetcode 470](https://leetcode.com/problems/implement-rand10-using-rand7/)

implement Ran10 given that Ran7 is already implemented as a function generates uniform random integer in range [1, 7], write a function rand10() that generates a uniform random integer in the range [1, 10]


## 我的解法

首先是因為是 uniform random 
 
Rand7  [1, 7] 

把一個 function 對應到 7 個數字

Rand10 [1, 10]

把一個 function 對應到 10 個數字

假設做了兩次 Rand7 

兩次各得 一個隨機整數 a, b

透過以下 算式  b + (a - 1) * 7 可以把 f(a,b) 對應到 [1, 49]

其中 [1, 40] 剛好可以是 10的倍數 % 運算後 再加上 1 可以對應到 [1, 10] 

而 [41, 49] 無法做 uniform map 則可以再重算一次

## 參考Kata

[Recca Chao's leetcode 470](https://gitpage.reccachao.net/kotlin/leetcode/470/)