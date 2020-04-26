#include <string>
#include <iostream>
#include <map>
#include <vector>
#include <set>
#include<algorithm>

using namespace std;

class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int i=0,j=0,ans=0;
        int n = s.length();
        set<char> ss;
        while(i<n&&j<n){
            char c = s.at(j);
            if(ss.find(c) == ss.end()){
                ss.insert(c);
                j++;
                ans = max(ans, j-i);
            }else{
                ss.erase(s.at(i++));
            }
        }
        return ans;
    }
};

Solution soul;

void printFun(string str){
    cout<<"test:" << str<<endl;
    cout<<"result: "<< soul.lengthOfLongestSubstring(str)<<endl;
}

int main(){
    vector<string> test_case;
    test_case.push_back("abcabcbb");
    test_case.push_back("bbbbb");
    test_case.push_back("pwwkew");

    for_each(test_case.begin(), test_case.end(), printFun);

    return 0;
}