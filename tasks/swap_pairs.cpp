#include "../includes/listf.h"


Node *swapPairs(Node *list)
{
    Node *prev;
    Node *first;
    Node *second;
    Node *tmp;
    Node *head;
    Node *next;
    if (list == NULL || list->next == NULL)
        return (list);
    prev = new Node(0, list);
    head = prev;
    while (list != NULL && list->next != NULL)
    {
        first = list;
        second = list->next;
        next = list->next->next;
        tmp = first;
        first = second;
        prev->next = first;
        second = tmp;
        first->next = second;
        second->next = next;
        prev = second;
        list = next;
    }
    return (head->next);
}