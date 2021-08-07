
#include "../includes/listf.h"
#include <iostream>
#include <vector>

using namespace std;


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

Node *mergeKList(vector<Node *>& lists)
{
    Node *tmp;
    int     i;
    tmp = lists.at(0);
    for (i = 1; i < lists.size(); i++)
    {
        Node    *head;
        Node    *first;
        Node    *list1 = tmp;
        Node    *list2 = lists.at(i);

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

           if (list1 == NULL && list2 != NULL)
            {
                head->next = list2;
                head = head->next;
                list2 = list2->next;
            }
            else if (list2 == NULL && list1 != NULL)
            {
                head->next = list1;
                head = head->next;
                list1 = list1->next;
            }
            else if (list1->val >= list2->val)
            {
                head->next = list2;
                head = head->next;
                list2 = list2->next;
            }
            else if (list2->val > list1->val)
            {
                head->next = list1;
                head = head->next;
                list1 = list1->next;
            }
        }
        tmp = first->next;
    }
    return (tmp);
}