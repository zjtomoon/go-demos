cmake_minimum_required(VERSION 3.13)
project(sm3 C)

set(CMAKE_C_STANDARD 11)

include_directories(.)

# 生成静态库
#add_library(sm3 sm3.c sm3.h)
# 生成动态库
add_library(sm3 SHARED sm3.c sm3.h)

# 编译生成release
# mkdir build
# cd build
# cmake -G "MinGW Makefiles" ../
# cmake .. -DCMAKE_BUILD_TYPE=Release
# cmake --build ./  --target sm3 --config Release
