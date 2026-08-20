package validAnagram;

import java.util.Arrays;

public class Solution {
    public boolean isAnagram(String s, String t) {
        int[] sCharacters = new int[26];
        int[] tCharacters = new int[26];

        if(s.length() != t.length()) {
            return false;
        }

        for(int i = 0; i < s.length(); i++) {
            sCharacters[s.charAt(i) - 'a']++;
            tCharacters[t.charAt(i) - 'a']++;
        }

        if(Arrays.equals(sCharacters, tCharacters)) {
            return true;
        }
        return false;
    }

    public boolean betterIsAnagram(String s, String t) {
        char[] sChars = s.toCharArray();
        char[] tChars = t.toCharArray();

        Arrays.sort(sChars);
        Arrays.sort(tChars);

        return Arrays.equals(sChars, tChars);
    }
}