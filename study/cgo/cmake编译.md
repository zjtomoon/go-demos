# cmake编译生成静态库

```cmake
add_library(hello hello.c hello.h)
```

# cmake编译生成动态库

```cmake
add_library(hello SHARED hello.c hello.h)
```

# cmake编译指令，指定Release或者Debug版本

+ 生成cache

```bash
cmake -G "MinGW Makefiles" ../
cmake -G "Unix Makefiles" ../
```

+ 设置Release或者Debug

```bash
cmake .. -DCMAKE_BUILD_TYPE=Release
```

+ 编译生成目标文件

```bash
 cmake --build ./  --target sm3 --config Release
```


