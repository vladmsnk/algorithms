#include "../includes/listf.h"

Node *deleteDuplicates(Node *head)
{
    Node *prev;
    Node *list = head;
    while (head != NULL)
    {
        if (prev->val == head->val)
        {
            Node *next = head->next;
            prev->next = next;
            delete head;
            head = next;
            
        }
        else
        {
            prev = head;
            head = head->next;
        }
    }
    return (list);
} 