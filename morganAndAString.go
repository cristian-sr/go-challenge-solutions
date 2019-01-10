/* Jack and Daniel are friends. Both of them like letters, especially upper-case ones. 
They are cutting upper-case letters from newspapers, and each one of them has his collection of letters stored in a stack.

One beautiful day, Morgan visited Jack and Daniel. He saw their collections. He wondered what is the lexicographically minimal string made of those two collections. He can take a letter from a collection only when it is on the top of the stack. Morgan wants to use all of the letters in their collections.

As an example, assume Jack has collected a = [A, C, A] and Daniel has b = [B, C, F]. The example shows the top at index 0 for each stack of letters. Assembling the string would go as follows:

Jack	Daniel	result
ACA	BCF
CA	BCF	A
CA	CF	AB
A	CF	ABC
A	CF	ABCA
    	F	ABCAC
    		ABCACF
Note the choice when there was a tie at CA and CF.

Function Description

Complete the morganAndString function in the editor below. It should return the completed string.

morganAndString has the following parameter(s):

a: a string representing Jack's letters, top at index 0 
b: a string representing Daniel's letters, top at index 0 
Input Format

The first line contains the an integer t, the number of test cases.

The next t pairs of lines are as follows: 
- The first line contains string a  
- The second line contains string b.

Constraints
1. 1<=T<=5
2. 1<=|a|,|b|<=10^5
3. a and b contain upper-case letters onlay, ascii[A-Z].
Output Format

Output the lexicographically minimal string result for each test case in new line.

Sample Input

2
JACK
DANIEL
ABACABA
ABACABA
Sample Output

DAJACKNIEL
AABABACABACABA
Explanation

The first letters to choose from were J and D since they were at the top of the stack. D was chosen, the options then were J and A. A chosen. Then the two stacks have J and N, so J is chosen. (Current string is DAJ) Continuing this way till the end gives us the resulting string. */

/*
Initial thoughts:
Since we can only pull from the top of each 
letter stack each time, we simply need to just
compare the letters at the top of each stack
In each comparison we just need to choose the
letter that is lexicographically smaller than
the other and print it out and continue this 
until 1 of the stacks is empty.
If we have equivalent characters, we have to decide 
which one to pick
When they are equivalent, we choose the one from the 
stack that is overall lexicographically smaller
Read inline comments for better understanding of how
I handle this
Then if we have finished 1 string early we just need 
to add letters from the other stack to the end of the 
string we built
Time Complexity: O(|a|+|b|^2)   //We only view each letter once
Space Complexity: O(|a| + |b|)  //We store out output in a SB to speed up run time
 */


 package main 
 import "bufio"
 import "fmt"
 import "os"

 func morgan(s1, s2 string) string {
	 output := make([]byte, 0, len(s1)+len(s2))
	 for i, j := 0, 0; i < len(s1) || j < len(s2); {
		 if i == len(s1) {
			 output = append(output, s2[j:]...)
			 break
		 } else if j == len(s2) {
			 output = append(output, s1[i:]...)
			 break
		 } else if s1[i] < s2[j] {
			 output = append(output, s1[i])
			 i++
		 } else if s1[i] > s2[j] {
			 output = append(output, s2[j])
			 j++
		 } else if s1[i+1:]+"[" <= s2[j+1:]+"[" {
			 for c := s1[i]; i < len(s1) && s1[i] == c; i++ {
				 output = append(output, c)
			 }
		 } else if s1[i+1:]+"[" > s2[j+1:]+"[" {
			 for c := s2[j]; j < len(s2) && s2[j] == c; j++ {
				 output = append(output, c)
			 }
		 }
	 }
	 return string(output)
 }

 func main() {
	 in := bufio.NewReader(os.Stdin)
	 N, s1, s2 := 0, "", ""
	 fmt.Fscan(in, &N)
	 for i := 0; i < N; i++ {
		 fmt.Fscan(in, &s1, &s2)
		 fmt.Println(morgan(s1, s2))
	 }
 }
 