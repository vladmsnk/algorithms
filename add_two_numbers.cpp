#include <iostream>

using namespace std;

class Node 
{
    public:
        int val;
        Node *next;
};

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

int make_number(Node *list)
{
    int number;

    number = 0;
    while (list != NULL)
    {
        number = number * 10 + list->val;
        list = list->next;
    }
    return (number);
}

int reverse_number(int number)
{
    int new_number;

    new_number = 0;
    while (number != 0)
    {
        new_number = new_number * 10 + number % 10;
        number /= 10;
    }
    return (new_number);
}


Node    *addTwoNumbers(Node *l1, Node *l2)
{
    Node    *final_list;
    int     output;

    output = reverse_number(make_number(l1)) +  reverse_number(make_number(l2));
    final_list = new Node();
    final_list = NULL;
    if (output == 0)
    {
        pushBack(&final_list, 0);
        return (final_list);
    }
    while (output != 0)
    {
        pushBack(&final_list, output  % 10);
        output /= 10;
    }
    return (final_list);
}

int main(void)
{
    Node    *list1;
    Node    *list2;
    Node    *output;

    list1 = NULL;
    list2 = NULL;
    pushBack(&list1, 2);
    pushBack(&list1, 4);
    pushBack(&list1, 3);   

    pushBack(&list2, 5);
    pushBack(&list2, 6);
    pushBack(&list2, 4);

    output = addTwoNumbers(list1, list2);
    printList(output);
   return (0);
}