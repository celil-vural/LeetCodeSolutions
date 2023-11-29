using System;

public class Solution {
    public int ThreeSumClosest(int[] nums, int target) {
        if (nums == null || nums.Length < 3) {
            throw new ArgumentException("Invalid input");
        }

        Array.Sort(nums);
        int closestSum = nums[0] + nums[1] + nums[2];

        for (int i = 0; i < nums.Length - 2; i++) {
            int left = i + 1;
            int right = nums.Length - 1;

            while (left < right) {
                int sum = nums[i] + nums[left] + nums[right];

                if (Math.Abs(sum - target) < Math.Abs(closestSum - target)) {
                    closestSum = sum;
                }

                if (sum < target) {
                    left++;
                } else if (sum > target) {
                    right--;
                } else {
                    return sum; // Eşitlik durumunda en yakın toplamı bulduk
                }
            }
        }

        return closestSum;
    }
}
