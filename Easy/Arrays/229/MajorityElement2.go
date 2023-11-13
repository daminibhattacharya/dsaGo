package main

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