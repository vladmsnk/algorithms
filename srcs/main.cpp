#include "../includes/listf.h"
#include <vector>

using namespace std;

// Node    *addTwoNumbers(Node *l1, Node *l2);
// Node *reverse_list(Node *head);
// Node    *removeNthFromEnd(Node *list, int n);
// // Node    *mergeTwoLists(Node *list1, Node *list2);
// Node     *swapPairs(Node *list);
ListNode *mergeKList(vector<ListNode *>& lists);
ListNode *rotateRight(ListNode *head, int k);


int main(void)
{
    ListNode    *list1;
    ListNode    *list2;
    ListNode    *list3;
    ListNode    *list4;
    ListNode    *output;
    vector<ListNode *> arr;

    list1 = NULL;
    list2 = NULL;
    list3 = NULL;  
    list4 = NULL;
    pushBack(&list1, 1);   
    pushBack(&list1, 2);
    pushBack(&list1, 5);
    pushBack(&list1, 3);
    pushBack(&list1, 7);
    pushBack(&list1, 9);
    pushBack(&list1, 3);
    pushBack(&list1, 3);


  

    pushBack(&list2, 2);
    pushBack(&list2, 5);
    pushBack(&list2, 7);
    pushBack(&list2, 7);

    


    output = rotateRight(list1, 6);
    // output = swapPairs(list1);
    printList(output);
   return (0);
}