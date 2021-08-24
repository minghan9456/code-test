// https://leetcode.com/problems/unique-binary-search-trees/

func numTrees(n int) int {
    // [1 - n]
    // 0 -> 1
    // 1 => 1
    // 2 => 2
    // 3 => 5
    // 4 => 14
    
    ans := make([]int, n+1)
    ans[0] = 1
    ans[1] = 1
    
    for i := 2; i <= n; i++ {
        for j := 1; j <= i; j++ {
            // fmt.Println(i, j, "=", j-1, i-j)
            ans[i] += ans[j-1] * ans[i-j]
        }
    }
    
    return ans[n]
}

/*    
Hope it will help you to understand :
    
    n = 0;     null   
    
    count[0] = 1
    
    
    n = 1;      1       
    
    count[1] = 1 
    
    
    n = 2;    1__       			 __2     
    		      \					/                 
    		     count[1]	   	count[1]	
    
    count[2] = 1 + 1 = 2
    
    
    
    n = 3;    1__				      __2__	                   __3
    		      \		            /       \			      /		
    		  count[2]		  count[1]    count[1]		count[2]
    
    count[3] = 2 + 1 + 2  = 5
    
    
    
    n = 4;    1__  					__2__					   ___3___                  
    		      \				 /        \					  /		  \			
    		  count[3]		 count[1]    count[2]		  count[2]   count[1]
    
                 __4				
               /
           count[3]   
    
    count[4] = 5 + 2 + 2 + 5 = 14     
    

And  so on...
*/
