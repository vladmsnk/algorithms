#ifndef LISTF_H
# define LISTF_H

# include <iostream>

using namespace std;

class Node
{
    public:
        int val;
        Node *next;
};

Node   *init(int val);
void    pushBack(Node **head, int val);
void    printList(Node *head);


#endif