IN: 1 -> 2 -> 2 -> 2 -> 3 -> 4 -> 4  
OUT: 1 -> 2 -> 3 -> 4

**Notice**:  
Without buffer

###Pseudocode:
```
l = List
currNode = List.Head
prevNode = nil

FOR currNode != nil    
    if prevNode != nil && currNode == prevNode {
        prevNode.next = currNode.next
        currNode = currNode.next
        l.len--
        continue
    }
    
    prevNode = currNode
    currNode = currNode.next
ENDFOR 
``` 