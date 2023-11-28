public class Solution {
    /*public double FindMedianSortedArrays(int[] nums1, int[] nums2) {
        int totalLength =nums1.Length + nums2.Length;
        int[] combinedArray = new int[totalLength];
        Array.Copy(nums1, combinedArray, nums1.Length);
        Array.Copy(nums2, 0, combinedArray, nums1.Length, nums2.Length);
        Array.Sort(combinedArray);
        return totalLength % 2 == 0 ? (combinedArray[totalLength / 2 - 1] + combinedArray[totalLength / 2]) / 2.0 : combinedArray[totalLength / 2];
    }*/
    public double FindMedianSortedArrays(int[] nums1, int[] nums2) {
        int m = nums1.Length;
        int n = nums2.Length;
        if (m > n) {
            return FindMedianSortedArrays(nums2, nums1); 
        }
        int iMin = 0, iMax = m, halfLen = (m + n + 1) / 2;
        while (iMin <= iMax) {
            int i = (iMin + iMax) / 2;
            int j = halfLen - i;
            if (i < iMax && nums2[j - 1] > nums1[i]) {
                iMin = i + 1; 
            } else if (i > iMin && nums1[i - 1] > nums2[j]) {
                iMax = i - 1;
            } else {
                int maxLeft = 0;
                if (i == 0) {
                    maxLeft = nums2[j - 1];
                } else if (j == 0) {
                    maxLeft = nums1[i - 1];
                } else {
                    maxLeft = Math.Max(nums1[i - 1], nums2[j - 1]);
                }
                if ((m + n) % 2 == 1) {
                    return maxLeft;
                }
                int minRight = 0;
                if (i == m) {
                    minRight = nums2[j];
                } else if (j == n) {
                    minRight = nums1[i];
                } else {
                    minRight = Math.Min(nums1[i], nums2[j]);
                }
                return (maxLeft + minRight) / 2.0;
            }
        }
        return 0.0;
    }
}