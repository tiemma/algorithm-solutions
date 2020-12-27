

# https://leetcode.com/problems/two-city-scheduling/

# Two City Scheduling
class Solution(object):
    def twoCitySchedCost(self, costs):
        """
        :type costs: List[List[int]]
        :rtype: int
        """
        cost = 0
        costs = sorted(costs, key = lambda x: x[0] - x[1])
        avg = len(costs) / 2
        cost += sum([x[0] for x in costs[:avg]])
        cost += sum([x[1] for x in costs[avg:]])
        return cost

