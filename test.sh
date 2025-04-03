#!/bin/bash

# 设置颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 打印带颜色的信息
print_info() {
    echo -e "${YELLOW}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 检查是否安装了Go
if ! command -v go &> /dev/null; then
    print_error "Go未安装，请先安装Go"
    exit 1
fi

# 运行mock包测试（包括gomonkey、gock、testify和goconvey）
run_mock_tests() {
    print_info "正在运行 mock 包测试..."
    if go test -gcflags=all=-l -v ./mock; then
        print_success "mock 包测试完成"
        return 0
    else
        print_error "mock 包测试失败"
        return 1
    fi
}

# 运行bytes包测试
run_byte_tests() {
    print_info "正在运行 byte 包测试..."
    if go test -v ./byte; then
        print_success "byte 包测试完成"
        return 0
    else
        print_error "byte 包测试失败"
        return 1
    fi
}

# 运行json包测试
run_json_tests() {
    print_info "正在运行 json 包测试..."
    if go test -v ./json; then
        print_success "json 包测试完成"
        return 0
    else
        print_error "json 包测试失败"
        return 1
    fi
}

# 运行leetcode包测试
run_leetcode_tests() {
    print_info "正在运行 leetcode 包测试..."
    if go test -v ./leetcode; then
        print_success "leetcode 包测试完成"
        return 0
    else
        print_error "leetcode 包测试失败"
        return 1
    fi
}

# 运行funk包测试
run_funk_tests() {
    print_info "正在运行 funk 包测试..."
    if go test -v ./funk; then
        print_success "funk 包测试完成"
        return 0
    else
        print_error "funk 包测试失败"
        return 1
    fi
}

# 主函数
main() {
    print_info "开始运行所有测试..."
    echo "----------------------------------------"

    # 记录失败的测试
    failed_tests=()

    # 运行mock包测试（包括所有mock相关的测试）
    if ! run_mock_tests; then
        failed_tests+=("mock")
    fi
    echo "----------------------------------------"

    # 运行bytes包测试
    if ! run_byte_tests; then
        failed_tests+=("byte")
    fi
    echo "----------------------------------------"

    # 运行json包测试
    if ! run_json_tests; then
        failed_tests+=("json")
    fi
    echo "----------------------------------------"

    # 运行leetcode包测试
    if ! run_leetcode_tests; then
        failed_tests+=("leetcode")
    fi
    echo "----------------------------------------"

    # 运行funk包测试
    if ! run_funk_tests; then
        failed_tests+=("funk")
    fi
    echo "----------------------------------------"

    # 显示测试结果摘要
    if [ ${#failed_tests[@]} -eq 0 ]; then
        print_success "所有测试通过！"
    else
        print_error "以下包的测试失败："
        for test in "${failed_tests[@]}"; do
            echo "  - $test"
        done
        exit 1
    fi
}

# 运行主函数
main
