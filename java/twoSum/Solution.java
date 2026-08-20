package twoSum;

import java.util.HashMap;

public class Solution {
    public int[] twoSum(int[] nums, int target) {

        int[] indicies = new int[2];
        
        HashMap<Integer, Integer> ints = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            if(ints.containsKey(Integer.valueOf(target - nums[i]))) {
                indicies[0] = ints.get(Integer.valueOf(target - nums[i]));
                indicies[1] = i;
                return indicies;
            }
            ints.put(Integer.valueOf(nums[i]), i);
        }

        return indicies;
    }
}