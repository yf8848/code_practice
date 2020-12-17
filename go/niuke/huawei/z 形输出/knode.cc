#include<iostream>

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
    TreeNode* ret = NULL;
    int count = 0;
public:
    TreeNode* KthNode(TreeNode* pRoot, int k)
    {
        if(!pRoot || k<1) return NULL;
        traval(pRoot, k);
        return ret;
    }
    
    void traval(TreeNode* root, int& k ){
        if(!root) return;
        
        if(count < k)traval(root->left, k);

        if(++count ==k){
            ret = root;
            return; 
        }
            
        if(count < k)traval(root->right, k);
    }
    
};

TreeNode* createTree()  {
    //{8,6,10,5,7,9,11},1
	TreeNode* root = new TreeNode(8);

	root->left = new TreeNode(6);
	root->right = new TreeNode(10);

    root->left->left= new TreeNode(5);
    root->left->right= new TreeNode(7);

    root->right->left= new TreeNode(9);
    root->right->right= new TreeNode(11);
	return root;
}

int main(){
    Solution s;

    cout<<s.KthNode(createTree(),3)->val << endl;
}