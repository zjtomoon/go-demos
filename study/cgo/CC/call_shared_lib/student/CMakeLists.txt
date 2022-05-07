cmake_minimum_required(VERSION 3.17)
project(stu)

set(CMAKE_CXX_STANDARD 11)

#add_executable(call_shared_lib main.c student.h student.cpp interface.h interface.cpp )

add_library(stu SHARED student.h student.cpp interface.h interface.cpp )
