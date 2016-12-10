import java.io.*;
import java.util.*;

public class Solution {
    private Stack palindromeStack;
    private LinkedList palindromeQueue;
    

    public Solution() {
        this.palindromeStack = new Stack<>();
        this.palindromeQueue  = new LinkedList<>();
    }

    public void pushCharacter(char ch) {
        this.palindromeStack.push(ch);
    }

    public void enqueueCharacter(char c) {
        this.palindromeQueue.addFirst(c);
    }

    public char popCharacter() {
        return (char) this.palindromeStack.pop();
    }

    public char dequeueCharacter() {
        return (char) this.palindromeQueue.removeLast();
    }




        public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        String input = scan.nextLine();
        scan.close();

        // Convert input String to an array of characters:
        char[] s = input.toCharArray();

        // Create a Solution object:
        Solution p = new Solution();

        // Enqueue/Push all chars to their respective data structures:
        for (char c : s) {
            p.pushCharacter(c);
            p.enqueueCharacter(c);
        }

        // Pop/Dequeue the chars at the head of both data structures and compare them:
        boolean isPalindrome = true;
        for (int i = 0; i < s.length/2; i++) {
            if (p.popCharacter() != p.dequeueCharacter()) {
                isPalindrome = false;                
                break;
            }
        }

        //Finally, print whether string s is palindrome or not.
        System.out.println( "The word, " + input + ", is " 
                           + ( (!isPalindrome) ? "not a palindrome." : "a palindrome." ) );
    }
}