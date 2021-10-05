#ifndef WRAPPER_H
#define WRAPPER_H

typedef struct Fix {
    int *ADEP_1;
    int *ADEP_2;
    int *ADEP_3;
    int *ADEP_4;
    int *ADEP_5;
} Data;

void wrapper(Data *in);

#endif