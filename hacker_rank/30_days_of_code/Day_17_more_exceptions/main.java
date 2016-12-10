import java.util.*;
import java.io.*;


class Calculator {
    private int n;
    private int p;
    
    public Calculator() {
        
    }
    
    protected int power(int n, int p) throws Exception {
        //System.out.println("n=" + n + ". p=" + p);
        if ((n < 0) || (p < 0)) {
            throw new Exception("n and p should be non-negative");
        }
        return (int)Math.pow(n,  p);
        //return n * power(n, p-1);
    }
    
}

class Solution{

    public static void main(String []argh)
    {
        Scanner in = new Scanner(System.in);
        int T=in.nextInt();
        while(T-->0)
        {
            int n = in.nextInt();
            int p = in.nextInt();
            Calculator myCalculator = new Calculator();
            try
            {
                int ans=myCalculator.power(n,p);
                System.out.println(ans);
                
            }
            catch(Exception e)
            {
                System.out.println(e.getMessage());
            }
        }

    }
}
