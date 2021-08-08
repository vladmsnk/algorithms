#ifndef LISTF_H
# define LISTF_H

# include <iostream>

using namespace std;

struct Node
{
    int val;
    Node *next;
    Node() : val(0), next(nullptr) {}
    Node(int x) : val(x), next(nullptr) {}
    Node(int x, Node *next) : val(x), next(next) {}
};

Node   *init(int val);
void    pushBack(Node **head, int val);
void    printList(Node *head);



#endif