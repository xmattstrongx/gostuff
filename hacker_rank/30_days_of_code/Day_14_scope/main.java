import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;


class Difference {
  	private int[] elements;
  	public int maximumDifference;

	
public Difference(int[] elements) {
    this.elements = elements;
}

public void computeDifference() {
    int first = 0;
    int least = 0;
    int max = 0; 
    for(int value : elements) {
        if (first == 0) {
            first = value;
            least = value;
            max = value;
        }
        if (value < least) {
            least = value;
        }
        if (value > max) {
            max = value;
        }
    }
    maximumDifference =  Math.abs(max - least);
    }
    } // End of Difference class

public class Solution {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int[] a = new int[n];
        for (int i = 0; i < n; i++) {
            a[i] = sc.nextInt();
        }
        sc.close();

        Difference difference = new Difference(a);

        difference.computeDifference();

        System.out.print(difference.maximumDifference);
    }
}