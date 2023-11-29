public class Solution {
    public string LongestCommonPrefix(string[] strs) {
        if (strs == null || strs.Length == 0) {
            return "";
        }

        string prefix = strs[0];

        for (int i = 1; i < strs.Length; i++) {
            int j = 0;
            while (j < prefix.Length && j < strs[i].Length && prefix[j] == strs[i][j]) {
                j++;
            }

            prefix = prefix.Substring(0, j);

            if (string.IsNullOrEmpty(prefix)) {
                return "";
            }
        }

        return prefix;
    }
}
