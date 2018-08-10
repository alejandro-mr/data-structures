class LinkedList:
    def __init__(self, head=None):
        self.head = head
        self.size = 0;

    def empty(self):
        if not self.head:
            return true
        return false

    def size(self):
        return self.size

    def print_list(self):
        current = self.head
        if not current:
            print("List is empty")
        else:
            while current:
                print("Node: ", current.value)
                current = current.next
        print()
        return

    def insert(self, node):
        if not self.head:
            self.head = node
            self.size += 1
        else:
            current = self.head
            while current.next:
                current = current.next
            current.next = node
            self.size += 1

    def insert_beginning(self, node):
        if not self.head:
            self.head = none
            self.size += 1
        else:
            node.next = self.head
            self.head = node
            self.size += 1

    def insert_after(self, node, afterNode):
        current = self.head

        if current.value == afterNode.value:
            node.next = current.next
            current.next = node
            self.size += 1
            return
        else:
            while current and current.next:
                if current.next.value == afterNode.value:
                    node.next = current.next.next
                    current.next.next = node
                    self.size += 1
                    return
                else:
                    current = current.next

    def remove(self, node):
        #If Node that we are trying to remove it's head
        if self.head and self.head.value == node.value:
            self.head = self.head.next
            self.size -= 1
        else:
            current = self.head
            while current and current.next:
                if current.next.value == node.value:
                    current.next = current.next.next
                    self.size -= 1
                    return
                else:
                    current = current.next

    def pop(self):
        current = self.head
        if current and not current.next:
            self.head = None
            self.size = 0
            return
        else:
            while current and current.next:
                if not current.next.next:
                    current.next = None
                    self.size -= 1
                    return
                else:
                    current = current.next


