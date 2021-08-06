#include "../includes/listf.h"

int count_list_len(Node *list)
{
    int len;

    len = 0;
    while (list != NULL)
    {
        len++;
        list = list->next;
    }
    return (len);
}

Node *reverse_list(Node *list)
{
    Node *previous;
    Node *next_elem;

    previous = NULL;
    while (list != NULL)
    {
        next_elem = list->next;
        list->next = previous;
        previous = list;
        list = next_elem;
    }
    return (previous);
}

void    reorderList(Node *List)
{
    int i;
    Node *cur;
    Node *reversed;
    Node *next;
    Node *tmp;
    Node *prev;

    i = 0;
    reversed = reverse_list(List);
    cur = List;
    while (i < count_list_len(List) / 2)
    {
        if (i % 2 == 1)
        {
            prev = List;
            next = List->next;
            tmp = reversed;
            prev->next = tmp;
            prev->next->next = next;
            reversed = reversed->next;
        }
        List = List->next;
        i++;
    }
    return (cur)
}