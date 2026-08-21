package com.example;

public class Solution {

    class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }

    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode solutionHead = new ListNode();

        ListNode solutionCurrent = null;
        int carry = 0;
        while(l1 != null || l2 != null) {

            // Unneeded check, use dummy head
            if(solutionCurrent == null) {
                solutionCurrent = solutionHead;
            } else {
                solutionCurrent.next = new ListNode();
                solutionCurrent = solutionCurrent.next;
            }

            int sum = 0;

            if(l1 == null) {
                sum = l2.val + carry;
                l2 = l2.next;
            } else if(l2 == null) {
                sum = l1.val + carry;
                l1 = l1.next;
            } else {
                sum = l1.val + l2.val + carry;
                l1 = l1.next;
                l2 = l2.next;
            }

            if(sum >= 10) {
                carry = 1;
                solutionCurrent.val = sum % 10;
            } else {
                carry = 0;
                solutionCurrent.val = sum;
            }
        }

        if (carry == 1) {
            solutionCurrent.next = new ListNode(carry);
        }

        return solutionHead;
    }

    public ListNode addTwoNumbersSimplified(ListNode l1, ListNode l2) {
        ListNode dummy = new ListNode(0);
        ListNode curr = dummy;
        int carry = 0;

        while (l1 != null || l2 != null || carry != 0) {   // carry drains inside the loop
            
            //simplify if statements my using the insight that a null node gives a 0 value
            int v1 = (l1 == null) ? 0 : l1.val;
            int v2 = (l2 == null) ? 0 : l2.val;

            int sum = v1 + v2 + carry;
            
            //using sum / 10 allows us to modify this to non base-10 if needed
            carry = sum / 10;

            curr.next = new ListNode(sum % 10);
            curr = curr.next;

            if (l1 != null) l1 = l1.next;
            if (l2 != null) l2 = l2.next;
        }

        // dummy head allows us to avoid the extra check I added in the loop.
        return dummy.next;
    }
}