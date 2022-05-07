//
// Created by alex on 2022/5/7.
//

#include "interface.h"

int main() {
    void *p = stuCreate();
    char *name = "test";
    initName(p,name);
    getStuName(p);
    getName();
    return 0;
}