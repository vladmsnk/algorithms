#include "listf.h"

Node *reverse_list(Node *head)
{
    Node    *previous;
    Node    *next_elem;

    previous = NULL;
    while (head != NULL)
    {
        next_elem = head->next;
        head->next = previous;
        previous = head;
        head = next_elem;
    }
    return (previous);
}