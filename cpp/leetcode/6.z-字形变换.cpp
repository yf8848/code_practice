/*
 * @lc app=leetcode.cn id=6 lang=cpp
 *
 * [6] Z 字形变换
 */
#include <string>
#include <vector>
#include<algorithm>

using namespace std;

// @lc code=start
class Solution {
public:
    string convert(string s, int numRows) {
        if (numRows == 1){
            return s;
        }

        vector<string> rows(min(numRows, int(s.size())));
        bool goingDown = false;
        int curRow = 0;
        for(char c: s){
            rows[curRow]+=c;
            if (curRow == numRows-1|| curRow == 0){
                goingDown = !goingDown;
            }
            curRow += goingDown?1:-1;
        }

        string ret;
        for(string row: rows){
            ret+=row;
        }
        return ret;
    }
};
// @lc code=end

