#include "../includes/listf.h"
int count_list_len(Node *list)
{
    int i;

    i = 0;
    while (list != NULL)
    {
        i++;
        list = list->next;
    }
    return (i);
}


Node    *removeNthFromEnd(Node *list, int n)
{
    int     j;
    Node    *head;
    int     len;

    j = 0;
    head = list;
    len = count_list_len(list) - n;
    if (len == 0)
    {
        list = list->next;
        return (list);
    }
    while (j < len)
    {
        if (len - j == 1)
        {
            list->next = list->next->next;
            list = list->next;
            return (head);
        }
        list = list->next;
        j++;
    }
    return (head);
}

