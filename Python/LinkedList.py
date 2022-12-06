class Node:
    def __init__(self, data = None, next = None) -> None:
        self.data = data
        self.next = next

class LinkedList:
    def __init__(self) -> None:
        self.head = None
    
    def insertAtFront(self, data):
        node = Node(data,self.head)
        self.head = node
    
    def getLen(self):
        count = 0
        itr = self.head
        while itr:
            count += 1
            itr= itr.next
        return count
       
    def print(self):
        if self.head == None:
            print(" Empty LinkedList")
            return
        
        llstr = ""
        itr = self.head
        while itr:
            llstr += str(itr.data)+"-->"
            itr = itr.next
        print(llstr)

    def insertAtEnd(self, data):
        if self.head == None:
            node = Node(data,None)
            self.head = node
            return
        
        itr = self.head
        while itr.next:
            itr = itr.next
        itr.next = Node(data,None)

    def insertValues(self, dataList):
        self.head = None
        for data in dataList:
            self.insertAtEnd(data)
    
    def removeAt(self, index):
        if index < 0 or index >= self.getLen():
            raise Exception("Invalid Index")
        
        if index == 0:
            self.head = self.head.next
            return
        
        count = 1

        itr = self.head
        
        while itr:
            if count == index:
                itr.next = itr.next.next
                break
            count += 1
            itr = itr.next

    def insertAt(self,index,data):
        if index < 0 or index >= self.getLen():
            raise Exception("Invalid index")
        
        if index == 0:
            self.insertAtFront(data)
            return
        
        count = 1
        itr = self.head
        while itr:
            if count == index:
                node = Node(data,itr.next)
                itr.next = node
                break
            itr = itr.next
            count += 1
        
    def removeByVal(self,data):
        itr = self.head

        if itr.data == data:
            self.head = itr.next
            return

        while itr.next:
            if itr.next.data == data:
                itr = itr.next.next
                return
            itr = itr.next.next
        print(" Value not found")

            
            

if __name__ == "__main__":
    ll = LinkedList()
    ll.insertValues(["a","b","c","d"])
    ll.print()
    print(ll.getLen())
    ll.removeAt(1)
    ll.print()
    ll.insertAt(2,"e")
    ll.print()
    ll.removeByVal("e")
    ll.print()

    