#include "./includes/main.h"
#include "./lib1/includes/lib1.h"
#include "./lib2/includes/lib2.h"


int main()
{
    printf("Nombre: %s\r\n",NOMBRE);
    printf("Librerias Lib1:%d Lib2:%d\r\n",SumaLib1(10,10),SumaLib2(10,10));
    return 0;
}