//
// Created by alex on 2022/5/7.
//

#include "student.h"
#include "interface.h"

#ifdef __cplusplus
extern "C"{
#endif

void *stuCreate()
{
    return new Student(); //构造
}

void getStuName(void *p)
{
    static_cast<Student *>(p)->Operation();
}

void initName(void *p, char* name1)
{
    static_cast<Student *>(p)->SetName(name1);
}

void getName()
{
    Student obj;
    obj.Operation();
}

#ifdef __cplusplus
}
#endif