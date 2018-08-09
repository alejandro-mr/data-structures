#!/usr/bin/env python
from LinkedList import LinkedList
from Node import Node

def main():
    ll = LinkedList()
    # Insertion methods
    ll.insert(Node(1))
    ll.insert(Node(2))
    ll.insert(Node(3))
    ll.insert_after(Node(4), Node(3))
    # Print current llst
    ll.print_list()

    print("Inserting element at beggining of llst")
    ll.insert_beginning(Node(5))
    # Print current llst
    ll.print_list()
    print("Count: ", ll.count, "\n")

    # Removal methods
    print("Removing node with value 2")
    ll.remove(Node(2))
    ll.print_list()
    print("Count: ", ll.count, "\n")

    print("Removing last element of llst")
    ll.pop()
    ll.print_list()
    print("Count: ", ll.count, "\n")

    # trying to remove non existing node
    print("Removing not existent element Node(2)")
    ll.remove(Node(2))
    ll.print_list()
    print("Count: ", ll.count, "\n")

    # empty whole llst with pop calls
    print("About to empty list with `LinkedList.pop` calls")
    for _ in range(ll.count + 10):
        ll.pop()
    ll.print_list()
    #ll.empty()
    print("Count: ", ll.count, "\n")

if __name__ == '__main__' :
    main()
