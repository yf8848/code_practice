#include <vector>
#include <queue>
#include <iostream>

using namespace std;

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
    TreeNode(int x) :
            val(x), left(NULL), right(NULL) {
    }
};

class Solution {
public:
    vector<vector<int> > Print(TreeNode* pRoot) {
        vector<vector<int>> res;
        bool reverse = false;
        queue<vector<TreeNode*>> queue;
        queue.push({pRoot});
        
        while(!queue.empty()){
            vector<TreeNode*> row;
            vector<int> raw;
            vector<TreeNode*> line = queue.front();
            queue.pop();
            
            //cout<<queue.size()<<"-"<< line.size()<<"-"<< res.size()<<endl;
            for(int i=0; i<line.size();++i){
                TreeNode* node = line[i];
                if (!node) continue;

                //cout<<"node: "<<node->val <<" left:"<<node->left<<" right:"<<node->right<<endl;
                raw.push_back(node->val);

                if(node->left) row.push_back(node->left);
                if(node->right) row.push_back(node->right);
            }
            //cout<<raw.size()<< "-" << row.size()<<endl;
            if(raw.size()>0){
                if(reverse){
                    int i = 0;
                    int j = raw.size()-1;
                    while(i<j){
                        int t = raw[i];
                        raw[i] = raw[j];
                        raw[j] = t;
                        ++i;
                        --j;
                    }
                }
                res.push_back(raw);
            }
            if(row.size()>0){
                queue.push(row);
            }
            reverse = !reverse;
        }
        return res;
    }
};

TreeNode* createTree()  {
	TreeNode* root = new TreeNode(1);

	root->left = new TreeNode(2);
	root->right = new TreeNode(3);
	return root;
}

int main() {
    Solution s;
    vector<vector<int>> res = s.Print(createTree());
	for(int i =0; i<res.size();++i){
        
        for (size_t j = 0; j < res[i].size(); j++)
        {
            cout<<res[i][j]<<" ";
        }
        cout<<endl;  
    }
}
