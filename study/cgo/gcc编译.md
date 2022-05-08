# 编译生成动态库

```bash
# linux
gcc -shared -o libnumber.so number.c number.h\

# windows mingw
gcc -shared -o libnumber.dll  number.c number.h
gcc -shared -o libnumber.dll.a number.c number.h 
```

## 编译生成静态库

```bash
gcc -c ./sm3.c ./sm3.h


ar rcs  libsm3.a sm3.o
```


