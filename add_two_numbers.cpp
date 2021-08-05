#include "listf.h"

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