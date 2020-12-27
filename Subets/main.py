

# https://leetcode.com/problems/subsets/

# O(n * 2^n)
class Solution(object):
    def subsets(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        arr = []
        for i in range(0, pow(2, len(nums))):
            val = format(i, "#0{}b".format(len(nums)+2)).replace('0b', '')
            print(val)
            temp = [x[1] for x in zip(val, nums) if x[0] == '1']
            arr.append(temp)
        return arr




    [1, 2, 3, 4, 5]


    1 - 0 0 0 0 1
