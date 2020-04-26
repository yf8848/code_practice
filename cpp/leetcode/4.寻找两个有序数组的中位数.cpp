/*
 * @lc app=leetcode.cn id=4 lang=cpp
 *
 * [4] 寻找两个有序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (37.32%)
 * Likes:    2465
 * Dislikes: 0
 * Total Accepted:    177.7K
 * Total Submissions: 476.2K
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
 * 
 * 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 * 
 * 你可以假设 nums1 和 nums2 不会同时为空。
 * 
 * 示例 1:
 * 
 * nums1 = [1, 3]
 * nums2 = [2]
 * 
 * 则中位数是 2.0
 * 
 * 
 * 示例 2:
 * 
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 * 
 * 则中位数是 (2 + 3)/2 = 2.5
 * 
 * 
 */

// @lc code=start
class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int n = nums1.size();
        int m = nums2.size();

        int left = (m+n+1)/2;
        int right = (m+n+2)/2;

        return (getKth(nums1, 0, n-1, nums2,0,m-1,left)+getKth(nums1, 0, n-1, nums2,0,m-1,right))*0.5;
    }

    int getKth(vector<int> nums1, int start, int end, vector<int>& nums2, int start2, int end2, int k){
        int len1 = end - start +1;
        int len2 = end2 - start2 +1;

        if (len1 > len2){
            return getKth(nums2, start2, end2, nums1, start, end, k);
        }
        if (len1 == 0){
            return nums2[start2+k-1];
        }

        if (k == 1) {
            return min(nums1[start], nums2[start2]);
        }

        int i = start+ min(len1, k/2)-1;
        int j = start2 +min(len2, k/2)-1;

        if (nums1[i]>nums2[j]){
            return getKth(nums1, start, end, nums2, j+1, end2, k-(j-start2+1));
        }else{
            return getKth(nums1, i+1, end, nums2, start2, end2, k-(i-start+1));
        }

    }
};
// @lc code=end

