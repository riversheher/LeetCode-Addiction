package majorityElement;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class Solution {
    public int majorityElement(int[] nums) {

        HashMap<Integer, Integer> counts = new HashMap<>();

        for (int num : nums) {
            counts.put(num, counts.getOrDefault(num, 0) + 1);
            if (counts.get(Integer.valueOf(num)) > nums.length/2) {
                return num;
            }
        } 
        
        return -1;
    }

    /*
    Takes advantage of the fact that there can only be one majority element
    Therefore, the majority element will be the one with a "leftover" copy
    after discarding down.    
    */ 
    public int majorityElementBoyerMoore(int[] nums) {
        int candidate = nums[0];
        int count = 0;
        for (int num : nums) {
            if (count == 0) {
                candidate = num;
            }
            
            if (candidate == num) {
                count++;
            } else {
                count--;
            }
        }

        return candidate;
    }

    public List<Integer> majorityElement2(int[] nums) {
        Set<Integer> solutionSet = new HashSet<>();
        HashMap<Integer, Integer> counts = new HashMap<>();

        for (int n : nums) {
            counts.put(n, counts.getOrDefault(n, 0) + 1);

            if (counts.get(Integer.valueOf(n)) > nums.length / 3) {
                solutionSet.add(n);
            }
        }

        return new ArrayList<>(solutionSet);
    }

    /*
    Extension of original BoyerMoore,
    Takes advanatage of the fact that there can only be 2 elements
    that take up more than a third of the length.

    Same principle applies, since they number greater than the remaining set,
    there will always be leftovers for the ones who have greater than a third.    
    */
    public List<Integer> majorityElement2BoyerMoore(int[] nums) {
        int candidateOne = 0, candidateTwo = 0;
        int countOne = 0, countTwo = 0;
        List<Integer> solutions = new ArrayList<>();

        for(int n : nums) {
            if (n == candidateOne) {
                countOne++;
            } else if (n == candidateTwo){
                countTwo++;
            } else if (countOne == 0) {
                candidateOne = n;
                countOne++;
            } else if (countTwo == 0){
                candidateTwo = n;
                countTwo++;
            } else {
                countOne--;
                countTwo--;
            }
        }

        countOne = 0;
        countTwo = 0;

        for(int n : nums) {
            if(candidateOne == n) {
                countOne++;
            } else if (candidateTwo == n) {
                countTwo++;
            }
        }


        if(countOne > nums.length / 3) {
            solutions.add(candidateOne);
        }
        if(countTwo > nums.length / 3) {
            solutions.add(candidateTwo);
        }

        return solutions;
    }
}
