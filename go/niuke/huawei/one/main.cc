#include <iostream>
#include <vector>


using namespace std;

int main(){
    string str;
    while(cin>>str){
        size_t pos = str.size()-4;
        cout<<str.size() << " " << pos<<endl;
        string pre = str.substr(0, pos);
        string last = str.substr(pos,3);

        cout<< pre<< endl;
        cout << last<<endl;
    }

    return 0;
}