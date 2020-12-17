
#include <string>
#include <iostream>
#include <sstream>
#include <vector>

using namespace std;

class Solution {
public:
    vector<vector<int> > FindContinuousSequence(int sum) {
        vector<vector<int>> res;
        
        // 切分因子个数
        for(int n = sum-1; n >1;n--){
            //cout<<"check n: "<<n <<endl;
            int v = sum/n;
            int delt = n/2;
            if(v - delt < 0){
                continue;
            }
            
            vector<int> val;
            int num=0, cal=0, start=0, end=0;
            if(sum%n==0){
                if(n%2==0){
                    start = v-delt;
                    end = v+delt-1;
                }else{
                    start = v-delt;
                    end = v+delt;
                }
            }else{
                if(n%2==0){
                    start = v-delt+1;
                    end = v+delt;
                }else{
                    start = v-delt;
                    end = v+delt;
                }
            }
            if (start==0){
                continue;
            }
            
            for(int j=start;j<=end;j++){
                    val.push_back(j);
                    num++;
                    cal+=j;
            }
            
            if(cal==sum){
                res.push_back(val);
            }

            //cout<<start<<"-"<<end<<":"<<cal<<endl;
        }
        return res;
    }
}; 

void fmt2vector(const vector<vector<int>>& vv){
    for(vector<vector<int>>::const_iterator it = vv.begin();it!=vv.end();it++){
        ostringstream str;
        for(vector<int>::const_iterator sub_it = it->begin(); sub_it != it->end(); ++sub_it){
            str<< *sub_it<<", ";
        }
        cout<<str.str()<<endl;
    }
}

int main(){
    Solution s;
    int num;
    while(cin>>num){
        fmt2vector(s.FindContinuousSequence(num));
    }
    return 0;
}