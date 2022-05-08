# 编译静态库
# gcc -c ./sm3.c ./sm3.h

# ar rcs  libsm3.a sm3.o

mkdir build
cd build
cmake -G "MinGW Makefiles" ../
cmake .. -DCMAKE_BUILD_TYPE=Release
cmake --build ./  --target sm3 --config Release