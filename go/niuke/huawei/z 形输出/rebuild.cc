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
    char* Serialize(TreeNode *root) {    
        if(!root) return "#";
        
        string res = to_string(root->val);
        res.push_back(',');
        
        char* left = Serialize(root->left);
        char* right = Serialize(root->right);
        
        char* ret = new char[strlen(left)+strlen(right)+res.size()];
        strcpy(ret, res.c_str());
        strcat(ret, left);
        strcat(ret, right);
        return ret;
    }
    TreeNode* Deserialize(char *str) {
        return deser(str);
    }
    
    TreeNode* deser(char *&str){
        if(*str == '#'){
            ++str;
            return NULL;
        }
        
        cout<<*str;
        int num = 0;
        while(*str!=','){
            num = num*10+(*str-'0');
            ++str;
        }
        ++str;
        TreeNode* root = new TreeNode(num);
        root->left = deser(str);
        root->right = deser(str);
        return root;
    }

};

TreeNode* createTree()  {
	TreeNode* root = new TreeNode(1);

	root->left = new TreeNode(2);
	root->right = new TreeNode(3);
	return root;
}

int main(){
    Solution s;
    char* str = s.Serialize(createTree());
    cout<< str <<endl;
    s.Deserialize(str);
}