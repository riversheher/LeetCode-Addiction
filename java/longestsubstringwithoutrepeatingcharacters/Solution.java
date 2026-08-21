package longestsubstringwithoutrepeatingcharacters;

import java.util.ArrayDeque;
import java.util.Deque;
import java.util.HashMap;
import java.util.Map;

public class Solution {

    public int lengthOfLongestSubstring(String s) {

        int candidate = 0;
        Deque<Character> deque = new ArrayDeque<>();

        for(char curr : s.toCharArray()) {
            if(deque.contains(curr)) {
                //Update candidate if deque is larger
                if(deque.size() > candidate) {
                    candidate = deque.size();
                }
                while (deque.pollFirst() != curr){
                }
            }

            deque.addLast(curr);
        }

        if(deque.size() > candidate) {
            candidate = deque.size();
        }

        return candidate;
    }

    //sliding window solution
    public int lengthOfLongestSubstringSW(String s) {
        Map<Character, Integer> lastSeen = new HashMap<>();
        int max = 0;
        int left = 0;

        for(int right = 0; right < s.length(); right++) {
            char curr = s.charAt(right);

            if(lastSeen.containsKey(curr) && lastSeen.get(curr) >= left) {
                left = lastSeen.get(curr) + 1;
            }

            lastSeen.put(curr, right);
            max = Math.max(max, right - left + 1);
        }

        return max;
    }
}
