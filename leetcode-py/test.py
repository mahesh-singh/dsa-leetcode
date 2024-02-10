import heapq

if __name__ == "__main__":
    heap = []
    heapq.heappush(heap, 2)
    heapq.heappush(heap, 3)
    heapq.heappush(heap, 1)
    heapq.heappush(heap, 8)

    print(heap)
