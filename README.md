gofractions
===========

[![Build Status](https://travis-ci.org/xyproto/gofractions.svg?branch=master)](https://travis-ci.org/xyproto/gofractions) [![GoDoc](https://godoc.org/github.com/xyproto/gofractions?status.svg)](http://godoc.org/github.com/xyproto/gofractions)

A package for dealing with fractions.

It can convert from floating point numbers or strings to fractions and back.


Output from tests
-----------------

```
11/1
11/1
-8/5
123/1
0/1
3/7
-3/7
-3/7
1/6 looks nicer than 0.16666666666666666
y is 1/1
z is 1/6 0 ( 0.16666666666666666 )
3/2 2
3/4 1
num dom i 		 fraction 	 float 		 rounded
{1 2 400} 		 1/2 		 0.5 		 1
{15708 5000 400} 	 15708/5000 	 3.1416 	 3
0.7 + 2 = 27/10 3 2.7 2.7
0.5 - 4 = -7/2 -3 -3.5 -3.5
  1/3 0.3333333333333333
+ 1/2 0.5
= 5/6 0.8333333333333334
  1/2 0.5
- 1/3 0.3333333333333333
= 1/6 0.16666666666666666
3/2 is also 1 and 1/2
PASS
ok  	gofractions	0.005s
```
