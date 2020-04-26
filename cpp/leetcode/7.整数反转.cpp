/*
 * @lc app=leetcode.cn id=7 lang=cpp
 *
 * [7] 整数反转
 */


#include<string>
#include<iostream>
#include<vector>

using namespace std;
// @lc code=start
class Solution {
public:
    int reverse(int x) {
        int old = x;
        int ret =0;
        while (old!=0)
        {
            int num =abs(old%10);
            if (ret > (maxNum-num)/10){
                // cout<<"1 bigger: "<< ret<<" - "<<num<<endl;
                return 0;
            }
            ret = ret*10+num;
            old = old/10;
        }

        if (ret >= maxNum && x>0){
            // cout<<"2 bigger"<< ret<<endl;
            return 0;
        }
        if (ret>maxNum&&x<0){
            // cout<<"3 bigger"<< ret<<endl;
            return 0;
        }
        
        if (x>=0){
            return ret;
        }else{
            return -ret;
        }
        
    }

    int maxNum = ((1<<30)-1)*2+1;

};
// @lc code=end

Solution s;
void printFun(int num){
    cout<<"test case:"<<num<<" res: "<<s.reverse(num)<<endl;
}
int main(){
    vector<int> test_case;
    test_case.push_back(123);
    test_case.push_back(-123);
    cout<<"max:"<<s.maxNum<<endl;
    
    for_each(test_case.begin(), test_case.end(), printFun);

    return 0;
}


