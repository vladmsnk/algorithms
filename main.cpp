#include "listf.h"

Node    *addTwoNumbers(Node *l1, Node *l2);
Node *reverse_list(Node *head);


int main(void)
{
    Node    *list1;
    Node    *list2;
    Node    *output;

    list1 = NULL;
    list2 = NULL;
    pushBack(&list1, 2);
    pushBack(&list1, 4);
    pushBack(&list1, 3);   

    pushBack(&list2, 5);
    pushBack(&list2, 6);
    pushBack(&list2, 4);

    // output = addTwoNumbers(list1, list2);
    output = reverse_list(list1);
    printList(output);
   return (0);
}