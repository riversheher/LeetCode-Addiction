package validParenthases;

import java.util.ArrayDeque;
import java.util.Deque;

public class Solution {

    public boolean isValid(String s) {

        Deque<Character> parenStack = new ArrayDeque<>();

        for(char paren : s.toCharArray()) {
            
            if(paren == '{' || paren == '(' || paren == '['){
                parenStack.push(paren);
            } else {
                if (parenStack.isEmpty()) {
                    return false;
                }
                switch (paren) {
                    case '}':
                        if(parenStack.peek().equals('{')) {
                            parenStack.pop();
                        } else {
                            return false;
                        }
                        break;
                    case ')':
                        if(parenStack.peek().equals('(')) {
                            parenStack.pop();
                        } else {
                            return false;
                        }
                        break;

                    case ']':
                        if(parenStack.peek().equals('[')) {
                            parenStack.pop();
                        } else {
                            return false;
                        }
                        break;
                    default:
                        break;
                }
            }

        }

        return parenStack.isEmpty();
    }

    /*
    This solution relies on a different thought process.
    Instead of pushing the open parentheses, just push 
    the expected close parenthases instead!

    This simplifies the logic greatly.
    */
    public boolean betterIsValid(String s) {
        Deque<Character> parenStack = new ArrayDeque<>();

        for(char paren : s.toCharArray()) {
            switch (paren) {
                case '(':
                    parenStack.push(')');
                    break;
                case '[':
                    parenStack.push(']');
                    break;
                case '{':
                    parenStack.push('}');
                    break;
                default:
                    if(parenStack.isEmpty() || paren != parenStack.pop()) {
                        return false;
                    }
                    break;
            }
        }

        return parenStack.isEmpty();
    }

    /*
    Example with modern switch statement.
    */
    public boolean betterSwitchIsValid(String s) {
        Deque<Character> stack = new ArrayDeque<>();

        for(char paren : s.toCharArray()) {
            switch (paren) {
                case '(' -> stack.push(')');
                case '{' -> stack.push('}');
                case '[' -> stack.push(']');
                default -> { if (stack.isEmpty() || stack.pop() != paren) return false; }
            }
        }

        return stack.isEmpty();
    }
    
}
