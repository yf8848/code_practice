#include <iostream>

using namespace std;

void hash_prictice(const string &str){
    uint32_t hash = std::hash<std::string>()(str);
    std::cout << str <<" hash_code: "<<hash<<std::endl;
}

int main(){
    hash_prictice("test_string_hash");


    return 0;
}
