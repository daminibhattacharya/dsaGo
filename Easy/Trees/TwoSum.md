# [Two Sum](https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/)

## Constraints

- The number of nodes in the tree is in the range [1, 104].
- 104 <= Node.val <= 104
- root is guaranteed to be a valid binary search tree.
- 105 <= k <= 105

## Brute Force Approach

Loop through the tree and try to find the value k-currentNode.val for each currentNode

### Time Complexity

- O(N*N)

### Space Complexity

- O(1)

## Optimized Approach

-Store the values of node visit in a dictionary. For every node vistied, try to find if k-currentNode.val already exists in dictionary, if so return true else store currentNode.val in the dictionary as true. 

-We call the same function to find the value in currentNode.Left or currentNode.Right, if anyone of these two return true, return true. 


### Explanation (Pseudocode)

```plaintext
func twoSum( root *treeNode, k int)bool{
    if root==nil{
        return false
    }
    dict:=make(map[int]bool)
    return getTwoSum(root,k,dict)
}
func getTwoSum(root *treeNode, k int, dict make(map[int]bool)) bool{
    if root==nil{
        return false
    }
    if(dict[k-root.val]){
        return true 
    }// already have save the k-root.val in previous steps
    dict[root.val]=true
    return getTwoSum(root.left,k,dict) || getTwoSum(root.right,k,dict)
}
```
