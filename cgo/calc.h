#ifndef CALC_H
#define CALC_H

// 基本算术运算
int add(int a, int b);
int subtract(int a, int b);
int multiply(int a, int b);
int divide(int a, int b);

// 边界值检查
int add_with_overflow_check(int a, int b, int* has_overflow);
long long add_long(long long a, long long b);

// 特殊运算
int abs_value(int a);
int max_value(int a, int b);
int min_value(int a, int b);

#endif
