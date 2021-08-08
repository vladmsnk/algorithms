#include "../includes/listf.h"

Node    *mergeTwoLists(Node *list1, Node *list2)
{
    Node    *head;
    Node    *first;
    if (list1 == NULL)
        return (list2);
    else if (list2 == NULL)
        return (list1);
    head = new Node();
    head->val = 0;
    head->next = NULL;
    first = head;
    while (list1 != NULL || list2 != NULL)
    {
        if (list1 == NULL || list1->val >= list2->val)
        {
            head->next = list2;
            list2 = list2->next;
        }
        else if (list2 == NULL || list2->val > list1->val)
        {
            head->next = list1;
            list1 = list1->next;
        }
        head = head->next;
    }
    return (first->next);
}