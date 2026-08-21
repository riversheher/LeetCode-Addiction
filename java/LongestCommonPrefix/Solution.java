package LongestCommonPrefix;

public class Solution {

    public String longestCommonPrefix(String[] strs) {
        if(strs.length == 1){
            return strs[0];
        }

        String prefix = "";

        if(strs[0].length() == 0) {
            return prefix;
        }

        for(int i = 0; i < strs[0].length(); i++){
            char candidate = strs[0].charAt(i);

            for(String str : strs) {
                if (i >= str.length() || str.charAt(i) != candidate){
                    return prefix;
                }
            }

            prefix += candidate;
        }

        return prefix;
    }
    
}
