# [Majority Element II](https://leetcode.com/problems/majority-element-ii/description/)

## Constraints

- T.c. O(N)
- S.C. O(1)

## Brute Force Approach

- Making a dictionary to store the occurence of every element

### Time Complexity

- O(N)

### Space Complexity

- O(N)

## Optimized Approach

- Using Boyer-Moore Voting Algorithm.
- To find N/K elements, we need to first find the K-1 majority elements.
    - We start with a k-1 counts set to 0 & index of k-1 max elements i.e. count_1:=1 & max_element_1:=0 
    - We loop from index=0 till length of the array.
    - For each K-1 max elements having k-1 count:
        - If the next element is the same element, we increase the count.
        - else If the next element is a different element , we decrease the count.
        - else, if we check if our count has reached 0. Once count has reached 0, we make the current element as our max element & reset count to 1.
        - else, we decrease the count of count variables
    - Now, we loop through the array and maintain a frequency_count for each max elements.
    - If the max elements, are present more than N/K times and we check if the indexs are different for all max elements, we append it to the answer array.
    - We return the answer array.

### Pseudocode

```golang
    //for N/3 times
    func majorityElement(nums []int) []int {
	ans := []int{}
	
    major1:=0
    major2:=0
    count1,count2:=0,0
    for i:=0;i<len(nums);i++{
        if(nums[i]==nums[major1]){
            count1+=1
        }else if nums[i]==nums[major2]{
            count2+=1
        }else  if count1==0{
            major1=i
            count1=1
        }else if count2==0{
            major2=i
            count2=1
        }else{
            count1-=1
            count2-=1
        }
    }
    freq1,freq2:=0,0
    for i:=0;i<len(nums);i++{
        if nums[i]==nums[major1]{
            freq1+=1
        }else if nums[i]==nums[major2]{
            freq2+=1
        }
        
    }
    if(freq1>len(nums)/3){
            ans=append(ans,nums[major1])
    }
    if(freq2>len(nums)/3 && major1!=major2){
        ans=append(ans,nums[major2])
    }
    return ans

}
```