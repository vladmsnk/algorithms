#include "../includes/listf.h"


Node *deleteDuplicates(Node *list)
{
    Node *tmp;
    Node *prev = new Node(0, list);
    tmp = list;
    Node *first = prev;
    while (tmp != NULL)
    {
        Node *for_check = tmp;
        int flag = 0;
        while ((for_check->next != NULL) && (for_check->val == for_check->next->val))
        {
            for_check = for_check->next;
            flag = 1;
        }
        if (flag == 0)
        {
            prev = tmp;
            tmp = tmp->next;
        }
        else
        {
            prev->next = for_check->next;;
            flag = 0;
            tmp = for_check->next;
        }
    }
    return (first->next);
}