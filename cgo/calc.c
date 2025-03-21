#include "calc.h"
#include <limits.h>

/**
 * @brief 计算两个整数的和。
 * 
 * @param a 第一个整数。
 * @param b 第二个整数。
 * @return 返回两个整数的和。
 */
int add(int a, int b) {
    return a + b;
}

/**
 * @brief 计算两个整数的差。
 * 
 * @param a 被减数。
 * @param b 减数。
 * @return 返回两个整数的差。
 */
int subtract(int a, int b) {
    return a - b;
}

/**
 * @brief 计算两个整数的乘积。
 * 
 * @param a 第一个整数。
 * @param b 第二个整数。
 * @return 返回两个整数的乘积。
 */
int multiply(int a, int b) {
    return a * b;
}

/**
 * @brief 计算两个整数的商。
 * 
 * @param a 被除数。
 * @param b 除数。
 * @return 如果除数为0返回0，否则返回两个整数的商。
 */
int divide(int a, int b) {
    if (b == 0) {
        return 0;
    }
    return a / b;
}

/**
 * @brief 带溢出检查的加法。
 * 
 * @param a 第一个整数。
 * @param b 第二个整数。
 * @param has_overflow 溢出标志，如果发生溢出则设为1，否则设为0。
 * @return 返回两个整数的和，如果发生溢出，返回0。
 */
int add_with_overflow_check(int a, int b, int* has_overflow) {
    long long result = (long long)a + (long long)b;
    if (result > INT_MAX || result < INT_MIN) {
        *has_overflow = 1;
        return 0;
    }
    *has_overflow = 0;
    return (int)result;
}

/**
 * @brief 计算两个长整型的和。
 * 
 * @param a 第一个长整型。
 * @param b 第二个长整型。
 * @return 返回两个长整型的和。
 */
long long add_long(long long a, long long b) {
    return a + b;
}

/**
 * @brief 计算整数的绝对值。
 * 
 * @param a 输入整数。
 * @return 返回输入整数的绝对值。
 */
int abs_value(int a) {
    return a < 0 ? -a : a;
}

/**
 * @brief 返回两个整数中的较大值。
 * 
 * @param a 第一个整数。
 * @param b 第二个整数。
 * @return 返回两个整数中的较大值。
 */
int max_value(int a, int b) {
    return a > b ? a : b;
}

/**
 * @brief 返回两个整数中的较小值。
 * 
 * @param a 第一个整数。
 * @param b 第二个整数。
 * @return 返回两个整数中的较小值。
 */
int min_value(int a, int b) {
    return a < b ? a : b;
}
