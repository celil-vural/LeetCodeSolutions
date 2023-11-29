public class Solution {
    public bool IsValid(string s) {
        if (string.IsNullOrEmpty(s)) {
            return true;
        }
        Stack<char> stack = new Stack<char>();
        foreach (char c in s) {
            if (c == '(' || c == '{' || c == '[') {
                stack.Push(c);
            } else {
                if (stack.Count == 0) {
                    return false;
                }
                char open = stack.Pop();
                if ((c == ')' && open != '(') || (c == '}' && open != '{') || (c == ']' && open != '[')) {
                    return false; 
                }
            }
        }
        return stack.Count == 0;
    }
}
