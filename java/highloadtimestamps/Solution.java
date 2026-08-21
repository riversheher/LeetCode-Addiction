package highloadtimestamps;

import java.util.ArrayList;
import java.util.List;

public class Solution {

    // Return list of indicies where load is higher than average.
    public List<Integer> highLoadsIndices(int[] loads) {
        List<Integer> highs = new ArrayList<>();
        if (loads.length == 0) return highs;

        long sum = 0;
        for (int load : loads) sum += load;

        long n = loads.length;
        for (int i = 0; i < n; i++) {
            if (loads[i] * n > 2 * sum) {   // loads[i] > 2*(sum/n), no float, no truncation
                highs.add(i);
            }
        }
        return highs;
    }
    
}
