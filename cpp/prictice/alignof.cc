#include <iostream>


struct Storage
{
    char a;
    int b;
    long long c;
    double d;
};


struct alignas(std::max_align_t) AlignStorge
{
    char a;
    int b;
    long long c;
    double d;
};


int main(){
    std::cout << alignof(Storage) << std::endl;
    std::cout << alignof(AlignStorge)<<std::endl;

    return 0;
}



