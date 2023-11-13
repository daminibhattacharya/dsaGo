package main
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