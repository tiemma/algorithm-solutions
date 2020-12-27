type set map[int]map[int]bool
type counter map[int]int

func findOrder(numCourses int, prerequisites [][]int) []int {
    courses := set{}
    degrees := counter{}
    result := []int{}
    if len(prerequisites) == 0{
        for i := numCourses - 1; i >= 0; i -- {
            result = append(result, i)
        }
        return result
    }
    for i := 0; i < numCourses; i++ {
        degrees[i] = 0
        if _, ok := courses[i]; !ok {
            courses[i] = make(map[int]bool)
        }
    }
    for _, course := range prerequisites {
        dep := course[1]
        preq := course[0]
        if _, ok := courses[preq][dep]; ok {
            return []int{}
        }
        courses[dep][preq] = true
        degrees[preq] += 1
    }
    queue := make([]int, 0)
    for c, degree := range degrees {
        if degree == 0 {
            queue = append(queue, c)
        }
    }
    for len(queue) > 0 {
        course := queue[0]
        queue = queue[1:]
        result = append(result, course)
        for c := range courses[course] {
            degrees[c] -= 1
            if degrees[c] == 0 {
                queue = append(queue, c)
            }
        }
    }
    if len(result) != len(degrees) {
        return []int{}
    }
    return result
}
