package dp

// leetcode #746


// NOTES:

// STATE VARIABLE: state variable is the index of the current step and the cost to get to it without its own cost. 
//                 We only count it's cost when we are taking this step.             
// TERMINATION CONDITION: final step should be above last one, but might not include the last step as the step size can be 1 OR 2.

func MinCostClimbingStairsTopDown(cost []int) int {
    memo := map[int]int{}
    cost = append(cost, 0)
    return climb(cost, len(cost)-1, &memo)
}

func climb(cost []int, i int, memo *map[int]int) int {
    if stepCost, exist := (*memo)[i]; exist {
        return stepCost
    }
    if i <= 1 {
        return cost[i]
    }
    (*memo)[i] = min(climb(cost, i-1, memo), climb(cost, i-2, memo)) + cost[i]
    return (*memo)[i]
}


func MinCostClimbingStairsBottomUp(cost []int) int {
   states := make([]int, len(cost) + 1)

   for step := 2; step < len(states); step++ {
       states[step] = min(states[step-2] + cost[step-2], states[step-1] + cost[step-1])
   }
   return states[len(cost)]
}

func min(v1, v2 int) int {
    if v1 < v2 {
        return v1
    }
    return v2
}
