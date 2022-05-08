# 编译静态库
# gcc -c ./sm4.c ./sm4.h

# ar rcs  libsm4.a sm4.o

mkdir build
cd build
cmake -G "MinGW Makefiles" ../
cmake .. -DCMAKE_BUILD_TYPE=Release
cmake --build ./  --target sm4 --config Release