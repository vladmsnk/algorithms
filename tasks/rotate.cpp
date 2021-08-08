#include "../includes/listf.h"


Node *rotateRight(Node *head, int k)
{
    if (head == NULL)
        return (NULL);
    Node    *slow = head;
    Node    *fast = head;
    Node    *next;
    int i = 0;
    while (fast != NULL)
    {
        fast = fast->next;
        i++;
    }
    fast = head;
    if ((k == 0) || (k >= i && k % i == 0))
        return (head);
    else if (k > i)
        k = k % i;
    for (int i = 0; i < k; i++)
        fast = fast->next;
    Node *first = head;
    Node *prev;
    while (fast != NULL)
    {
        if (fast->next == NULL)
            prev = slow;
        fast = fast->next;
        slow = slow->next;

    }
    if (prev)
        prev->next = NULL;
    Node *list = slow;
    while (slow->next != NULL)
        slow = slow->next;
    slow->next= first;
    return (list);
}