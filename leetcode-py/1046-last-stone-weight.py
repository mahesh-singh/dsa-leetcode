class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:

        stones_heap = [-s for s in stones]
        heapq.heapify(stones_heap)

        while len(stones_heap) > 1:
            first_stone = -heapq.heappop(stones_heap)
            second_stone = -heapq.heappop(stones_heap)

            if first_stone > second_stone:
                heapq.heappush(stones_heap, -(first_stone - second_stone))


        if len(stones_heap) == 1:
            return -stones_heap[0]

        return 0
