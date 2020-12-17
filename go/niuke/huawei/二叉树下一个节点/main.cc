#include <iostream>

using namespace std;

struct TreeLinkNode {
    int val;
    struct TreeLinkNode *left;
    struct TreeLinkNode *right;
    struct TreeLinkNode *next;
    TreeLinkNode(int x) :val(x), left(NULL), right(NULL), next(NULL) {
        
    }
};
class Solution {
public:
    TreeLinkNode* GetNext(TreeLinkNode* pNode)
    {
        if(!pNode) return pNode;
        
        if(pNode->right){
            TreeLinkNode* temp = pNode->right;
            while(temp->left){
                temp = temp->left;
            }
            return temp;
        }
        
        if(pNode->next){
            TreeLinkNode* temp = pNode->next;
            while(temp){
                if(temp->left == pNode) return temp;
                pNode = pNode->next;
            }
        }
        
        return NULL;
    }
};

int main(){
    Solution s;
    cout<<"outPut" <<  s.GetNext(NULL) <<endl;
}