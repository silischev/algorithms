IN: (7->1->6) + (5->9->2) = 617 + 295 = 912
OUT: 2->1->9

###Pseudocode:
```
l1 = List
l2 = List
l3 = List{nil, 0}

currNodeLarger = l1.head
currNodeSmaller = l2.head
IF l2.len > l1.len
    currNodeLarger = l2.head
    currNodeSmaller = l1.head
ENDIF

currRes = 0

FOR currNodeLarger != nil
    IF currNode2 != nil
        currRes += currNodeLarger.data + currNode2.data
        currNode2 = currNode2.next
    ELSE
        currRes += currNodeLarger.data
    ENDIF
    
    currNodeLarger = currNodeLarger.next
    
    IF currRes > 9
        l3.AddFront(currRes MOD 10)
        currRes = 1    
    ELSE
        l3.AddFront(currRes)
        currRes = 0 
    ENDIF
    
    l3.AddFront(currRes)
ENDFOR
``` 