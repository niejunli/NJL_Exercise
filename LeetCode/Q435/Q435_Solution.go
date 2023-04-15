package Q435

/*
贪心算法 --- 无重叠区间
按照右边界进行排序-因为区间的右边界end越小，可以留给后面的空间就越大。因为是要统计最小移除的区间数量，也就是保留最多的区间数量，既然是数量，那么为后面预留越多空间越好
*/

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1] // 将intervals按照end大小升序排序
    })
    end := intervals[0][1] // 结束最早的区间
    cnt := 1 // 统计互不重叠的区间数量
    for i:=1; i<len(intervals); i++ {
        if intervals[i][0] >= end { // 某个区间的起始点>end，表示不重叠
            cnt++
            end = intervals[i][1] // 更新end
        } 
    }
    return len(intervals)-cnt // 总数量-不重叠组数量=需要移除的数量
}