package validPalindrome;

public class Solution {

    public boolean isPalindrome(String s) {

        for(int i = updateLeft(0, s), j = updateRight(s.length() - 1, s); 
        i < j; 
        i = updateLeft(++i, s), j = updateRight(--j, s)){

            if(!(Character.toLowerCase(s.charAt(i)) == Character.toLowerCase(s.charAt(j)))) {
                return false;
            }

        }

        return true;
    }

    public int updateLeft(int i, String s){

        while(i < s.length() && !Character.isLetterOrDigit(s.charAt(i))){
            i++;
        }

        return i;
    }

    public int updateRight(int j, String s) {

        while(j >= 0 && !Character.isLetterOrDigit(s.charAt(j))) {
            j--;
        }

        return j;

    }


    /*
    This is probably the better option as it avoids over complication,
    making it easier to read and write during an assessment.
    */
    public boolean isPalindromeSimplified(String s) {
        int i = 0;
        int j = s.length() - 1;

        while (i < j) {
            while (i < j && !Character.isLetterOrDigit(s.charAt(i))){
                i++;
            }
            while (i < j && !Character.isLetterOrDigit(s.charAt(j))) {
                j--;
            }

            if(!(Character.toLowerCase(s.charAt(i)) == Character.toLowerCase(s.charAt(j)))) {
                return false;
            }
            i++;
            j++;
        }

        return true;
    }
}
