#include "../includes/listf.h"

Node   *init(int val)
{
    Node *newnode = new Node();

    newnode->next = NULL;
    newnode->val = val;
    return (newnode);
}

void    pushBack(Node **head, int val)
{
    Node *to_push;
    Node *last;

    to_push = init(val);
    if (*head == NULL)
    {
        *head = to_push;
        return;
    }
    last = *head;
    while (last->next != NULL)
        last = last->next;
    last->next = to_push;
}

void    printList(Node *head)
{
    while (head != NULL)
    {
        cout << head->val << endl;
        head = head->next;
    }
}