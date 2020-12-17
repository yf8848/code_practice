#include <vector>
#include <queue>
#include <iostream>

using namespace std;

class Solution {
    struct Pos{
        int x;
        int y;
    };
    int arrx[4]={-1,0,1,0};
    int arry[4]={0,-1,0,1};
public:
    int movingCount(int threshold, int rows, int cols)
    {
        vector<vector<bool>> visited(rows);
        for( int i =0; i< rows; i++){
            visited[i] = vector<bool>(cols, false);
        }
        
        queue<Pos> q;
        q.push(Pos{0,0});
        int sum =0;
        while(!q.empty()){
            Pos p = q.front();
            q.pop();
            int v = val(p.x)+val(p.y);
            cout<<"point: "<< p.x << "-" <<p.y << " ->"<<v<<endl;

            if(v <= threshold && !visited[p.x][p.y]){
                visited[p.x][p.y] = true;
                sum++;
                
                for(int i =0; i<4;i++){
                    int x = p.x + arrx[i];
                    int y = p.y + arry[i];
                    if( x>=0&&x<rows && y>=0 && y<cols && !visited[x][y]){
                        q.push(Pos{x,y});
                    }
                }
            }
        }
        return sum;
    }

    int val(int v){
        int sum =0;
        while(v!=0){
            sum+=v%10;
            v /=10;
        }
        return sum;
    }
};

int main(){
    Solution s;
    int t,r,c;
    while(cin>>t>>r>>c){
        cout << "inputs:"<<t<<" - "<<r<<" - "<<c<<": "<<s.movingCount(t,r,c)<<endl;
    }
}