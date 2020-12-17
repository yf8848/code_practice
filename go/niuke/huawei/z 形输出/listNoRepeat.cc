
#include <iostream>

using namespace std;

struct ListNode {
    int val;
    struct ListNode *next;
    ListNode(int x) :
        val(x), next(NULL) {
    }
};

class Solution {
public:
    ListNode* deleteDuplication(ListNode* pHead)
    {
        if(!pHead||!pHead->next) return pHead;
        
        ListNode* r = pHead;
        ListNode* behind = pHead;
        ListNode* forward = pHead->next;
        bool equaled = false;
        bool isFirst = false;
        
        while(forward!=NULL){
            if(forward->val == behind->val){
                equaled = true;
                forward = forward->next;
            }else{
                if(equaled &&pHead->val == r->val && pHead->val==behind->val){
                    //cout<<pHead->val<<"/"<<r->val<<"/"<<behind->val<<endl;
                    isFirst = true;
                }
                if(!equaled) r = behind;
                behind->next=forward;
                behind = behind->next;
                
                if(equaled){
                    r->next=forward;
                    //cout<<r->val<<"/"<<behind->val<<"/"<<forward->val<<endl;
                }
                forward = forward->next;
                equaled = false;
            }
        }
        //r->next=forward;
        
        if(isFirst) return pHead->next;
        return pHead;
    }
};

ListNode* createList(){
    ListNode* pHead = new ListNode(1);

    pHead->next = new ListNode(2);
    pHead->next->next = new ListNode(3);
    pHead->next->next->next = new ListNode(3);
    pHead->next->next->next->next = new ListNode(4);
    pHead->next->next->next->next->next = new ListNode(4);
    pHead->next->next->next->next->next->next = new ListNode(5);

    return pHead;
}

void PrintList(ListNode* head){
    while(head!=NULL){
        cout<< head->val <<"-";
        head=head->next;
    }
    cout<<endl;
}

int main(){
    Solution s;
    ListNode* l = createList();
    PrintList(l);
    PrintList(s.deleteDuplication(l));
}