public class Solution {
    public int Reverse(int x) {
        int INT_MAX = int.MaxValue;
        int INT_MIN = int.MinValue;
        int reversedNum = 0;
        while (x != 0) {
            int digit = x % 10;
            x /= 10;
            if (reversedNum > INT_MAX / 10 || (reversedNum == INT_MAX / 10 && digit > 7)) {
                return 0;
            }
            if (reversedNum < INT_MIN / 10 || (reversedNum == INT_MIN / 10 && digit < -8)) {
                return 0;
            }
            reversedNum = reversedNum * 10 + digit;
        }
        return reversedNum;
    }
}