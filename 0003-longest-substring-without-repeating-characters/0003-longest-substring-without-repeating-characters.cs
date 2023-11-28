public class Solution {
    public int LengthOfLongestSubstring(string s) {
        int n = s.Length;
        int maxLength = 0;
        Dictionary<char, int> lastIndex = new Dictionary<char, int>();
        for (int j = 0, i = 0; j < n; j++) {
            if (lastIndex.ContainsKey(s[j])) {
                i = Math.Max(lastIndex[s[j]], i);
            }
            maxLength = Math.Max(maxLength, j - i + 1);
            lastIndex[s[j]] = j + 1;
        }
        return maxLength;
    }
}