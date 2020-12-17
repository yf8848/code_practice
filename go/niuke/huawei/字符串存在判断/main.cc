#include <string>
#include <iostream>

using namespace std;


class Solution {
public:
    bool hasPath(char* matrix, int rows, int cols, char* str)
    {
        for(int i = 0; i< rows; i++){
            for(int j =0; j<cols; j++){
                if(matrix[getPos(cols, i, j)] == str[0]&&
                    isPath(string(matrix), i, j,rows, cols, str, 0)){
                        return true;
                }
            }
        }
        return false;
    }

    bool isPath(string matrix, int posx, int posy,int rows, int cols, string str, int ps){
        int pos0 = getPos(cols, posx, posy);
        //cout<< matrix <<": "<< matrix[pos0] << "-" << pos0 << "|" <<str[ps] <<"-"<<ps <<endl;
        if(matrix[pos0] != str[ps]){
            return false;
        }
        
        if(ps == str.length()-1){
            return true;
        }
        
        char old = matrix[pos0];
        matrix[pos0] = '-';
        int xArr[] ={-1,0,1,0};
        int yArr[] ={0, -1, 0, 1};

        for(int j=0; j<4;j++){
            int x = xArr[j] + posx;
            int y = yArr[j] + posy;
            if (x>=0 && y>=0 && x<rows && y<cols &&
                matrix[getPos(cols, x, y)]!='-' &&
                isPath(matrix, x, y, rows, cols, str, ps+1)){
                    return true;
            }
        }
        matrix[pos0] = old;
        return false;
    }
    
    int getPos(int lenx, int posx, int posy){
        return posx*lenx+posy;
    }
};


int main(){
    Solution s;
    string matrix = "ABCEHJIGSFCSLOPQADEEMNOEADIDEJFMVCEIFGGS";
    int x = 5;
    int y = 8;
    string str = "SGGFIECVAASABCEHJIGQEM";
    /**
    ABCEHJIG
    SFCSLOPQ
    ADEEMNOE
    ADIDEJFM
    VCEIFGGS
    */

    cout << s.hasPath(const_cast<char*>(matrix.c_str()), x, y,const_cast<char*>(str.c_str())) << endl;
    return 0;
}